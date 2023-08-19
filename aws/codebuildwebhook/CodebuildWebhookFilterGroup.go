package codebuildwebhook


type CodebuildWebhookFilterGroup struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/codebuild_webhook#filter CodebuildWebhook#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

