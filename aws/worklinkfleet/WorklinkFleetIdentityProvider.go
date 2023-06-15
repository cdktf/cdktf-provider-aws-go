package worklinkfleet


type WorklinkFleetIdentityProvider struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/worklink_fleet#saml_metadata WorklinkFleet#saml_metadata}.
	SamlMetadata *string `field:"required" json:"samlMetadata" yaml:"samlMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/worklink_fleet#type WorklinkFleet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

