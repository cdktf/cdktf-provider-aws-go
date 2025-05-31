// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchpackageassociation


type OpensearchPackageAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opensearch_package_association#create OpensearchPackageAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opensearch_package_association#delete OpensearchPackageAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

