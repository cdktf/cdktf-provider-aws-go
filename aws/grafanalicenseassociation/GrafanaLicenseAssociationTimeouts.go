package grafanalicenseassociation


type GrafanaLicenseAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/grafana_license_association#create GrafanaLicenseAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/grafana_license_association#delete GrafanaLicenseAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

