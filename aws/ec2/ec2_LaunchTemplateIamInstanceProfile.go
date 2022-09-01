package ec2


type LaunchTemplateIamInstanceProfile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/launch_template#arn LaunchTemplate#arn}.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/launch_template#name LaunchTemplate#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

