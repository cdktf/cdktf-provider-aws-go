package wafv2webaclloggingconfiguration


type Wafv2WebAclLoggingConfigurationRedactedFields struct {
	// all_query_arguments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/wafv2_web_acl_logging_configuration#all_query_arguments Wafv2WebAclLoggingConfiguration#all_query_arguments}
	AllQueryArguments *Wafv2WebAclLoggingConfigurationRedactedFieldsAllQueryArguments `field:"optional" json:"allQueryArguments" yaml:"allQueryArguments"`
	// body block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/wafv2_web_acl_logging_configuration#body Wafv2WebAclLoggingConfiguration#body}
	Body *Wafv2WebAclLoggingConfigurationRedactedFieldsBody `field:"optional" json:"body" yaml:"body"`
	// method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/wafv2_web_acl_logging_configuration#method Wafv2WebAclLoggingConfiguration#method}
	Method *Wafv2WebAclLoggingConfigurationRedactedFieldsMethod `field:"optional" json:"method" yaml:"method"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/wafv2_web_acl_logging_configuration#query_string Wafv2WebAclLoggingConfiguration#query_string}
	QueryString *Wafv2WebAclLoggingConfigurationRedactedFieldsQueryString `field:"optional" json:"queryString" yaml:"queryString"`
	// single_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/wafv2_web_acl_logging_configuration#single_header Wafv2WebAclLoggingConfiguration#single_header}
	SingleHeader *Wafv2WebAclLoggingConfigurationRedactedFieldsSingleHeader `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// single_query_argument block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/wafv2_web_acl_logging_configuration#single_query_argument Wafv2WebAclLoggingConfiguration#single_query_argument}
	SingleQueryArgument *Wafv2WebAclLoggingConfigurationRedactedFieldsSingleQueryArgument `field:"optional" json:"singleQueryArgument" yaml:"singleQueryArgument"`
	// uri_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/wafv2_web_acl_logging_configuration#uri_path Wafv2WebAclLoggingConfiguration#uri_path}
	UriPath *Wafv2WebAclLoggingConfigurationRedactedFieldsUriPath `field:"optional" json:"uriPath" yaml:"uriPath"`
}

