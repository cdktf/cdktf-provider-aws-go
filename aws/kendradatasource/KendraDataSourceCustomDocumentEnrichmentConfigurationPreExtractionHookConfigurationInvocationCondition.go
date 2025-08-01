// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kendradatasource


type KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kendra_data_source#condition_document_attribute_key KendraDataSource#condition_document_attribute_key}.
	ConditionDocumentAttributeKey *string `field:"required" json:"conditionDocumentAttributeKey" yaml:"conditionDocumentAttributeKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kendra_data_source#operator KendraDataSource#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// condition_on_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kendra_data_source#condition_on_value KendraDataSource#condition_on_value}
	ConditionOnValue *KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValue `field:"optional" json:"conditionOnValue" yaml:"conditionOnValue"`
}

