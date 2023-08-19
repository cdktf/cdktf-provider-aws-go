package appstreamfleet


type AppstreamFleetDomainJoinInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/appstream_fleet#directory_name AppstreamFleet#directory_name}.
	DirectoryName *string `field:"optional" json:"directoryName" yaml:"directoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/appstream_fleet#organizational_unit_distinguished_name AppstreamFleet#organizational_unit_distinguished_name}.
	OrganizationalUnitDistinguishedName *string `field:"optional" json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
}

