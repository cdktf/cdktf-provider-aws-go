// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterClientAuthentication struct {
	// sasl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/msk_cluster#sasl MskCluster#sasl}
	Sasl *MskClusterClientAuthenticationSasl `field:"optional" json:"sasl" yaml:"sasl"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/msk_cluster#tls MskCluster#tls}
	Tls *MskClusterClientAuthenticationTls `field:"optional" json:"tls" yaml:"tls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/msk_cluster#unauthenticated MskCluster#unauthenticated}.
	Unauthenticated interface{} `field:"optional" json:"unauthenticated" yaml:"unauthenticated"`
}

