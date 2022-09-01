package fsx


type FsxOpenzfsVolumeNfsExportsClientConfigurations struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#clients FsxOpenzfsVolume#clients}.
	Clients *string `field:"required" json:"clients" yaml:"clients"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#options FsxOpenzfsVolume#options}.
	Options *[]*string `field:"required" json:"options" yaml:"options"`
}

