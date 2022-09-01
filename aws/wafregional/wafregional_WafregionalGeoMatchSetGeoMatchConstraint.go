package wafregional


type WafregionalGeoMatchSetGeoMatchConstraint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafregional_geo_match_set#type WafregionalGeoMatchSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafregional_geo_match_set#value WafregionalGeoMatchSet#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

