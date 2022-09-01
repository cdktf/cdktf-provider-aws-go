package wafregional


type WafregionalRateBasedRulePredicate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafregional_rate_based_rule#data_id WafregionalRateBasedRule#data_id}.
	DataId *string `field:"required" json:"dataId" yaml:"dataId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafregional_rate_based_rule#negated WafregionalRateBasedRule#negated}.
	Negated interface{} `field:"required" json:"negated" yaml:"negated"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafregional_rate_based_rule#type WafregionalRateBasedRule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

