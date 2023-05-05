package autoscalinggroup


type AutoscalingGroupInstanceRefreshPreferences struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/autoscaling_group#auto_rollback AutoscalingGroup#auto_rollback}.
	AutoRollback interface{} `field:"optional" json:"autoRollback" yaml:"autoRollback"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/autoscaling_group#checkpoint_delay AutoscalingGroup#checkpoint_delay}.
	CheckpointDelay *string `field:"optional" json:"checkpointDelay" yaml:"checkpointDelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/autoscaling_group#checkpoint_percentages AutoscalingGroup#checkpoint_percentages}.
	CheckpointPercentages *[]*float64 `field:"optional" json:"checkpointPercentages" yaml:"checkpointPercentages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/autoscaling_group#instance_warmup AutoscalingGroup#instance_warmup}.
	InstanceWarmup *string `field:"optional" json:"instanceWarmup" yaml:"instanceWarmup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/autoscaling_group#min_healthy_percentage AutoscalingGroup#min_healthy_percentage}.
	MinHealthyPercentage *float64 `field:"optional" json:"minHealthyPercentage" yaml:"minHealthyPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/autoscaling_group#skip_matching AutoscalingGroup#skip_matching}.
	SkipMatching interface{} `field:"optional" json:"skipMatching" yaml:"skipMatching"`
}

