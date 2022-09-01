package timestreamwrite


type TimestreamwriteTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocation struct {
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/timestreamwrite_table#s3_configuration TimestreamwriteTable#s3_configuration}
	S3Configuration *TimestreamwriteTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocationS3Configuration `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

