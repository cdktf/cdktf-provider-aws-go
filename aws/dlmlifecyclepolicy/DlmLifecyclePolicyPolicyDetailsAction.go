package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsAction struct {
	// cross_region_copy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#cross_region_copy DlmLifecyclePolicy#cross_region_copy}
	CrossRegionCopy interface{} `field:"required" json:"crossRegionCopy" yaml:"crossRegionCopy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#name DlmLifecyclePolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

