// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type CeCostCategorySplitChargeRuleParameter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ce_cost_category#type CeCostCategory#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ce_cost_category#values CeCostCategory#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

