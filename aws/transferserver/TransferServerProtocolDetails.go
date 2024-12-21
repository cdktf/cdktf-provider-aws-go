// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferserver


type TransferServerProtocolDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/transfer_server#as2_transports TransferServer#as2_transports}.
	As2Transports *[]*string `field:"optional" json:"as2Transports" yaml:"as2Transports"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/transfer_server#passive_ip TransferServer#passive_ip}.
	PassiveIp *string `field:"optional" json:"passiveIp" yaml:"passiveIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/transfer_server#set_stat_option TransferServer#set_stat_option}.
	SetStatOption *string `field:"optional" json:"setStatOption" yaml:"setStatOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/transfer_server#tls_session_resumption_mode TransferServer#tls_session_resumption_mode}.
	TlsSessionResumptionMode *string `field:"optional" json:"tlsSessionResumptionMode" yaml:"tlsSessionResumptionMode"`
}

