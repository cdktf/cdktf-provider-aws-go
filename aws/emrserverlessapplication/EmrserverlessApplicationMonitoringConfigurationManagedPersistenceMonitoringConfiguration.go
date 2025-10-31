// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrserverlessapplication


type EmrserverlessApplicationMonitoringConfigurationManagedPersistenceMonitoringConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/emrserverless_application#enabled EmrserverlessApplication#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/emrserverless_application#encryption_key_arn EmrserverlessApplication#encryption_key_arn}.
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
}

