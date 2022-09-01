package dlm


type DlmLifecyclePolicyPolicyDetailsEventSource struct {
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#parameters DlmLifecyclePolicy#parameters}
	Parameters *DlmLifecyclePolicyPolicyDetailsEventSourceParameters `field:"required" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dlm_lifecycle_policy#type DlmLifecyclePolicy#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

