package mskconnectcustomplugin


type MskconnectCustomPluginLocationS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/mskconnect_custom_plugin#bucket_arn MskconnectCustomPlugin#bucket_arn}.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/mskconnect_custom_plugin#file_key MskconnectCustomPlugin#file_key}.
	FileKey *string `field:"required" json:"fileKey" yaml:"fileKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/mskconnect_custom_plugin#object_version MskconnectCustomPlugin#object_version}.
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

