// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ekscapability


type EksCapabilityConfiguration struct {
	// argo_cd block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/eks_capability#argo_cd EksCapability#argo_cd}
	ArgoCd interface{} `field:"optional" json:"argoCd" yaml:"argoCd"`
}

