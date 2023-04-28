package sesreceiptrule


type SesReceiptRuleS3Action struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/ses_receipt_rule#bucket_name SesReceiptRule#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/ses_receipt_rule#position SesReceiptRule#position}.
	Position *float64 `field:"required" json:"position" yaml:"position"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/ses_receipt_rule#kms_key_arn SesReceiptRule#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/ses_receipt_rule#object_key_prefix SesReceiptRule#object_key_prefix}.
	ObjectKeyPrefix *string `field:"optional" json:"objectKeyPrefix" yaml:"objectKeyPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/ses_receipt_rule#topic_arn SesReceiptRule#topic_arn}.
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

