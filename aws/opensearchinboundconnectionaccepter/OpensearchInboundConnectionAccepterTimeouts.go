// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchinboundconnectionaccepter


type OpensearchInboundConnectionAccepterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/opensearch_inbound_connection_accepter#create OpensearchInboundConnectionAccepter#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/opensearch_inbound_connection_accepter#delete OpensearchInboundConnectionAccepter#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

