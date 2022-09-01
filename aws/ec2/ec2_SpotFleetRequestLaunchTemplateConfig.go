package ec2


type SpotFleetRequestLaunchTemplateConfig struct {
	// launch_template_specification block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/spot_fleet_request#launch_template_specification SpotFleetRequest#launch_template_specification}
	LaunchTemplateSpecification *SpotFleetRequestLaunchTemplateConfigLaunchTemplateSpecification `field:"required" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// overrides block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/spot_fleet_request#overrides SpotFleetRequest#overrides}
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
}

