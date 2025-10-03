// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchdomain


type OpensearchDomainClusterConfigNodeOptions struct {
	// node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/opensearch_domain#node_config OpensearchDomain#node_config}
	NodeConfig *OpensearchDomainClusterConfigNodeOptionsNodeConfig `field:"optional" json:"nodeConfig" yaml:"nodeConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/opensearch_domain#node_type OpensearchDomain#node_type}.
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
}

