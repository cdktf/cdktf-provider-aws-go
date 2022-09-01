package ec2


type DataAwsLaunchTemplateFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/launch_template#name DataAwsLaunchTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/launch_template#values DataAwsLaunchTemplate#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

