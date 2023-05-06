package grafanaroleassociation


type GrafanaRoleAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/grafana_role_association#create GrafanaRoleAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/grafana_role_association#delete GrafanaRoleAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

