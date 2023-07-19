package configorganizationcustomrule


type ConfigOrganizationCustomRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/config_organization_custom_rule#create ConfigOrganizationCustomRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/config_organization_custom_rule#delete ConfigOrganizationCustomRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/config_organization_custom_rule#update ConfigOrganizationCustomRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

