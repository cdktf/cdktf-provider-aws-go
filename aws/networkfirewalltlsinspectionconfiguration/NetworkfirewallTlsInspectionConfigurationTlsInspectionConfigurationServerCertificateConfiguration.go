// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewalltlsinspectionconfiguration


type NetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/networkfirewall_tls_inspection_configuration#certificate_authority_arn NetworkfirewallTlsInspectionConfiguration#certificate_authority_arn}.
	CertificateAuthorityArn *string `field:"optional" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// check_certificate_revocation_status block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/networkfirewall_tls_inspection_configuration#check_certificate_revocation_status NetworkfirewallTlsInspectionConfiguration#check_certificate_revocation_status}
	CheckCertificateRevocationStatus interface{} `field:"optional" json:"checkCertificateRevocationStatus" yaml:"checkCertificateRevocationStatus"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/networkfirewall_tls_inspection_configuration#scope NetworkfirewallTlsInspectionConfiguration#scope}
	Scope interface{} `field:"optional" json:"scope" yaml:"scope"`
	// server_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/networkfirewall_tls_inspection_configuration#server_certificate NetworkfirewallTlsInspectionConfiguration#server_certificate}
	ServerCertificate interface{} `field:"optional" json:"serverCertificate" yaml:"serverCertificate"`
}

