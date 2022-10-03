package wafv2webacl


type Wafv2WebAclDefaultActionBlock struct {
	// custom_response block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#custom_response Wafv2WebAcl#custom_response}
	CustomResponse *Wafv2WebAclDefaultActionBlockCustomResponse `field:"optional" json:"customResponse" yaml:"customResponse"`
}

