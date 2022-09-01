package connect


type ConnectInstanceStorageConfigStorageConfigKinesisStreamConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/connect_instance_storage_config#stream_arn ConnectInstanceStorageConfig#stream_arn}.
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
}

