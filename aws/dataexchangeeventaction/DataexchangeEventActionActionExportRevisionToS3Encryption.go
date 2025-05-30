// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataexchangeeventaction


type DataexchangeEventActionActionExportRevisionToS3Encryption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/dataexchange_event_action#kms_key_arn DataexchangeEventAction#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/dataexchange_event_action#type DataexchangeEventAction#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

