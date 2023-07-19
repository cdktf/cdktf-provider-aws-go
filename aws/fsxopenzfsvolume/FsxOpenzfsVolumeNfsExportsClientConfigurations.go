package fsxopenzfsvolume


type FsxOpenzfsVolumeNfsExportsClientConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/fsx_openzfs_volume#clients FsxOpenzfsVolume#clients}.
	Clients *string `field:"required" json:"clients" yaml:"clients"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/fsx_openzfs_volume#options FsxOpenzfsVolume#options}.
	Options *[]*string `field:"required" json:"options" yaml:"options"`
}

