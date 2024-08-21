// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshiftserverlessworkgroup


type RedshiftserverlessWorkgroupConfigParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/redshiftserverless_workgroup#parameter_key RedshiftserverlessWorkgroup#parameter_key}.
	ParameterKey *string `field:"required" json:"parameterKey" yaml:"parameterKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/redshiftserverless_workgroup#parameter_value RedshiftserverlessWorkgroup#parameter_value}.
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

