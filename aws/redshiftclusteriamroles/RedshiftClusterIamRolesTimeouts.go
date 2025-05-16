// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshiftclusteriamroles


type RedshiftClusterIamRolesTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/redshift_cluster_iam_roles#create RedshiftClusterIamRoles#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/redshift_cluster_iam_roles#delete RedshiftClusterIamRoles#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/redshift_cluster_iam_roles#update RedshiftClusterIamRoles#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

