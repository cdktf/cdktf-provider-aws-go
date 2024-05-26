// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskserverlesscluster


type MskServerlessClusterClientAuthentication struct {
	// sasl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/msk_serverless_cluster#sasl MskServerlessCluster#sasl}
	Sasl *MskServerlessClusterClientAuthenticationSasl `field:"required" json:"sasl" yaml:"sasl"`
}

