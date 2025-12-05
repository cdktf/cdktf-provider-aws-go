// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ekscapability


type EksCapabilityConfigurationArgoCdNetworkAccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/eks_capability#vpce_ids EksCapability#vpce_ids}.
	VpceIds *[]*string `field:"optional" json:"vpceIds" yaml:"vpceIds"`
}

