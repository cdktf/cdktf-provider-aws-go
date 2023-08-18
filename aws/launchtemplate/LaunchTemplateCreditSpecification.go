package launchtemplate


type LaunchTemplateCreditSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/launch_template#cpu_credits LaunchTemplate#cpu_credits}.
	CpuCredits *string `field:"optional" json:"cpuCredits" yaml:"cpuCredits"`
}

