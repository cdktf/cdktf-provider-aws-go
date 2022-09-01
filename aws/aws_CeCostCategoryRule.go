// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type CeCostCategoryRule struct {
	// inherited_value block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ce_cost_category#inherited_value CeCostCategory#inherited_value}
	InheritedValue *CeCostCategoryRuleInheritedValue `field:"optional" json:"inheritedValue" yaml:"inheritedValue"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ce_cost_category#rule CeCostCategory#rule}
	Rule *CeCostCategoryRuleRule `field:"optional" json:"rule" yaml:"rule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ce_cost_category#type CeCostCategory#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ce_cost_category#value CeCostCategory#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

