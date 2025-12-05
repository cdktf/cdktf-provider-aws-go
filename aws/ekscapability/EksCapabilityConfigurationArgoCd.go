// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ekscapability


type EksCapabilityConfigurationArgoCd struct {
	// aws_idc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/eks_capability#aws_idc EksCapability#aws_idc}
	AwsIdc interface{} `field:"optional" json:"awsIdc" yaml:"awsIdc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/eks_capability#namespace EksCapability#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// network_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/eks_capability#network_access EksCapability#network_access}
	NetworkAccess interface{} `field:"optional" json:"networkAccess" yaml:"networkAccess"`
	// rbac_role_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/eks_capability#rbac_role_mapping EksCapability#rbac_role_mapping}
	RbacRoleMapping interface{} `field:"optional" json:"rbacRoleMapping" yaml:"rbacRoleMapping"`
}

