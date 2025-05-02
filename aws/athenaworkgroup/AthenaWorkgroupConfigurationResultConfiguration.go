// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkgroupConfigurationResultConfiguration struct {
	// acl_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/athena_workgroup#acl_configuration AthenaWorkgroup#acl_configuration}
	AclConfiguration *AthenaWorkgroupConfigurationResultConfigurationAclConfiguration `field:"optional" json:"aclConfiguration" yaml:"aclConfiguration"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/athena_workgroup#encryption_configuration AthenaWorkgroup#encryption_configuration}
	EncryptionConfiguration *AthenaWorkgroupConfigurationResultConfigurationEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/athena_workgroup#expected_bucket_owner AthenaWorkgroup#expected_bucket_owner}.
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/athena_workgroup#output_location AthenaWorkgroup#output_location}.
	OutputLocation *string `field:"optional" json:"outputLocation" yaml:"outputLocation"`
}

