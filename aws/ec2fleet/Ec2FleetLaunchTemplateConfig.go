package ec2fleet


type Ec2FleetLaunchTemplateConfig struct {
	// launch_template_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ec2_fleet#launch_template_specification Ec2Fleet#launch_template_specification}
	LaunchTemplateSpecification *Ec2FleetLaunchTemplateConfigLaunchTemplateSpecification `field:"optional" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ec2_fleet#override Ec2Fleet#override}
	Override interface{} `field:"optional" json:"override" yaml:"override"`
}

