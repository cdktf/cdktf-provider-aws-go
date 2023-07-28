package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsEventSource struct {
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/dlm_lifecycle_policy#parameters DlmLifecyclePolicy#parameters}
	Parameters *DlmLifecyclePolicyPolicyDetailsEventSourceParameters `field:"required" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/dlm_lifecycle_policy#type DlmLifecyclePolicy#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

