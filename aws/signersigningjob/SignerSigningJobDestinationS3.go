// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package signersigningjob


type SignerSigningJobDestinationS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/signer_signing_job#bucket SignerSigningJob#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/signer_signing_job#prefix SignerSigningJob#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

