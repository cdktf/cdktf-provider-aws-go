package signersigningjob


type SignerSigningJobSourceS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/signer_signing_job#bucket SignerSigningJob#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/signer_signing_job#key SignerSigningJob#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/signer_signing_job#version SignerSigningJob#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
}

