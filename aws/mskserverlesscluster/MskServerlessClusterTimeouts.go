// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskserverlesscluster


type MskServerlessClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/msk_serverless_cluster#create MskServerlessCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/msk_serverless_cluster#delete MskServerlessCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

