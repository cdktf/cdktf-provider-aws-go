package wafv2webacl


type Wafv2WebAclRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatch struct {
	// all_query_arguments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/wafv2_web_acl#all_query_arguments Wafv2WebAcl#all_query_arguments}
	AllQueryArguments *Wafv2WebAclRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments `field:"optional" json:"allQueryArguments" yaml:"allQueryArguments"`
	// body block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/wafv2_web_acl#body Wafv2WebAcl#body}
	Body *Wafv2WebAclRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchBody `field:"optional" json:"body" yaml:"body"`
	// cookies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/wafv2_web_acl#cookies Wafv2WebAcl#cookies}
	Cookies *Wafv2WebAclRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchCookies `field:"optional" json:"cookies" yaml:"cookies"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/wafv2_web_acl#headers Wafv2WebAcl#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// json_body block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/wafv2_web_acl#json_body Wafv2WebAcl#json_body}
	JsonBody *Wafv2WebAclRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchJsonBody `field:"optional" json:"jsonBody" yaml:"jsonBody"`
	// method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/wafv2_web_acl#method Wafv2WebAcl#method}
	Method *Wafv2WebAclRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchMethod `field:"optional" json:"method" yaml:"method"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/wafv2_web_acl#query_string Wafv2WebAcl#query_string}
	QueryString *Wafv2WebAclRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchQueryString `field:"optional" json:"queryString" yaml:"queryString"`
	// single_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/wafv2_web_acl#single_header Wafv2WebAcl#single_header}
	SingleHeader *Wafv2WebAclRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleHeader `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// single_query_argument block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/wafv2_web_acl#single_query_argument Wafv2WebAcl#single_query_argument}
	SingleQueryArgument *Wafv2WebAclRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument `field:"optional" json:"singleQueryArgument" yaml:"singleQueryArgument"`
	// uri_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/wafv2_web_acl#uri_path Wafv2WebAcl#uri_path}
	UriPath *Wafv2WebAclRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchUriPath `field:"optional" json:"uriPath" yaml:"uriPath"`
}

