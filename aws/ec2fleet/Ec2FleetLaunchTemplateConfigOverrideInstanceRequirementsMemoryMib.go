package ec2fleet


type Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryMib struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/ec2_fleet#min Ec2Fleet#min}.
	Min *float64 `field:"required" json:"min" yaml:"min"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/ec2_fleet#max Ec2Fleet#max}.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
}

