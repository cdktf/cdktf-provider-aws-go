package ec2


type Ec2FleetLaunchTemplateConfigLaunchTemplateSpecification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_fleet#version Ec2Fleet#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_fleet#launch_template_id Ec2Fleet#launch_template_id}.
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_fleet#launch_template_name Ec2Fleet#launch_template_name}.
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
}

