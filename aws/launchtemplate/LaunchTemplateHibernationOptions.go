package launchtemplate


type LaunchTemplateHibernationOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/launch_template#configured LaunchTemplate#configured}.
	Configured interface{} `field:"required" json:"configured" yaml:"configured"`
}

