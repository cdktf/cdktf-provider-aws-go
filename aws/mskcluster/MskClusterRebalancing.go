// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterRebalancing struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/msk_cluster#status MskCluster#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
}

