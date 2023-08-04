package kendradatasource


type KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsTargetTargetDocumentAttributeValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/kendra_data_source#date_value KendraDataSource#date_value}.
	DateValue *string `field:"optional" json:"dateValue" yaml:"dateValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/kendra_data_source#long_value KendraDataSource#long_value}.
	LongValue *float64 `field:"optional" json:"longValue" yaml:"longValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/kendra_data_source#string_list_value KendraDataSource#string_list_value}.
	StringListValue *[]*string `field:"optional" json:"stringListValue" yaml:"stringListValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/kendra_data_source#string_value KendraDataSource#string_value}.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

