package acmcertificate


type AcmCertificateOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/acm_certificate#certificate_transparency_logging_preference AcmCertificate#certificate_transparency_logging_preference}.
	CertificateTransparencyLoggingPreference *string `field:"optional" json:"certificateTransparencyLoggingPreference" yaml:"certificateTransparencyLoggingPreference"`
}

