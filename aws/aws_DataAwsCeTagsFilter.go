// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type DataAwsCeTagsFilter struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ce_tags#and DataAwsCeTags#and}
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ce_tags#cost_category DataAwsCeTags#cost_category}
	CostCategory *DataAwsCeTagsFilterCostCategory `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ce_tags#dimension DataAwsCeTags#dimension}
	Dimension *DataAwsCeTagsFilterDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// not block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ce_tags#not DataAwsCeTags#not}
	Not *DataAwsCeTagsFilterNot `field:"optional" json:"not" yaml:"not"`
	// or block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ce_tags#or DataAwsCeTags#or}
	Or interface{} `field:"optional" json:"or" yaml:"or"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ce_tags#tags DataAwsCeTags#tags}
	Tags *DataAwsCeTagsFilterTags `field:"optional" json:"tags" yaml:"tags"`
}

