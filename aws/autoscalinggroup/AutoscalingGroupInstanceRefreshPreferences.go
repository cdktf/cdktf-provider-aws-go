// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupInstanceRefreshPreferences struct {
	// alarm_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/autoscaling_group#alarm_specification AutoscalingGroup#alarm_specification}
	AlarmSpecification *AutoscalingGroupInstanceRefreshPreferencesAlarmSpecification `field:"optional" json:"alarmSpecification" yaml:"alarmSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/autoscaling_group#auto_rollback AutoscalingGroup#auto_rollback}.
	AutoRollback interface{} `field:"optional" json:"autoRollback" yaml:"autoRollback"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/autoscaling_group#checkpoint_delay AutoscalingGroup#checkpoint_delay}.
	CheckpointDelay *string `field:"optional" json:"checkpointDelay" yaml:"checkpointDelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/autoscaling_group#checkpoint_percentages AutoscalingGroup#checkpoint_percentages}.
	CheckpointPercentages *[]*float64 `field:"optional" json:"checkpointPercentages" yaml:"checkpointPercentages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/autoscaling_group#instance_warmup AutoscalingGroup#instance_warmup}.
	InstanceWarmup *string `field:"optional" json:"instanceWarmup" yaml:"instanceWarmup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/autoscaling_group#max_healthy_percentage AutoscalingGroup#max_healthy_percentage}.
	MaxHealthyPercentage *float64 `field:"optional" json:"maxHealthyPercentage" yaml:"maxHealthyPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/autoscaling_group#min_healthy_percentage AutoscalingGroup#min_healthy_percentage}.
	MinHealthyPercentage *float64 `field:"optional" json:"minHealthyPercentage" yaml:"minHealthyPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/autoscaling_group#scale_in_protected_instances AutoscalingGroup#scale_in_protected_instances}.
	ScaleInProtectedInstances *string `field:"optional" json:"scaleInProtectedInstances" yaml:"scaleInProtectedInstances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/autoscaling_group#skip_matching AutoscalingGroup#skip_matching}.
	SkipMatching interface{} `field:"optional" json:"skipMatching" yaml:"skipMatching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/autoscaling_group#standby_instances AutoscalingGroup#standby_instances}.
	StandbyInstances *string `field:"optional" json:"standbyInstances" yaml:"standbyInstances"`
}

