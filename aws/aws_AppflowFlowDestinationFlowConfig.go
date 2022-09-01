// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AppflowFlowDestinationFlowConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#connector_type AppflowFlow#connector_type}.
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// destination_connector_properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#destination_connector_properties AppflowFlow#destination_connector_properties}
	DestinationConnectorProperties *AppflowFlowDestinationFlowConfigDestinationConnectorProperties `field:"required" json:"destinationConnectorProperties" yaml:"destinationConnectorProperties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#api_version AppflowFlow#api_version}.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#connector_profile_name AppflowFlow#connector_profile_name}.
	ConnectorProfileName *string `field:"optional" json:"connectorProfileName" yaml:"connectorProfileName"`
}

