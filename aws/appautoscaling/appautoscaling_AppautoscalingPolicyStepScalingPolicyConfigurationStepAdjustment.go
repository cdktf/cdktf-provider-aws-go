package appautoscaling


type AppautoscalingPolicyStepScalingPolicyConfigurationStepAdjustment struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#scaling_adjustment AppautoscalingPolicy#scaling_adjustment}.
	ScalingAdjustment *float64 `field:"required" json:"scalingAdjustment" yaml:"scalingAdjustment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#metric_interval_lower_bound AppautoscalingPolicy#metric_interval_lower_bound}.
	MetricIntervalLowerBound *string `field:"optional" json:"metricIntervalLowerBound" yaml:"metricIntervalLowerBound"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#metric_interval_upper_bound AppautoscalingPolicy#metric_interval_upper_bound}.
	MetricIntervalUpperBound *string `field:"optional" json:"metricIntervalUpperBound" yaml:"metricIntervalUpperBound"`
}

