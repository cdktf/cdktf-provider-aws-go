package ec2


type AmiEphemeralBlockDevice struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ami#device_name Ami#device_name}.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ami#virtual_name Ami#virtual_name}.
	VirtualName *string `field:"required" json:"virtualName" yaml:"virtualName"`
}

