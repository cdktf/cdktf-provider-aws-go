package connectinstancestorageconfig


type ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfigEncryptionConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/connect_instance_storage_config#encryption_type ConnectInstanceStorageConfig#encryption_type}.
	EncryptionType *string `field:"required" json:"encryptionType" yaml:"encryptionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/connect_instance_storage_config#key_id ConnectInstanceStorageConfig#key_id}.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
}

