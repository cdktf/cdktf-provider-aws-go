// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataexchangeeventaction


type DataexchangeEventActionActionExportRevisionToS3RevisionDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/dataexchange_event_action#bucket DataexchangeEventAction#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/dataexchange_event_action#key_pattern DataexchangeEventAction#key_pattern}.
	KeyPattern *string `field:"optional" json:"keyPattern" yaml:"keyPattern"`
}

