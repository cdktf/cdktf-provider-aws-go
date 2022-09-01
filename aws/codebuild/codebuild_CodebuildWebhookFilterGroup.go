package codebuild


type CodebuildWebhookFilterGroup struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_webhook#filter CodebuildWebhook#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

