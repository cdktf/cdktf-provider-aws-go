package evidentlylaunch


type EvidentlyLaunchMetricMonitors struct {
	// metric_definition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/evidently_launch#metric_definition EvidentlyLaunch#metric_definition}
	MetricDefinition *EvidentlyLaunchMetricMonitorsMetricDefinition `field:"required" json:"metricDefinition" yaml:"metricDefinition"`
}

