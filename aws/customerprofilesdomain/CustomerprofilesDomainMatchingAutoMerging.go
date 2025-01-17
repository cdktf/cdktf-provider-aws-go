// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customerprofilesdomain


type CustomerprofilesDomainMatchingAutoMerging struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/customerprofiles_domain#enabled CustomerprofilesDomain#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// conflict_resolution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/customerprofiles_domain#conflict_resolution CustomerprofilesDomain#conflict_resolution}
	ConflictResolution *CustomerprofilesDomainMatchingAutoMergingConflictResolution `field:"optional" json:"conflictResolution" yaml:"conflictResolution"`
	// consolidation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/customerprofiles_domain#consolidation CustomerprofilesDomain#consolidation}
	Consolidation *CustomerprofilesDomainMatchingAutoMergingConsolidation `field:"optional" json:"consolidation" yaml:"consolidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/customerprofiles_domain#min_allowed_confidence_score_for_merging CustomerprofilesDomain#min_allowed_confidence_score_for_merging}.
	MinAllowedConfidenceScoreForMerging *float64 `field:"optional" json:"minAllowedConfidenceScoreForMerging" yaml:"minAllowedConfidenceScoreForMerging"`
}

