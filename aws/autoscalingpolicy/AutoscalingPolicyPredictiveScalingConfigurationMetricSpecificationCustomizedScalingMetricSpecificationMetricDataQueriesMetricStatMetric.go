package autoscalingpolicy


type AutoscalingPolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueriesMetricStatMetric struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/autoscaling_policy#metric_name AutoscalingPolicy#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/autoscaling_policy#namespace AutoscalingPolicy#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/autoscaling_policy#dimensions AutoscalingPolicy#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
}

