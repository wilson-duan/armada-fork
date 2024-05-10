package util

import (
	v1 "k8s.io/api/core/v1"

	"github.com/armadaproject/armada/pkg/api"
)

func PodSpecFromJob(job *api.Job) *v1.PodSpec {
	// TODO: Remove
	if job.JobSpec != nil {
		return &job.JobSpec.Template.Spec
	}
	for _, jobSpec := range job.JobSpecs {
		if jobSpec != nil {
			return &jobSpec.Template.Spec
		}
	}
	if job.PodSpec != nil {
		return job.PodSpec
	}
	for _, podSpec := range job.PodSpecs {
		if podSpec != nil {
			return podSpec
		}
	}
	return nil
}
