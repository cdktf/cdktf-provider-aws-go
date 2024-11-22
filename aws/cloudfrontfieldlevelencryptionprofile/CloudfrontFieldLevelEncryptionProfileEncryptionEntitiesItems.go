// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontfieldlevelencryptionprofile


type CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItems struct {
	// field_patterns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.77.0/docs/resources/cloudfront_field_level_encryption_profile#field_patterns CloudfrontFieldLevelEncryptionProfile#field_patterns}
	FieldPatterns *CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsFieldPatterns `field:"required" json:"fieldPatterns" yaml:"fieldPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.77.0/docs/resources/cloudfront_field_level_encryption_profile#provider_id CloudfrontFieldLevelEncryptionProfile#provider_id}.
	ProviderId *string `field:"required" json:"providerId" yaml:"providerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.77.0/docs/resources/cloudfront_field_level_encryption_profile#public_key_id CloudfrontFieldLevelEncryptionProfile#public_key_id}.
	PublicKeyId *string `field:"required" json:"publicKeyId" yaml:"publicKeyId"`
}

