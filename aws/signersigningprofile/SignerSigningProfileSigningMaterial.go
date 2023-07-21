package signersigningprofile


type SignerSigningProfileSigningMaterial struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/signer_signing_profile#certificate_arn SignerSigningProfile#certificate_arn}.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
}

