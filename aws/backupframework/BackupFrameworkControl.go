// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupframework


type BackupFrameworkControl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/backup_framework#name BackupFramework#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// input_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/backup_framework#input_parameter BackupFramework#input_parameter}
	InputParameter interface{} `field:"optional" json:"inputParameter" yaml:"inputParameter"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/backup_framework#scope BackupFramework#scope}
	Scope *BackupFrameworkControlScope `field:"optional" json:"scope" yaml:"scope"`
}

