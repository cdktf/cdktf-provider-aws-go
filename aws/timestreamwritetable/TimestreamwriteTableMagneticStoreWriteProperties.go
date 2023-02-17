package timestreamwritetable


type TimestreamwriteTableMagneticStoreWriteProperties struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/timestreamwrite_table#enable_magnetic_store_writes TimestreamwriteTable#enable_magnetic_store_writes}.
	EnableMagneticStoreWrites interface{} `field:"optional" json:"enableMagneticStoreWrites" yaml:"enableMagneticStoreWrites"`
	// magnetic_store_rejected_data_location block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/timestreamwrite_table#magnetic_store_rejected_data_location TimestreamwriteTable#magnetic_store_rejected_data_location}
	MagneticStoreRejectedDataLocation *TimestreamwriteTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocation `field:"optional" json:"magneticStoreRejectedDataLocation" yaml:"magneticStoreRejectedDataLocation"`
}

