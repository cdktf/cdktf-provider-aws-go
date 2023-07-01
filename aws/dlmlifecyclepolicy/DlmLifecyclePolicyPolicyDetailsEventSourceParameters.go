package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsEventSourceParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/dlm_lifecycle_policy#description_regex DlmLifecyclePolicy#description_regex}.
	DescriptionRegex *string `field:"required" json:"descriptionRegex" yaml:"descriptionRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/dlm_lifecycle_policy#event_type DlmLifecyclePolicy#event_type}.
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/dlm_lifecycle_policy#snapshot_owner DlmLifecyclePolicy#snapshot_owner}.
	SnapshotOwner *[]*string `field:"required" json:"snapshotOwner" yaml:"snapshotOwner"`
}

