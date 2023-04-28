package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/dlm_lifecycle_policy#interval DlmLifecyclePolicy#interval}.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/dlm_lifecycle_policy#interval_unit DlmLifecyclePolicy#interval_unit}.
	IntervalUnit *string `field:"required" json:"intervalUnit" yaml:"intervalUnit"`
}

