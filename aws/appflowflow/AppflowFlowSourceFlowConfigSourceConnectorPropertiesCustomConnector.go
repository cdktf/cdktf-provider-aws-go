package appflowflow


type AppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnector struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/appflow_flow#entity_name AppflowFlow#entity_name}.
	EntityName *string `field:"required" json:"entityName" yaml:"entityName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/appflow_flow#custom_properties AppflowFlow#custom_properties}.
	CustomProperties *map[string]*string `field:"optional" json:"customProperties" yaml:"customProperties"`
}

