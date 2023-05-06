package evidentlylaunch


type EvidentlyLaunchMetricMonitors struct {
	// metric_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/evidently_launch#metric_definition EvidentlyLaunch#metric_definition}
	MetricDefinition *EvidentlyLaunchMetricMonitorsMetricDefinition `field:"required" json:"metricDefinition" yaml:"metricDefinition"`
}

