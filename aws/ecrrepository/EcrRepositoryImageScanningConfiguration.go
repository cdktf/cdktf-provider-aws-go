// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecrrepository


type EcrRepositoryImageScanningConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ecr_repository#scan_on_push EcrRepository#scan_on_push}.
	ScanOnPush interface{} `field:"required" json:"scanOnPush" yaml:"scanOnPush"`
}

