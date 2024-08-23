// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewalltlsinspectionconfiguration


type NetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationCheckCertificateRevocationStatus struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/networkfirewall_tls_inspection_configuration#revoked_status_action NetworkfirewallTlsInspectionConfiguration#revoked_status_action}.
	RevokedStatusAction *string `field:"optional" json:"revokedStatusAction" yaml:"revokedStatusAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/networkfirewall_tls_inspection_configuration#unknown_status_action NetworkfirewallTlsInspectionConfiguration#unknown_status_action}.
	UnknownStatusAction *string `field:"optional" json:"unknownStatusAction" yaml:"unknownStatusAction"`
}

