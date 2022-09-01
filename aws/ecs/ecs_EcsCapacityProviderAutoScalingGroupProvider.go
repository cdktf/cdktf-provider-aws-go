package ecs


type EcsCapacityProviderAutoScalingGroupProvider struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_capacity_provider#auto_scaling_group_arn EcsCapacityProvider#auto_scaling_group_arn}.
	AutoScalingGroupArn *string `field:"required" json:"autoScalingGroupArn" yaml:"autoScalingGroupArn"`
	// managed_scaling block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_capacity_provider#managed_scaling EcsCapacityProvider#managed_scaling}
	ManagedScaling *EcsCapacityProviderAutoScalingGroupProviderManagedScaling `field:"optional" json:"managedScaling" yaml:"managedScaling"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_capacity_provider#managed_termination_protection EcsCapacityProvider#managed_termination_protection}.
	ManagedTerminationProtection *string `field:"optional" json:"managedTerminationProtection" yaml:"managedTerminationProtection"`
}

