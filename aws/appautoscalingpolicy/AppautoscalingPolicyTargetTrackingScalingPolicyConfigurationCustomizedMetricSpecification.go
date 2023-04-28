package appautoscalingpolicy


type AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appautoscaling_policy#metric_name AppautoscalingPolicy#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appautoscaling_policy#namespace AppautoscalingPolicy#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appautoscaling_policy#statistic AppautoscalingPolicy#statistic}.
	Statistic *string `field:"required" json:"statistic" yaml:"statistic"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appautoscaling_policy#dimensions AppautoscalingPolicy#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/appautoscaling_policy#unit AppautoscalingPolicy#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

