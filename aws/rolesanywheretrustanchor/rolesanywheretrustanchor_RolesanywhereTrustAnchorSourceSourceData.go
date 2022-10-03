package rolesanywheretrustanchor


type RolesanywhereTrustAnchorSourceSourceData struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/rolesanywhere_trust_anchor#acm_pca_arn RolesanywhereTrustAnchor#acm_pca_arn}.
	AcmPcaArn *string `field:"optional" json:"acmPcaArn" yaml:"acmPcaArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/rolesanywhere_trust_anchor#x509_certificate_data RolesanywhereTrustAnchor#x509_certificate_data}.
	X509CertificateData *string `field:"optional" json:"x509CertificateData" yaml:"x509CertificateData"`
}

