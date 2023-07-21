package autoscalinggroup


type AutoscalingGroupInitialLifecycleHook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/autoscaling_group#lifecycle_transition AutoscalingGroup#lifecycle_transition}.
	LifecycleTransition *string `field:"required" json:"lifecycleTransition" yaml:"lifecycleTransition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/autoscaling_group#name AutoscalingGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/autoscaling_group#default_result AutoscalingGroup#default_result}.
	DefaultResult *string `field:"optional" json:"defaultResult" yaml:"defaultResult"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/autoscaling_group#heartbeat_timeout AutoscalingGroup#heartbeat_timeout}.
	HeartbeatTimeout *float64 `field:"optional" json:"heartbeatTimeout" yaml:"heartbeatTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/autoscaling_group#notification_metadata AutoscalingGroup#notification_metadata}.
	NotificationMetadata *string `field:"optional" json:"notificationMetadata" yaml:"notificationMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/autoscaling_group#notification_target_arn AutoscalingGroup#notification_target_arn}.
	NotificationTargetArn *string `field:"optional" json:"notificationTargetArn" yaml:"notificationTargetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/autoscaling_group#role_arn AutoscalingGroup#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

