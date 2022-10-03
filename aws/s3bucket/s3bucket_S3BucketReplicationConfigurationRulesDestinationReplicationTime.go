package s3bucket


type S3BucketReplicationConfigurationRulesDestinationReplicationTime struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket#minutes S3Bucket#minutes}.
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket#status S3Bucket#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

