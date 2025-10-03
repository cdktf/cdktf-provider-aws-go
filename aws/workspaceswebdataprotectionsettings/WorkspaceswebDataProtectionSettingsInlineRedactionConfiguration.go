// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspaceswebdataprotectionsettings


type WorkspaceswebDataProtectionSettingsInlineRedactionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/workspacesweb_data_protection_settings#global_confidence_level WorkspaceswebDataProtectionSettings#global_confidence_level}.
	GlobalConfidenceLevel *float64 `field:"optional" json:"globalConfidenceLevel" yaml:"globalConfidenceLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/workspacesweb_data_protection_settings#global_enforced_urls WorkspaceswebDataProtectionSettings#global_enforced_urls}.
	GlobalEnforcedUrls *[]*string `field:"optional" json:"globalEnforcedUrls" yaml:"globalEnforcedUrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/workspacesweb_data_protection_settings#global_exempt_urls WorkspaceswebDataProtectionSettings#global_exempt_urls}.
	GlobalExemptUrls *[]*string `field:"optional" json:"globalExemptUrls" yaml:"globalExemptUrls"`
	// inline_redaction_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/workspacesweb_data_protection_settings#inline_redaction_pattern WorkspaceswebDataProtectionSettings#inline_redaction_pattern}
	InlineRedactionPattern interface{} `field:"optional" json:"inlineRedactionPattern" yaml:"inlineRedactionPattern"`
}

