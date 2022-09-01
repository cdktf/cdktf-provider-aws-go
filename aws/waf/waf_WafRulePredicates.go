package waf


type WafRulePredicates struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/waf_rule#data_id WafRule#data_id}.
	DataId *string `field:"required" json:"dataId" yaml:"dataId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/waf_rule#negated WafRule#negated}.
	Negated interface{} `field:"required" json:"negated" yaml:"negated"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/waf_rule#type WafRule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

