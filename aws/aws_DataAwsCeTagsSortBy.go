// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type DataAwsCeTagsSortBy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ce_tags#key DataAwsCeTags#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ce_tags#sort_order DataAwsCeTags#sort_order}.
	SortOrder *string `field:"optional" json:"sortOrder" yaml:"sortOrder"`
}

