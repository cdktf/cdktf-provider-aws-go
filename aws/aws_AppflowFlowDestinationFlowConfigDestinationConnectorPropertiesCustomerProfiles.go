// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesCustomerProfiles struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#domain_name AppflowFlow#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#object_type_name AppflowFlow#object_type_name}.
	ObjectTypeName *string `field:"optional" json:"objectTypeName" yaml:"objectTypeName"`
}

