// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type CeCostCategoryRuleInheritedValue struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ce_cost_category#dimension_key CeCostCategory#dimension_key}.
	DimensionKey *string `field:"optional" json:"dimensionKey" yaml:"dimensionKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ce_cost_category#dimension_name CeCostCategory#dimension_name}.
	DimensionName *string `field:"optional" json:"dimensionName" yaml:"dimensionName"`
}

