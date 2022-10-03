package rolesanywheretrustanchor


type RolesanywhereTrustAnchorSource struct {
	// source_data block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/rolesanywhere_trust_anchor#source_data RolesanywhereTrustAnchor#source_data}
	SourceData *RolesanywhereTrustAnchorSourceSourceData `field:"required" json:"sourceData" yaml:"sourceData"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/rolesanywhere_trust_anchor#source_type RolesanywhereTrustAnchor#source_type}.
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
}

