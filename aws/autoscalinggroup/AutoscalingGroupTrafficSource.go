package autoscalinggroup


type AutoscalingGroupTrafficSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/autoscaling_group#identifier AutoscalingGroup#identifier}.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/autoscaling_group#type AutoscalingGroup#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

