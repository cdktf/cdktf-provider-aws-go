package wafregional


type WafregionalSqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatch struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafregional_sql_injection_match_set#type WafregionalSqlInjectionMatchSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafregional_sql_injection_match_set#data WafregionalSqlInjectionMatchSet#data}.
	Data *string `field:"optional" json:"data" yaml:"data"`
}

