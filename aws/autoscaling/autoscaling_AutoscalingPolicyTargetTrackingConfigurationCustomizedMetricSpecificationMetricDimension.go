package autoscaling


type AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/autoscaling_policy#name AutoscalingPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/autoscaling_policy#value AutoscalingPolicy#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

