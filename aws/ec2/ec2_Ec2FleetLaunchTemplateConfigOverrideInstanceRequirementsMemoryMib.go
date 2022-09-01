package ec2


type Ec2FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryMib struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_fleet#min Ec2Fleet#min}.
	Min *float64 `field:"required" json:"min" yaml:"min"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_fleet#max Ec2Fleet#max}.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
}

