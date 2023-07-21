package sagemakerendpoint


type SagemakerEndpointDeploymentConfig struct {
	// blue_green_update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/sagemaker_endpoint#blue_green_update_policy SagemakerEndpoint#blue_green_update_policy}
	BlueGreenUpdatePolicy *SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicy `field:"required" json:"blueGreenUpdatePolicy" yaml:"blueGreenUpdatePolicy"`
	// auto_rollback_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/sagemaker_endpoint#auto_rollback_configuration SagemakerEndpoint#auto_rollback_configuration}
	AutoRollbackConfiguration *SagemakerEndpointDeploymentConfigAutoRollbackConfiguration `field:"optional" json:"autoRollbackConfiguration" yaml:"autoRollbackConfiguration"`
}

