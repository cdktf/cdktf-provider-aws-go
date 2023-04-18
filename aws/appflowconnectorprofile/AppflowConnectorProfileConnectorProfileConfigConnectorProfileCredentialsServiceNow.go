package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsServiceNow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/appflow_connector_profile#password AppflowConnectorProfile#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/appflow_connector_profile#username AppflowConnectorProfile#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

