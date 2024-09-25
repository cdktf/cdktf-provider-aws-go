// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterOutpostConfigControlPlanePlacement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/eks_cluster#group_name EksCluster#group_name}.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
}

