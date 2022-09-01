// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KendraIndexDocumentMetadataConfigurationUpdates struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#name KendraIndex#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#type KendraIndex#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// relevance block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#relevance KendraIndex#relevance}
	Relevance *KendraIndexDocumentMetadataConfigurationUpdatesRelevance `field:"optional" json:"relevance" yaml:"relevance"`
	// search block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#search KendraIndex#search}
	Search *KendraIndexDocumentMetadataConfigurationUpdatesSearch `field:"optional" json:"search" yaml:"search"`
}

