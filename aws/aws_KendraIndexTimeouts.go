// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KendraIndexTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#create KendraIndex#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#delete KendraIndex#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_index#update KendraIndex#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

