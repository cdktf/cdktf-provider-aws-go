package datasync


type DatasyncLocationS3S3Config struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/datasync_location_s3#bucket_access_role_arn DatasyncLocationS3#bucket_access_role_arn}.
	BucketAccessRoleArn *string `field:"required" json:"bucketAccessRoleArn" yaml:"bucketAccessRoleArn"`
}

