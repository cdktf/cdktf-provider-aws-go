// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sesv2configurationset


type Sesv2ConfigurationSetVdmOptions struct {
	// dashboard_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/sesv2_configuration_set#dashboard_options Sesv2ConfigurationSet#dashboard_options}
	DashboardOptions *Sesv2ConfigurationSetVdmOptionsDashboardOptions `field:"optional" json:"dashboardOptions" yaml:"dashboardOptions"`
	// guardian_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/sesv2_configuration_set#guardian_options Sesv2ConfigurationSet#guardian_options}
	GuardianOptions *Sesv2ConfigurationSetVdmOptionsGuardianOptions `field:"optional" json:"guardianOptions" yaml:"guardianOptions"`
}

