// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type DataAwsCeTagsFilterNot struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ce_tags#cost_category DataAwsCeTags#cost_category}
	CostCategory *DataAwsCeTagsFilterNotCostCategory `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ce_tags#dimension DataAwsCeTags#dimension}
	Dimension *DataAwsCeTagsFilterNotDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ce_tags#tags DataAwsCeTags#tags}
	Tags *DataAwsCeTagsFilterNotTags `field:"optional" json:"tags" yaml:"tags"`
}

