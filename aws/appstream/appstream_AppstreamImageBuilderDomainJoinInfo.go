package appstream


type AppstreamImageBuilderDomainJoinInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#directory_name AppstreamImageBuilder#directory_name}.
	DirectoryName *string `field:"optional" json:"directoryName" yaml:"directoryName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#organizational_unit_distinguished_name AppstreamImageBuilder#organizational_unit_distinguished_name}.
	OrganizationalUnitDistinguishedName *string `field:"optional" json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
}

