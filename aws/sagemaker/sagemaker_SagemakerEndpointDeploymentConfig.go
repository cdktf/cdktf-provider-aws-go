package sagemaker


type SagemakerEndpointDeploymentConfig struct {
	// blue_green_update_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_endpoint#blue_green_update_policy SagemakerEndpoint#blue_green_update_policy}
	BlueGreenUpdatePolicy *SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicy `field:"required" json:"blueGreenUpdatePolicy" yaml:"blueGreenUpdatePolicy"`
	// auto_rollback_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_endpoint#auto_rollback_configuration SagemakerEndpoint#auto_rollback_configuration}
	AutoRollbackConfiguration *SagemakerEndpointDeploymentConfigAutoRollbackConfiguration `field:"optional" json:"autoRollbackConfiguration" yaml:"autoRollbackConfiguration"`
}

