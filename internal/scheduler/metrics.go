package scheduler

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/clock"

	commmonmetrics "github.com/armadaproject/armada/internal/common/metrics"
	"github.com/armadaproject/armada/internal/common/resource"
	"github.com/armadaproject/armada/internal/scheduler/database"
	"github.com/armadaproject/armada/internal/scheduler/jobdb"
)

// stores the metrics state associated with a queue
type queueState struct {
	numQueuedJobs      int
	queuedJobRecorder  *commmonmetrics.JobMetricsRecorder
	runningJobRecorder *commmonmetrics.JobMetricsRecorder
}

// a snapshot of metrics.  Implements QueueMetricProvider
type metricsState struct {
	queues      []*database.Queue
	queueStates map[string]*queueState
}

func (m metricsState) GetQueuedJobMetrics(queueName string) []*commmonmetrics.QueueMetrics {
	state, ok := m.queueStates[queueName]
	if ok {
		return state.queuedJobRecorder.Metrics()
	}
	return nil
}

func (m metricsState) GetRunningJobMetrics(queueName string) []*commmonmetrics.QueueMetrics {
	state, ok := m.queueStates[queueName]
	if ok {
		return state.runningJobRecorder.Metrics()
	}
	return nil
}

func (m metricsState) numQueuedJobs() map[string]int {
	queueCounts := make(map[string]int)
	for _, queue := range m.queues {
		state, ok := m.queueStates[queue.Name]
		count := 0
		if ok {
			count = state.numQueuedJobs
		}
		queueCounts[queue.Name] = count
	}
	return queueCounts
}

// MetricsCollector is a Prometheus Collector that handles scheduler metrics.
// The metrics themselves are calculated asynchronously every refreshPeriod
type MetricsCollector struct {
	jobDb           *jobdb.JobDb
	queueRepository database.QueueRepository
	poolAssigner    PoolAssigner
	refreshPeriod   time.Duration
	clock           clock.Clock
	state           atomic.Value
}

func NewMetricsCollector(
	jobDb *jobdb.JobDb,
	queueRepository database.QueueRepository,
	poolAssigner PoolAssigner,
	refreshPeriod time.Duration,
) *MetricsCollector {
	return &MetricsCollector{
		jobDb:           jobDb,
		queueRepository: queueRepository,
		poolAssigner:    poolAssigner,
		refreshPeriod:   refreshPeriod,
		clock:           clock.RealClock{},
		state:           atomic.Value{},
	}
}

// Run enters s a loop which updates the metrics every refreshPeriod until the supplied comtext is cancelled
func (c *MetricsCollector) Run(ctx context.Context) error {
	ticker := c.clock.NewTicker(c.refreshPeriod)
	log.Infof("Will update metrics every %s", c.refreshPeriod)
	for {
		select {
		case <-ctx.Done():
			log.Infof("Context cancelled, returning..")
			return nil
		case <-ticker.C():
			err := c.refresh(ctx)
			if err != nil {
				log.WithError(err).Warnf("error refreshing metrics state")
			}
		}
	}
}

// Describe returns all descriptions of the collector.
func (c *MetricsCollector) Describe(out chan<- *prometheus.Desc) {
	commmonmetrics.Describe(out)
}

// Collect returns the current state of all metrics of the collector.
func (c *MetricsCollector) Collect(metrics chan<- prometheus.Metric) {
	state, ok := c.state.Load().(metricsState)
	if ok {
		commmonmetrics.CollectQueueMetrics(state.numQueuedJobs(), state, metrics)
	}
}

func (c *MetricsCollector) refresh(ctx context.Context) error {
	log.Debugf("Refreshing prometheus metrics")
	start := time.Now()

	queues, err := c.queueRepository.GetAllQueues()
	if err != nil {
		return err
	}

	err = c.poolAssigner.Refresh(ctx)
	if err != nil {
		return err
	}

	metricsState := metricsState{
		queues:      queues,
		queueStates: map[string]*queueState{},
	}
	for _, queue := range queues {
		metricsState.queueStates[queue.Name] = &queueState{
			queuedJobRecorder:  commmonmetrics.NewJobMetricsRecorder(),
			runningJobRecorder: commmonmetrics.NewJobMetricsRecorder(),
		}
	}

	currentTime := c.clock.Now()
	iter, err := jobdb.NewAllJobsIterator(c.jobDb.ReadTxn())
	if err != nil {
		return err
	}
	job := iter.NextJobItem()
	for job != nil {
		// Don't calculate metrics for dead jobs
		if job.InTerminalState() {
			continue
		}
		queueState, ok := metricsState.queueStates[job.Queue()]
		if !ok {
			log.Warnf("Job %s is in queue %s, but this queue does not exist.  Skipping", job.Id(), job.Queue())
			continue
		}

		pool, err := c.poolAssigner.AssignPool(job)
		if err != nil {
			return err
		}

		priorityClass := job.JobSchedulingInfo().PriorityClassName
		resourceRequirements := job.JobSchedulingInfo().GetObjectRequirements()[0].GetPodRequirements().GetResourceRequirements().Requests
		jobResources := make(map[string]float64)
		for key, value := range resourceRequirements {
			jobResources[string(key)] = resource.QuantityAsFloat64(value)
		}

		var recorder *commmonmetrics.JobMetricsRecorder
		var timeInState time.Duration
		if job.Queued() {
			recorder = queueState.queuedJobRecorder
			timeInState = currentTime.Sub(time.Unix(0, job.Created()))
		} else if job.HasRuns() {
			run := job.LatestRun()
			timeInState = currentTime.Sub(time.UnixMicro(run.Created()))
			recorder = queueState.runningJobRecorder
		} else {
			log.Warnf("Job %s is marked as leased but has no runs", job.Id())
		}
		recorder.RecordJobRuntime(pool, priorityClass, timeInState)
		recorder.RecordResources(pool, priorityClass, jobResources)

		queueState.numQueuedJobs++
		job = iter.NextJobItem()
	}
	c.state.Store(metricsState)
	log.Debugf("Refreshed prometheus metrics in %s", time.Since(start))
	return nil
}