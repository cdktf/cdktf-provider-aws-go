package launchtemplate


type LaunchTemplateIamInstanceProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/launch_template#arn LaunchTemplate#arn}.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/launch_template#name LaunchTemplate#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

