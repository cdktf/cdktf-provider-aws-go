package autoscalinggroup


type AutoscalingGroupInstanceRefresh struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/autoscaling_group#strategy AutoscalingGroup#strategy}.
	Strategy *string `field:"required" json:"strategy" yaml:"strategy"`
	// preferences block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/autoscaling_group#preferences AutoscalingGroup#preferences}
	Preferences *AutoscalingGroupInstanceRefreshPreferences `field:"optional" json:"preferences" yaml:"preferences"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/autoscaling_group#triggers AutoscalingGroup#triggers}.
	Triggers *[]*string `field:"optional" json:"triggers" yaml:"triggers"`
}

