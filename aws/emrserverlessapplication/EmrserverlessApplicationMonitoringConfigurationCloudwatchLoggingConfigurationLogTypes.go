// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrserverlessapplication


type EmrserverlessApplicationMonitoringConfigurationCloudwatchLoggingConfigurationLogTypes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/emrserverless_application#name EmrserverlessApplication#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/emrserverless_application#values EmrserverlessApplication#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

