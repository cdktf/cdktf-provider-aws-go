package mskcluster


type MskClusterClientAuthenticationTls struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/msk_cluster#certificate_authority_arns MskCluster#certificate_authority_arns}.
	CertificateAuthorityArns *[]*string `field:"optional" json:"certificateAuthorityArns" yaml:"certificateAuthorityArns"`
}

