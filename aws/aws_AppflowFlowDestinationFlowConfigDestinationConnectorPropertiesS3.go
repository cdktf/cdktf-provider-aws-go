// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#bucket_name AppflowFlow#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#bucket_prefix AppflowFlow#bucket_prefix}.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// s3_output_format_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#s3_output_format_config AppflowFlow#s3_output_format_config}
	S3OutputFormatConfig *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfig `field:"optional" json:"s3OutputFormatConfig" yaml:"s3OutputFormatConfig"`
}

