// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type CeCostCategoryRuleRuleAnd struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ce_cost_category#cost_category CeCostCategory#cost_category}
	CostCategory *CeCostCategoryRuleRuleAndCostCategory `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ce_cost_category#dimension CeCostCategory#dimension}
	Dimension *CeCostCategoryRuleRuleAndDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ce_cost_category#tags CeCostCategory#tags}
	Tags *CeCostCategoryRuleRuleAndTags `field:"optional" json:"tags" yaml:"tags"`
}

