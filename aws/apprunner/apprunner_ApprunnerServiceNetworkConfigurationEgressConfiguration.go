package apprunner


type ApprunnerServiceNetworkConfigurationEgressConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service#egress_type ApprunnerService#egress_type}.
	EgressType *string `field:"optional" json:"egressType" yaml:"egressType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service#vpc_connector_arn ApprunnerService#vpc_connector_arn}.
	VpcConnectorArn *string `field:"optional" json:"vpcConnectorArn" yaml:"vpcConnectorArn"`
}

