package sagemaker


type SagemakerEndpointDeploymentConfigAutoRollbackConfiguration struct {
	// alarms block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_endpoint#alarms SagemakerEndpoint#alarms}
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
}

