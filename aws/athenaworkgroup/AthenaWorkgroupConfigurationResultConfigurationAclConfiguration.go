// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkgroupConfigurationResultConfigurationAclConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/athena_workgroup#s3_acl_option AthenaWorkgroup#s3_acl_option}.
	S3AclOption *string `field:"required" json:"s3AclOption" yaml:"s3AclOption"`
}

