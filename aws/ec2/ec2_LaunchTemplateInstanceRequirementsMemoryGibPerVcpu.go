package ec2


type LaunchTemplateInstanceRequirementsMemoryGibPerVcpu struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/launch_template#max LaunchTemplate#max}.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/launch_template#min LaunchTemplate#min}.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

