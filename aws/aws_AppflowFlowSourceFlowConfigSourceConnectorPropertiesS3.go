// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#bucket_name AppflowFlow#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#bucket_prefix AppflowFlow#bucket_prefix}.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// s3_input_format_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#s3_input_format_config AppflowFlow#s3_input_format_config}
	S3InputFormatConfig *AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3S3InputFormatConfig `field:"optional" json:"s3InputFormatConfig" yaml:"s3InputFormatConfig"`
}

