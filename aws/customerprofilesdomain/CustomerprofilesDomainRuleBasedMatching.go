// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customerprofilesdomain


type CustomerprofilesDomainRuleBasedMatching struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/customerprofiles_domain#enabled CustomerprofilesDomain#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// attribute_types_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/customerprofiles_domain#attribute_types_selector CustomerprofilesDomain#attribute_types_selector}
	AttributeTypesSelector *CustomerprofilesDomainRuleBasedMatchingAttributeTypesSelector `field:"optional" json:"attributeTypesSelector" yaml:"attributeTypesSelector"`
	// conflict_resolution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/customerprofiles_domain#conflict_resolution CustomerprofilesDomain#conflict_resolution}
	ConflictResolution *CustomerprofilesDomainRuleBasedMatchingConflictResolution `field:"optional" json:"conflictResolution" yaml:"conflictResolution"`
	// exporting_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/customerprofiles_domain#exporting_config CustomerprofilesDomain#exporting_config}
	ExportingConfig *CustomerprofilesDomainRuleBasedMatchingExportingConfig `field:"optional" json:"exportingConfig" yaml:"exportingConfig"`
	// matching_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/customerprofiles_domain#matching_rules CustomerprofilesDomain#matching_rules}
	MatchingRules interface{} `field:"optional" json:"matchingRules" yaml:"matchingRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/customerprofiles_domain#max_allowed_rule_level_for_matching CustomerprofilesDomain#max_allowed_rule_level_for_matching}.
	MaxAllowedRuleLevelForMatching *float64 `field:"optional" json:"maxAllowedRuleLevelForMatching" yaml:"maxAllowedRuleLevelForMatching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/customerprofiles_domain#max_allowed_rule_level_for_merging CustomerprofilesDomain#max_allowed_rule_level_for_merging}.
	MaxAllowedRuleLevelForMerging *float64 `field:"optional" json:"maxAllowedRuleLevelForMerging" yaml:"maxAllowedRuleLevelForMerging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/customerprofiles_domain#status CustomerprofilesDomain#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

