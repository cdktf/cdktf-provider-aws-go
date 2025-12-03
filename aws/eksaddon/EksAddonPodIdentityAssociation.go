// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eksaddon


type EksAddonPodIdentityAssociation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/eks_addon#role_arn EksAddon#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/eks_addon#service_account EksAddon#service_account}.
	ServiceAccount *string `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
}

