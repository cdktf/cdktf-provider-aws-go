// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewalltlsinspectionconfiguration


type NetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationScopeSourcePorts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/networkfirewall_tls_inspection_configuration#from_port NetworkfirewallTlsInspectionConfiguration#from_port}.
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/networkfirewall_tls_inspection_configuration#to_port NetworkfirewallTlsInspectionConfiguration#to_port}.
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
}

