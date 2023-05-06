package launchtemplate


type LaunchTemplateHibernationOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/launch_template#configured LaunchTemplate#configured}.
	Configured interface{} `field:"required" json:"configured" yaml:"configured"`
}

