// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchoutboundconnection


type OpensearchOutboundConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/opensearch_outbound_connection#create OpensearchOutboundConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/opensearch_outbound_connection#delete OpensearchOutboundConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

