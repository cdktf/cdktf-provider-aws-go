package autoscalinggrouptag


type AutoscalingGroupTagTag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/autoscaling_group_tag#key AutoscalingGroupTagA#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/autoscaling_group_tag#propagate_at_launch AutoscalingGroupTagA#propagate_at_launch}.
	PropagateAtLaunch interface{} `field:"required" json:"propagateAtLaunch" yaml:"propagateAtLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/autoscaling_group_tag#value AutoscalingGroupTagA#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

