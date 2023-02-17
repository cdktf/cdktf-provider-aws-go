package apprunnerservice


type ApprunnerServiceEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service#kms_key ApprunnerService#kms_key}.
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
}

