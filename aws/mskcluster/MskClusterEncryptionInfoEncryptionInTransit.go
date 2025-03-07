// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterEncryptionInfoEncryptionInTransit struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/msk_cluster#client_broker MskCluster#client_broker}.
	ClientBroker *string `field:"optional" json:"clientBroker" yaml:"clientBroker"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/msk_cluster#in_cluster MskCluster#in_cluster}.
	InCluster interface{} `field:"optional" json:"inCluster" yaml:"inCluster"`
}

