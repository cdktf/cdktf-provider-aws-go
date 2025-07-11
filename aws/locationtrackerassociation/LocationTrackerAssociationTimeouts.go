// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package locationtrackerassociation


type LocationTrackerAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/location_tracker_association#create LocationTrackerAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/location_tracker_association#delete LocationTrackerAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

