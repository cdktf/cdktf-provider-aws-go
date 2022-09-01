package ecs


type EcsClusterConfiguration struct {
	// execute_command_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_cluster#execute_command_configuration EcsCluster#execute_command_configuration}
	ExecuteCommandConfiguration *EcsClusterConfigurationExecuteCommandConfiguration `field:"optional" json:"executeCommandConfiguration" yaml:"executeCommandConfiguration"`
}

