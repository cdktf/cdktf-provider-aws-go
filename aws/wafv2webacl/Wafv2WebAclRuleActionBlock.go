package wafv2webacl


type Wafv2WebAclRuleActionBlock struct {
	// custom_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/wafv2_web_acl#custom_response Wafv2WebAcl#custom_response}
	CustomResponse *Wafv2WebAclRuleActionBlockCustomResponse `field:"optional" json:"customResponse" yaml:"customResponse"`
}

