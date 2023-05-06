package apprunnerservice


type ApprunnerServiceEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/apprunner_service#kms_key ApprunnerService#kms_key}.
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
}

