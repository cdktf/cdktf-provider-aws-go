package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsSchedule struct {
	// create_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/dlm_lifecycle_policy#create_rule DlmLifecyclePolicy#create_rule}
	CreateRule *DlmLifecyclePolicyPolicyDetailsScheduleCreateRule `field:"required" json:"createRule" yaml:"createRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/dlm_lifecycle_policy#name DlmLifecyclePolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// retain_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/dlm_lifecycle_policy#retain_rule DlmLifecyclePolicy#retain_rule}
	RetainRule *DlmLifecyclePolicyPolicyDetailsScheduleRetainRule `field:"required" json:"retainRule" yaml:"retainRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/dlm_lifecycle_policy#copy_tags DlmLifecyclePolicy#copy_tags}.
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// cross_region_copy_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/dlm_lifecycle_policy#cross_region_copy_rule DlmLifecyclePolicy#cross_region_copy_rule}
	CrossRegionCopyRule interface{} `field:"optional" json:"crossRegionCopyRule" yaml:"crossRegionCopyRule"`
	// deprecate_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/dlm_lifecycle_policy#deprecate_rule DlmLifecyclePolicy#deprecate_rule}
	DeprecateRule *DlmLifecyclePolicyPolicyDetailsScheduleDeprecateRule `field:"optional" json:"deprecateRule" yaml:"deprecateRule"`
	// fast_restore_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/dlm_lifecycle_policy#fast_restore_rule DlmLifecyclePolicy#fast_restore_rule}
	FastRestoreRule *DlmLifecyclePolicyPolicyDetailsScheduleFastRestoreRule `field:"optional" json:"fastRestoreRule" yaml:"fastRestoreRule"`
	// share_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/dlm_lifecycle_policy#share_rule DlmLifecyclePolicy#share_rule}
	ShareRule *DlmLifecyclePolicyPolicyDetailsScheduleShareRule `field:"optional" json:"shareRule" yaml:"shareRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/dlm_lifecycle_policy#tags_to_add DlmLifecyclePolicy#tags_to_add}.
	TagsToAdd *map[string]*string `field:"optional" json:"tagsToAdd" yaml:"tagsToAdd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/dlm_lifecycle_policy#variable_tags DlmLifecyclePolicy#variable_tags}.
	VariableTags *map[string]*string `field:"optional" json:"variableTags" yaml:"variableTags"`
}

