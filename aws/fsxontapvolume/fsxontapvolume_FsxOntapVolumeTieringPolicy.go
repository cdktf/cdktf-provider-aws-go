package fsxontapvolume


type FsxOntapVolumeTieringPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume#cooling_period FsxOntapVolume#cooling_period}.
	CoolingPeriod *float64 `field:"optional" json:"coolingPeriod" yaml:"coolingPeriod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume#name FsxOntapVolume#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

