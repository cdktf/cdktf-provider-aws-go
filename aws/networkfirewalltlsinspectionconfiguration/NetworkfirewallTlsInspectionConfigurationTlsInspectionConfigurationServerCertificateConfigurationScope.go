// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewalltlsinspectionconfiguration


type NetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationScope struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/networkfirewall_tls_inspection_configuration#protocols NetworkfirewallTlsInspectionConfiguration#protocols}.
	Protocols *[]*float64 `field:"required" json:"protocols" yaml:"protocols"`
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/networkfirewall_tls_inspection_configuration#destination NetworkfirewallTlsInspectionConfiguration#destination}
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// destination_ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/networkfirewall_tls_inspection_configuration#destination_ports NetworkfirewallTlsInspectionConfiguration#destination_ports}
	DestinationPorts interface{} `field:"optional" json:"destinationPorts" yaml:"destinationPorts"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/networkfirewall_tls_inspection_configuration#source NetworkfirewallTlsInspectionConfiguration#source}
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// source_ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/networkfirewall_tls_inspection_configuration#source_ports NetworkfirewallTlsInspectionConfiguration#source_ports}
	SourcePorts interface{} `field:"optional" json:"sourcePorts" yaml:"sourcePorts"`
}

