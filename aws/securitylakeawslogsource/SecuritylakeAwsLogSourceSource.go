// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitylakeawslogsource


type SecuritylakeAwsLogSourceSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/securitylake_aws_log_source#regions SecuritylakeAwsLogSource#regions}.
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/securitylake_aws_log_source#source_name SecuritylakeAwsLogSource#source_name}.
	SourceName *string `field:"required" json:"sourceName" yaml:"sourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/securitylake_aws_log_source#accounts SecuritylakeAwsLogSource#accounts}.
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/securitylake_aws_log_source#source_version SecuritylakeAwsLogSource#source_version}.
	SourceVersion *string `field:"optional" json:"sourceVersion" yaml:"sourceVersion"`
}

