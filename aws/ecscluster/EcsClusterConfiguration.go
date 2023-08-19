package ecscluster


type EcsClusterConfiguration struct {
	// execute_command_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ecs_cluster#execute_command_configuration EcsCluster#execute_command_configuration}
	ExecuteCommandConfiguration *EcsClusterConfigurationExecuteCommandConfiguration `field:"optional" json:"executeCommandConfiguration" yaml:"executeCommandConfiguration"`
}

