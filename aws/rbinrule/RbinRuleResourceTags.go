package rbinrule


type RbinRuleResourceTags struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/rbin_rule#resource_tag_key RbinRule#resource_tag_key}.
	ResourceTagKey *string `field:"required" json:"resourceTagKey" yaml:"resourceTagKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/rbin_rule#resource_tag_value RbinRule#resource_tag_value}.
	ResourceTagValue *string `field:"optional" json:"resourceTagValue" yaml:"resourceTagValue"`
}

