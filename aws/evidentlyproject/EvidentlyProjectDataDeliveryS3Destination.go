// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package evidentlyproject


type EvidentlyProjectDataDeliveryS3Destination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/evidently_project#bucket EvidentlyProject#bucket}.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/evidently_project#prefix EvidentlyProject#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

