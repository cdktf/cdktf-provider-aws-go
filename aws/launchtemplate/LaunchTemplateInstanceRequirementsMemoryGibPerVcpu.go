package launchtemplate


type LaunchTemplateInstanceRequirementsMemoryGibPerVcpu struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/launch_template#max LaunchTemplate#max}.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.12.0/docs/resources/launch_template#min LaunchTemplate#min}.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

