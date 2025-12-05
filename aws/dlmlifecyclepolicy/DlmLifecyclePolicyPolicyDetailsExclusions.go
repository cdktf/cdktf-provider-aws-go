// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsExclusions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dlm_lifecycle_policy#exclude_boot_volumes DlmLifecyclePolicy#exclude_boot_volumes}.
	ExcludeBootVolumes interface{} `field:"optional" json:"excludeBootVolumes" yaml:"excludeBootVolumes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dlm_lifecycle_policy#exclude_tags DlmLifecyclePolicy#exclude_tags}.
	ExcludeTags *map[string]*string `field:"optional" json:"excludeTags" yaml:"excludeTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dlm_lifecycle_policy#exclude_volume_types DlmLifecyclePolicy#exclude_volume_types}.
	ExcludeVolumeTypes *[]*string `field:"optional" json:"excludeVolumeTypes" yaml:"excludeVolumeTypes"`
}

