package mq


type MqBrokerEncryptionOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker#kms_key_id MqBroker#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mq_broker#use_aws_owned_key MqBroker#use_aws_owned_key}.
	UseAwsOwnedKey interface{} `field:"optional" json:"useAwsOwnedKey" yaml:"useAwsOwnedKey"`
}

