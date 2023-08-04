package rbinrule


type RbinRuleResourceTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/rbin_rule#resource_tag_key RbinRule#resource_tag_key}.
	ResourceTagKey *string `field:"required" json:"resourceTagKey" yaml:"resourceTagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/rbin_rule#resource_tag_value RbinRule#resource_tag_value}.
	ResourceTagValue *string `field:"optional" json:"resourceTagValue" yaml:"resourceTagValue"`
}

