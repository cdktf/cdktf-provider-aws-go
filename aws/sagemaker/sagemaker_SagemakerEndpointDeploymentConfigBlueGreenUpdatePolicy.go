package sagemaker


type SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicy struct {
	// traffic_routing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_endpoint#traffic_routing_configuration SagemakerEndpoint#traffic_routing_configuration}
	TrafficRoutingConfiguration *SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfiguration `field:"required" json:"trafficRoutingConfiguration" yaml:"trafficRoutingConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_endpoint#maximum_execution_timeout_in_seconds SagemakerEndpoint#maximum_execution_timeout_in_seconds}.
	MaximumExecutionTimeoutInSeconds *float64 `field:"optional" json:"maximumExecutionTimeoutInSeconds" yaml:"maximumExecutionTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_endpoint#termination_wait_in_seconds SagemakerEndpoint#termination_wait_in_seconds}.
	TerminationWaitInSeconds *float64 `field:"optional" json:"terminationWaitInSeconds" yaml:"terminationWaitInSeconds"`
}

