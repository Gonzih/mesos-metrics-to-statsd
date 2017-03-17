# Reprot mesos master metrics to statsd

## Installation

```
go get github.com/Gonzih/mesos-metrics-to-statsd
```

## Usage

```
mesos-metrics-to-statsd --mesos-master 192.168.9.144:5050 --statsd-host 192.168.227.149:8125 --statsd-prefix mesos.metrics.
```
