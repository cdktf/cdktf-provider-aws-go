package autoscalinggroup


type AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverride struct {
	// instance_requirements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/autoscaling_group#instance_requirements AutoscalingGroup#instance_requirements}
	InstanceRequirements *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirements `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/autoscaling_group#instance_type AutoscalingGroup#instance_type}.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// launch_template_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/autoscaling_group#launch_template_specification AutoscalingGroup#launch_template_specification}
	LaunchTemplateSpecification *AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecification `field:"optional" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/autoscaling_group#weighted_capacity AutoscalingGroup#weighted_capacity}.
	WeightedCapacity *string `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}

