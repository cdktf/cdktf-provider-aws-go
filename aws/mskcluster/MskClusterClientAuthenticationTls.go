package mskcluster


type MskClusterClientAuthenticationTls struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/msk_cluster#certificate_authority_arns MskCluster#certificate_authority_arns}.
	CertificateAuthorityArns *[]*string `field:"optional" json:"certificateAuthorityArns" yaml:"certificateAuthorityArns"`
}

