// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationRunConfigurationApplicationRestoreConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/kinesisanalyticsv2_application#application_restore_type Kinesisanalyticsv2Application#application_restore_type}.
	ApplicationRestoreType *string `field:"optional" json:"applicationRestoreType" yaml:"applicationRestoreType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/kinesisanalyticsv2_application#snapshot_name Kinesisanalyticsv2Application#snapshot_name}.
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
}

