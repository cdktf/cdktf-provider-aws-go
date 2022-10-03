package wafv2webacl


type Wafv2WebAclRuleActionAllow struct {
	// custom_request_handling block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#custom_request_handling Wafv2WebAcl#custom_request_handling}
	CustomRequestHandling *Wafv2WebAclRuleActionAllowCustomRequestHandling `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}

