// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontfieldlevelencryptionconfig


type CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfilesItems struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudfront_field_level_encryption_config#profile_id CloudfrontFieldLevelEncryptionConfig#profile_id}.
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudfront_field_level_encryption_config#query_arg CloudfrontFieldLevelEncryptionConfig#query_arg}.
	QueryArg *string `field:"required" json:"queryArg" yaml:"queryArg"`
}

