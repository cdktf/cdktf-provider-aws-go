// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasynctask


type DatasyncTaskTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/datasync_task#create DatasyncTask#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

