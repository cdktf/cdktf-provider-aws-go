// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasynclocationhdfs


type DatasyncLocationHdfsNameNode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/datasync_location_hdfs#hostname DatasyncLocationHdfs#hostname}.
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/datasync_location_hdfs#port DatasyncLocationHdfs#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

