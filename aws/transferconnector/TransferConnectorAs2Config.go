// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferconnector


type TransferConnectorAs2Config struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/transfer_connector#compression TransferConnector#compression}.
	Compression *string `field:"required" json:"compression" yaml:"compression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/transfer_connector#encryption_algorithm TransferConnector#encryption_algorithm}.
	EncryptionAlgorithm *string `field:"required" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/transfer_connector#local_profile_id TransferConnector#local_profile_id}.
	LocalProfileId *string `field:"required" json:"localProfileId" yaml:"localProfileId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/transfer_connector#mdn_response TransferConnector#mdn_response}.
	MdnResponse *string `field:"required" json:"mdnResponse" yaml:"mdnResponse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/transfer_connector#partner_profile_id TransferConnector#partner_profile_id}.
	PartnerProfileId *string `field:"required" json:"partnerProfileId" yaml:"partnerProfileId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/transfer_connector#signing_algorithm TransferConnector#signing_algorithm}.
	SigningAlgorithm *string `field:"required" json:"signingAlgorithm" yaml:"signingAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/transfer_connector#mdn_signing_algorithm TransferConnector#mdn_signing_algorithm}.
	MdnSigningAlgorithm *string `field:"optional" json:"mdnSigningAlgorithm" yaml:"mdnSigningAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/transfer_connector#message_subject TransferConnector#message_subject}.
	MessageSubject *string `field:"optional" json:"messageSubject" yaml:"messageSubject"`
}

