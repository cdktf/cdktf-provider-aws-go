package ec2


type LaunchTemplateMonitoring struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/launch_template#enabled LaunchTemplate#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

