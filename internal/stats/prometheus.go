package stats

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func MeasureLatency(observer prometheus.Observer, startTime time.Time) {
	observer.Observe(time.Since(startTime).Seconds())
}


func GetMetrics() *DKVMetrics {
	dkvMetrics := newDKVMetric()
	mfs, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		fmt.Println("yo...yo...yo...")
	}
	for _,mf :=  range mfs {
		switch mf.GetName() {
		case "storage_latency" :
			for _ , m := range mf.GetMetric(){
					dkvMetrics.StoreLatency[m.Label[0].GetValue()] = newPercentile(m.GetSummary().GetQuantile())
					dkvMetrics.StorageOpsCount[m.Label[0].GetValue()] = m.GetSummary().GetSampleCount()
			}
		case "nexus_latency" :
			for _ , m := range mf.GetMetric(){
				dkvMetrics.NexusLatency[m.Label[0].GetValue()] = newPercentile(m.GetSummary().GetQuantile())
				dkvMetrics.NexusOpsCount[m.Label[0].GetValue()] = m.GetSummary().GetSampleCount()
			}
		case "grpc_server_handled_total":
			for _ , m := range mf.GetMetric(){
				if _ , exist:= dkvMetrics.GrpcResponseCount[m.Label[0].GetValue()] ; exist {
					dkvMetrics.GrpcResponseCount[m.Label[0].GetValue()] += m.GetCounter().GetValue()
				} else {
					dkvMetrics.GrpcResponseCount[m.Label[0].GetValue()] = m.GetCounter().GetValue()
				}
			}
		case "storage_error":
			for _ , m := range mf.GetMetric(){
				dkvMetrics.StorageOpsErrorCount[m.Label[0].GetValue()] = m.GetCounter().GetValue()
			}
		}
	}
	return dkvMetrics
}
