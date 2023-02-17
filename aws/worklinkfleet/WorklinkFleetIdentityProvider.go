package worklinkfleet


type WorklinkFleetIdentityProvider struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/worklink_fleet#saml_metadata WorklinkFleet#saml_metadata}.
	SamlMetadata *string `field:"required" json:"samlMetadata" yaml:"samlMetadata"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/worklink_fleet#type WorklinkFleet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

