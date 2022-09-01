// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KendraFaqS3Path struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_faq#bucket KendraFaq#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_faq#key KendraFaq#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
}

