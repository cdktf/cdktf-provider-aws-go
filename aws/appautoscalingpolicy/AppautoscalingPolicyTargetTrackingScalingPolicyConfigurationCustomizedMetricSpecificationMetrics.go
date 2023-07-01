package appautoscalingpolicy


type AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetrics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/appautoscaling_policy#id AppautoscalingPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/appautoscaling_policy#expression AppautoscalingPolicy#expression}.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/appautoscaling_policy#label AppautoscalingPolicy#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// metric_stat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/appautoscaling_policy#metric_stat AppautoscalingPolicy#metric_stat}
	MetricStat *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStat `field:"optional" json:"metricStat" yaml:"metricStat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/appautoscaling_policy#return_data AppautoscalingPolicy#return_data}.
	ReturnData interface{} `field:"optional" json:"returnData" yaml:"returnData"`
}

