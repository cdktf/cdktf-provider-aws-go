package opsworkscustomlayer


type OpsworksCustomLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/opsworks_custom_layer#mount_point OpsworksCustomLayer#mount_point}.
	MountPoint *string `field:"required" json:"mountPoint" yaml:"mountPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/opsworks_custom_layer#number_of_disks OpsworksCustomLayer#number_of_disks}.
	NumberOfDisks *float64 `field:"required" json:"numberOfDisks" yaml:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/opsworks_custom_layer#size OpsworksCustomLayer#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/opsworks_custom_layer#encrypted OpsworksCustomLayer#encrypted}.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/opsworks_custom_layer#iops OpsworksCustomLayer#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/opsworks_custom_layer#raid_level OpsworksCustomLayer#raid_level}.
	RaidLevel *string `field:"optional" json:"raidLevel" yaml:"raidLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/opsworks_custom_layer#type OpsworksCustomLayer#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

