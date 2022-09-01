package wafv2


type Wafv2WebAclDefaultAction struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#allow Wafv2WebAcl#allow}
	Allow *Wafv2WebAclDefaultActionAllow `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#block Wafv2WebAcl#block}
	Block *Wafv2WebAclDefaultActionBlock `field:"optional" json:"block" yaml:"block"`
}

