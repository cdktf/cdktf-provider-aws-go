// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KendraExperienceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_experience#create KendraExperience#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_experience#delete KendraExperience#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_experience#update KendraExperience#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

