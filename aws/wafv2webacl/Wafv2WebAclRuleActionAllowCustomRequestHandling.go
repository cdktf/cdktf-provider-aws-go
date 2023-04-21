package wafv2webacl


type Wafv2WebAclRuleActionAllowCustomRequestHandling struct {
	// insert_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/wafv2_web_acl#insert_header Wafv2WebAcl#insert_header}
	InsertHeader interface{} `field:"required" json:"insertHeader" yaml:"insertHeader"`
}

