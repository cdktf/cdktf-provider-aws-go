// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AppflowFlowSourceFlowConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#connector_type AppflowFlow#connector_type}.
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// source_connector_properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#source_connector_properties AppflowFlow#source_connector_properties}
	SourceConnectorProperties *AppflowFlowSourceFlowConfigSourceConnectorProperties `field:"required" json:"sourceConnectorProperties" yaml:"sourceConnectorProperties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#api_version AppflowFlow#api_version}.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#connector_profile_name AppflowFlow#connector_profile_name}.
	ConnectorProfileName *string `field:"optional" json:"connectorProfileName" yaml:"connectorProfileName"`
	// incremental_pull_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#incremental_pull_config AppflowFlow#incremental_pull_config}
	IncrementalPullConfig *AppflowFlowSourceFlowConfigIncrementalPullConfig `field:"optional" json:"incrementalPullConfig" yaml:"incrementalPullConfig"`
}

