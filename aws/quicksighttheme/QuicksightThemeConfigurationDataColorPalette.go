// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme


type QuicksightThemeConfigurationDataColorPalette struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/quicksight_theme#colors QuicksightTheme#colors}.
	Colors *[]*string `field:"optional" json:"colors" yaml:"colors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/quicksight_theme#empty_fill_color QuicksightTheme#empty_fill_color}.
	EmptyFillColor *string `field:"optional" json:"emptyFillColor" yaml:"emptyFillColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/quicksight_theme#min_max_gradient QuicksightTheme#min_max_gradient}.
	MinMaxGradient *[]*string `field:"optional" json:"minMaxGradient" yaml:"minMaxGradient"`
}

