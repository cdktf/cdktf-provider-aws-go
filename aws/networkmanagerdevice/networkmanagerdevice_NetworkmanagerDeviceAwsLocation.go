package networkmanagerdevice


type NetworkmanagerDeviceAwsLocation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkmanager_device#subnet_arn NetworkmanagerDevice#subnet_arn}.
	SubnetArn *string `field:"optional" json:"subnetArn" yaml:"subnetArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkmanager_device#zone NetworkmanagerDevice#zone}.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

