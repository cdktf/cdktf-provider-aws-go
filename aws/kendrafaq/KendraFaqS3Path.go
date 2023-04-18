package kendrafaq


type KendraFaqS3Path struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/kendra_faq#bucket KendraFaq#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/kendra_faq#key KendraFaq#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
}

