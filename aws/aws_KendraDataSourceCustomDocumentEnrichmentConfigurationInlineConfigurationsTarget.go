// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsTarget struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_data_source#target_document_attribute_key KendraDataSource#target_document_attribute_key}.
	TargetDocumentAttributeKey *string `field:"optional" json:"targetDocumentAttributeKey" yaml:"targetDocumentAttributeKey"`
	// target_document_attribute_value block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_data_source#target_document_attribute_value KendraDataSource#target_document_attribute_value}
	TargetDocumentAttributeValue *KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsTargetTargetDocumentAttributeValue `field:"optional" json:"targetDocumentAttributeValue" yaml:"targetDocumentAttributeValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_data_source#target_document_attribute_value_deletion KendraDataSource#target_document_attribute_value_deletion}.
	TargetDocumentAttributeValueDeletion interface{} `field:"optional" json:"targetDocumentAttributeValueDeletion" yaml:"targetDocumentAttributeValueDeletion"`
}

