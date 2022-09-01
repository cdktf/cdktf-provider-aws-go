// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KendraThesaurusTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_thesaurus#create KendraThesaurus#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_thesaurus#delete KendraThesaurus#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_thesaurus#update KendraThesaurus#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

