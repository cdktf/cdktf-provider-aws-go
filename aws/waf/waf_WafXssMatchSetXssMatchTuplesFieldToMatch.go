package waf


type WafXssMatchSetXssMatchTuplesFieldToMatch struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/waf_xss_match_set#type WafXssMatchSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/waf_xss_match_set#data WafXssMatchSet#data}.
	Data *string `field:"optional" json:"data" yaml:"data"`
}

