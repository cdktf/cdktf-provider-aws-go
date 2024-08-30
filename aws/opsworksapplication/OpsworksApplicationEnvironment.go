// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksapplication


type OpsworksApplicationEnvironment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/opsworks_application#key OpsworksApplication#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/opsworks_application#value OpsworksApplication#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/opsworks_application#secure OpsworksApplication#secure}.
	Secure interface{} `field:"optional" json:"secure" yaml:"secure"`
}

