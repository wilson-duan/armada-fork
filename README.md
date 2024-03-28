<div align="center">
<img src="./logo.svg" width="200"/>

<p><a href="https://circleci.com/gh/armadaproject/armada"><img src="https://circleci.com/gh/helm/helm.svg?style=shield" alt="CircleCI"></a>
<a href="https://goreportcard.com/report/github.com/armadaproject/armada"><img src="https://goreportcard.com/badge/github.com/armadaproject/armada" alt="Go Report Card"></a></p>
</div>

# Armada

Armada is a high-throughput [batch scheduler](https://en.wikipedia.org/wiki/Job_scheduler) built on top of [Kubernetes](https://kubernetes.io/docs/concepts/overview/). Armada is used for critical tasks in a live business environment. Armada can run millions of jobs per day across thousands of nodes. No other batch scheduler for Kubernetes can currently do this at the scale Armada is capable of.

Armada addresses the following limitations of Kubernetes:

1. Scaling and operating a single Kubernetes cluster beyond about 1000 nodes is [challenging](https://openai.com/blog/scaling-kubernetes-to-7500-nodes/). One contributing reason is that the number of [watch requests](https://etcd.io/docs/v3.2/learning/api/#watch-api) scales with the square of the number of nodes in some cases. Hence, Armada is designed to effectively schedule jobs across many Kubernetes clusters. Many thousands of nodes can be managed by Armada in this way. Further, Kubernetes clusters can be connected and disconnected in a manner transparent to users, which simplifies, e.g., maintenance and upgrades.
2. Achieving very high throughput using the in-cluster storage backend, etcd, is [challenging](https://etcd.io/docs/v3.5/op-guide/performance/) as a consequence of the stringent concistency and timeliness guarantees provided by etcd. Hence, Armada performs queueing and scheduling out-of-cluster using a specialized storage layer, which trades off timeliness for higher throughput. This allows Armada to maintain queues composed of millions of jobs.
3. The default [kube-scheduler](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-scheduler/) lacks many features necessary for anything but trivial batch workflows. Instead, Armada includes a novel multi-Kubernetes cluster scheduler with support for such features, including:
   * Sophisticated queuing and scheduling algorithms to fairly divide scarce cluster resources between competing requests.
   * Scheduling limits, which prevent individual users from consuming more than their allotted share of resources.
   * Gang-scheduling, i.e., atomically scheduling sets of related jobs, as is necessary for, e.g., modern distributed machine learning applications.
   * State-of-the-art preemption algorithms. For example, Armada can automatically preempt running jobs to balance resource allocation between multiple users, or to make resources available for an urgent job immediately.

Armada also provides features to help manage large compute clusters effectively, including:

* Detailed analytics exposed via [Prometheus](https://prometheus.io/) showing how the system behaves and how resources are allocated.
* Automatic removal of nodes exhibiting high failure rates from consideration for scheduling.
* A mechanism to earmark nodes for a particular set of jobs, but allowing them to be used by other jobs when not used for their primary purpose.

Armada is designed with the enterprise in mind; all components are secure and highly available.

Armada is a [CNCF](https://www.cncf.io/) Sandbox project and is used in production at [G-Research](https://www.gresearch.co.uk/).

### Limitations

## Documentation

For documentation, see the following:

- [System overview](./docs/system_overview.md)
- [Scheduler](./docs/scheduler.md)
- [User guide](./docs/user.md)
- [Quickstart](./docs/quickstart/index.md)
- [Development guide](./docs/developer.md)
- [API Documentation](./docs/developer/api.md)

We expect readers of the documentation to have a basic understanding of Docker and Kubernetes; see, e.g., the following links:

- [Docker overiew](https://docs.docker.com/get-started/overview/)
- [Kubernetes overview](https://kubernetes.io/docs/concepts/overview/)

## Presentations

Presentations relating to Armada:

- [Armada - high-throughput batch scheduling](https://www.youtube.com/watch?v=FT8pXYciD9A)
- [Building Armada - Running Batch Jobs at Massive Scale on Kubernetes](https://www.youtube.com/watch?v=B3WPxw3OUl4)

## Contributions

Thank you for considering contributing to Armada!
We want everyone to feel that they can contribute to the Armada Project.
Your contributions are valuable, whether it's fixing a bug, implementing a new feature, improving documentation, or suggesting enhancements.
We appreciate your time and effort in helping make this project better for everyone.
For more information about contributing to Armada see [CONTRIBUTING.md](https://github.com/armadaproject/armada/blob/master/CONTRIBUTING.md) and before proceeding to contributions see [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md)

The Armada project adheres to the CNCF [Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md).

## Discussion

If you are interested in discussing Armada you can find us on  [![slack](https://img.shields.io/badge/slack-armada-brightgreen.svg?logo=slack)](https://cloud-native.slack.com/?redir=%2Farchives%2FC03T9CBCEMC)
