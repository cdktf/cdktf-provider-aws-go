package wafv2webacl


type Wafv2WebAclRuleActionCount struct {
	// custom_request_handling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/wafv2_web_acl#custom_request_handling Wafv2WebAcl#custom_request_handling}
	CustomRequestHandling *Wafv2WebAclRuleActionCountCustomRequestHandling `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}

