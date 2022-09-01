// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KendraExperienceConfiguration struct {
	// content_source_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_experience#content_source_configuration KendraExperience#content_source_configuration}
	ContentSourceConfiguration *KendraExperienceConfigurationContentSourceConfiguration `field:"optional" json:"contentSourceConfiguration" yaml:"contentSourceConfiguration"`
	// user_identity_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_experience#user_identity_configuration KendraExperience#user_identity_configuration}
	UserIdentityConfiguration *KendraExperienceConfigurationUserIdentityConfiguration `field:"optional" json:"userIdentityConfiguration" yaml:"userIdentityConfiguration"`
}

