// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterUpgradePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/eks_cluster#support_type EksCluster#support_type}.
	SupportType *string `field:"optional" json:"supportType" yaml:"supportType"`
}

