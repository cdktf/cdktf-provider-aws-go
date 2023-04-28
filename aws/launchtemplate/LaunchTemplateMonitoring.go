package launchtemplate


type LaunchTemplateMonitoring struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/launch_template#enabled LaunchTemplate#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

