package waf


type WafGeoMatchSetGeoMatchConstraint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/waf_geo_match_set#type WafGeoMatchSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/waf_geo_match_set#value WafGeoMatchSet#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

