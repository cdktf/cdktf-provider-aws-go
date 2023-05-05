package opsworksphpapplayer


type OpsworksPhpAppLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/opsworks_php_app_layer#mount_point OpsworksPhpAppLayer#mount_point}.
	MountPoint *string `field:"required" json:"mountPoint" yaml:"mountPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/opsworks_php_app_layer#number_of_disks OpsworksPhpAppLayer#number_of_disks}.
	NumberOfDisks *float64 `field:"required" json:"numberOfDisks" yaml:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/opsworks_php_app_layer#size OpsworksPhpAppLayer#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/opsworks_php_app_layer#encrypted OpsworksPhpAppLayer#encrypted}.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/opsworks_php_app_layer#iops OpsworksPhpAppLayer#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/opsworks_php_app_layer#raid_level OpsworksPhpAppLayer#raid_level}.
	RaidLevel *string `field:"optional" json:"raidLevel" yaml:"raidLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/opsworks_php_app_layer#type OpsworksPhpAppLayer#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

