// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daxcluster


type DaxClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/dax_cluster#create DaxCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/dax_cluster#delete DaxCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/dax_cluster#update DaxCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

