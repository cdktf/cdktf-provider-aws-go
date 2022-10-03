package acmpcacertificateauthority


type AcmpcaCertificateAuthorityRevocationConfiguration struct {
	// crl_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#crl_configuration AcmpcaCertificateAuthority#crl_configuration}
	CrlConfiguration *AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration `field:"optional" json:"crlConfiguration" yaml:"crlConfiguration"`
	// ocsp_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#ocsp_configuration AcmpcaCertificateAuthority#ocsp_configuration}
	OcspConfiguration *AcmpcaCertificateAuthorityRevocationConfigurationOcspConfiguration `field:"optional" json:"ocspConfiguration" yaml:"ocspConfiguration"`
}

