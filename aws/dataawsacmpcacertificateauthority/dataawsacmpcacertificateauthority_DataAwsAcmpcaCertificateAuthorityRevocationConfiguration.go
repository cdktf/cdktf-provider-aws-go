package dataawsacmpcacertificateauthority


type DataAwsAcmpcaCertificateAuthorityRevocationConfiguration struct {
	// crl_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/acmpca_certificate_authority#crl_configuration DataAwsAcmpcaCertificateAuthority#crl_configuration}
	CrlConfiguration interface{} `field:"optional" json:"crlConfiguration" yaml:"crlConfiguration"`
	// ocsp_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/acmpca_certificate_authority#ocsp_configuration DataAwsAcmpcaCertificateAuthority#ocsp_configuration}
	OcspConfiguration interface{} `field:"optional" json:"ocspConfiguration" yaml:"ocspConfiguration"`
}

