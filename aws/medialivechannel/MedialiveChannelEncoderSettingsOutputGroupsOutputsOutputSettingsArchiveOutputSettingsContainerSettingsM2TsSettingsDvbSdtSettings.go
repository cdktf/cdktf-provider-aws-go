// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2TsSettingsDvbSdtSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/medialive_channel#output_sdt MedialiveChannel#output_sdt}.
	OutputSdt *string `field:"optional" json:"outputSdt" yaml:"outputSdt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/medialive_channel#rep_interval MedialiveChannel#rep_interval}.
	RepInterval *float64 `field:"optional" json:"repInterval" yaml:"repInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/medialive_channel#service_name MedialiveChannel#service_name}.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/medialive_channel#service_provider_name MedialiveChannel#service_provider_name}.
	ServiceProviderName *string `field:"optional" json:"serviceProviderName" yaml:"serviceProviderName"`
}

