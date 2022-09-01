package ec2


type InstanceNetworkInterface struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/instance#device_index Instance#device_index}.
	DeviceIndex *float64 `field:"required" json:"deviceIndex" yaml:"deviceIndex"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/instance#network_interface_id Instance#network_interface_id}.
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/instance#delete_on_termination Instance#delete_on_termination}.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/instance#network_card_index Instance#network_card_index}.
	NetworkCardIndex *float64 `field:"optional" json:"networkCardIndex" yaml:"networkCardIndex"`
}

