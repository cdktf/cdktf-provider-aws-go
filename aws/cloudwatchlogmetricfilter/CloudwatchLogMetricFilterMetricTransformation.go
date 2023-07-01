package cloudwatchlogmetricfilter


type CloudwatchLogMetricFilterMetricTransformation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/cloudwatch_log_metric_filter#name CloudwatchLogMetricFilter#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/cloudwatch_log_metric_filter#namespace CloudwatchLogMetricFilter#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/cloudwatch_log_metric_filter#value CloudwatchLogMetricFilter#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/cloudwatch_log_metric_filter#default_value CloudwatchLogMetricFilter#default_value}.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/cloudwatch_log_metric_filter#dimensions CloudwatchLogMetricFilter#dimensions}.
	Dimensions *map[string]*string `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/cloudwatch_log_metric_filter#unit CloudwatchLogMetricFilter#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

