// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkgroupConfigurationManagedQueryResultsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/athena_workgroup#enabled AthenaWorkgroup#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/athena_workgroup#encryption_configuration AthenaWorkgroup#encryption_configuration}
	EncryptionConfiguration *AthenaWorkgroupConfigurationManagedQueryResultsConfigurationEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
}

