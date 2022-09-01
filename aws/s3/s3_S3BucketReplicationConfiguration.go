package s3


type S3BucketReplicationConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket#role S3Bucket#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket#rules S3Bucket#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

