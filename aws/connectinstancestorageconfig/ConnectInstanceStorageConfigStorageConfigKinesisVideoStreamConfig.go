package connectinstancestorageconfig


type ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfig struct {
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/connect_instance_storage_config#encryption_config ConnectInstanceStorageConfig#encryption_config}
	EncryptionConfig *ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigEncryptionConfig `field:"required" json:"encryptionConfig" yaml:"encryptionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/connect_instance_storage_config#prefix ConnectInstanceStorageConfig#prefix}.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/connect_instance_storage_config#retention_period_hours ConnectInstanceStorageConfig#retention_period_hours}.
	RetentionPeriodHours *float64 `field:"required" json:"retentionPeriodHours" yaml:"retentionPeriodHours"`
}

