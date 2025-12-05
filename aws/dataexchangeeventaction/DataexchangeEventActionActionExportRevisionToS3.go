// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataexchangeeventaction


type DataexchangeEventActionActionExportRevisionToS3 struct {
	// encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dataexchange_event_action#encryption DataexchangeEventAction#encryption}
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// revision_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dataexchange_event_action#revision_destination DataexchangeEventAction#revision_destination}
	RevisionDestination interface{} `field:"optional" json:"revisionDestination" yaml:"revisionDestination"`
}

