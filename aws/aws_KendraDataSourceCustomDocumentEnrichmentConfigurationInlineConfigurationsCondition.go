// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_data_source#condition_document_attribute_key KendraDataSource#condition_document_attribute_key}.
	ConditionDocumentAttributeKey *string `field:"required" json:"conditionDocumentAttributeKey" yaml:"conditionDocumentAttributeKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_data_source#operator KendraDataSource#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// condition_on_value block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_data_source#condition_on_value KendraDataSource#condition_on_value}
	ConditionOnValue *KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsConditionConditionOnValue `field:"optional" json:"conditionOnValue" yaml:"conditionOnValue"`
}

