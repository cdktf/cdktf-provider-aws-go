package autoscalingpolicy


type AutoscalingPolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecification struct {
	// metric_data_queries block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/autoscaling_policy#metric_data_queries AutoscalingPolicy#metric_data_queries}
	MetricDataQueries interface{} `field:"required" json:"metricDataQueries" yaml:"metricDataQueries"`
}

