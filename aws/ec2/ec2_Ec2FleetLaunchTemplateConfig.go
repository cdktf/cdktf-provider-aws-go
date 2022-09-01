package ec2


type Ec2FleetLaunchTemplateConfig struct {
	// launch_template_specification block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_fleet#launch_template_specification Ec2Fleet#launch_template_specification}
	LaunchTemplateSpecification *Ec2FleetLaunchTemplateConfigLaunchTemplateSpecification `field:"required" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// override block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_fleet#override Ec2Fleet#override}
	Override interface{} `field:"optional" json:"override" yaml:"override"`
}

