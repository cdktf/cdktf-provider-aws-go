// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetails struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/dlm_lifecycle_policy#action DlmLifecyclePolicy#action}
	Action *DlmLifecyclePolicyPolicyDetailsAction `field:"optional" json:"action" yaml:"action"`
	// event_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/dlm_lifecycle_policy#event_source DlmLifecyclePolicy#event_source}
	EventSource *DlmLifecyclePolicyPolicyDetailsEventSource `field:"optional" json:"eventSource" yaml:"eventSource"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/dlm_lifecycle_policy#parameters DlmLifecyclePolicy#parameters}
	Parameters *DlmLifecyclePolicyPolicyDetailsParameters `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/dlm_lifecycle_policy#policy_type DlmLifecyclePolicy#policy_type}.
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/dlm_lifecycle_policy#resource_locations DlmLifecyclePolicy#resource_locations}.
	ResourceLocations *[]*string `field:"optional" json:"resourceLocations" yaml:"resourceLocations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/dlm_lifecycle_policy#resource_types DlmLifecyclePolicy#resource_types}.
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/dlm_lifecycle_policy#schedule DlmLifecyclePolicy#schedule}
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/dlm_lifecycle_policy#target_tags DlmLifecyclePolicy#target_tags}.
	TargetTags *map[string]*string `field:"optional" json:"targetTags" yaml:"targetTags"`
}

