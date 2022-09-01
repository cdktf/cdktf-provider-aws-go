package cloudwatch


type CloudwatchMetricStreamIncludeFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudwatch_metric_stream#namespace CloudwatchMetricStream#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

