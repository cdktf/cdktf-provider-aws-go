// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type CeCostCategoryRuleRuleOr struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ce_cost_category#cost_category CeCostCategory#cost_category}
	CostCategory *CeCostCategoryRuleRuleOrCostCategory `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ce_cost_category#dimension CeCostCategory#dimension}
	Dimension *CeCostCategoryRuleRuleOrDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ce_cost_category#tags CeCostCategory#tags}
	Tags *CeCostCategoryRuleRuleOrTags `field:"optional" json:"tags" yaml:"tags"`
}

