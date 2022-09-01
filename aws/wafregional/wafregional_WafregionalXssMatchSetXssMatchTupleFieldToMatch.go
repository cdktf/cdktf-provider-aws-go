package wafregional


type WafregionalXssMatchSetXssMatchTupleFieldToMatch struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafregional_xss_match_set#type WafregionalXssMatchSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafregional_xss_match_set#data WafregionalXssMatchSet#data}.
	Data *string `field:"optional" json:"data" yaml:"data"`
}

