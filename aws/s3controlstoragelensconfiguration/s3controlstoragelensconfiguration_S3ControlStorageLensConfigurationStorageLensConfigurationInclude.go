package s3controlstoragelensconfiguration


type S3ControlStorageLensConfigurationStorageLensConfigurationInclude struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3control_storage_lens_configuration#buckets S3ControlStorageLensConfiguration#buckets}.
	Buckets *[]*string `field:"optional" json:"buckets" yaml:"buckets"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3control_storage_lens_configuration#regions S3ControlStorageLensConfiguration#regions}.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

