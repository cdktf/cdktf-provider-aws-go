package ivschatloggingconfiguration


type IvschatLoggingConfigurationDestinationConfigurationS3 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ivschat_logging_configuration#bucket_name IvschatLoggingConfiguration#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
}

