// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshiftserverlessworkgroup


type RedshiftserverlessWorkgroupPricePerformanceTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/redshiftserverless_workgroup#enabled RedshiftserverlessWorkgroup#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/redshiftserverless_workgroup#level RedshiftserverlessWorkgroup#level}.
	Level *float64 `field:"optional" json:"level" yaml:"level"`
}

