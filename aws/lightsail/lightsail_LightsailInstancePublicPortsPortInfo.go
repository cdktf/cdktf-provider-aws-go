package lightsail


type LightsailInstancePublicPortsPortInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lightsail_instance_public_ports#from_port LightsailInstancePublicPorts#from_port}.
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lightsail_instance_public_ports#protocol LightsailInstancePublicPorts#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lightsail_instance_public_ports#to_port LightsailInstancePublicPorts#to_port}.
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lightsail_instance_public_ports#cidrs LightsailInstancePublicPorts#cidrs}.
	Cidrs *[]*string `field:"optional" json:"cidrs" yaml:"cidrs"`
}

