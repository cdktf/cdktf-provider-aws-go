// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme


type QuicksightThemeConfiguration struct {
	// data_color_palette block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/quicksight_theme#data_color_palette QuicksightTheme#data_color_palette}
	DataColorPalette *QuicksightThemeConfigurationDataColorPalette `field:"optional" json:"dataColorPalette" yaml:"dataColorPalette"`
	// sheet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/quicksight_theme#sheet QuicksightTheme#sheet}
	Sheet *QuicksightThemeConfigurationSheet `field:"optional" json:"sheet" yaml:"sheet"`
	// typography block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/quicksight_theme#typography QuicksightTheme#typography}
	Typography *QuicksightThemeConfigurationTypography `field:"optional" json:"typography" yaml:"typography"`
	// ui_color_palette block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/quicksight_theme#ui_color_palette QuicksightTheme#ui_color_palette}
	UiColorPalette *QuicksightThemeConfigurationUiColorPalette `field:"optional" json:"uiColorPalette" yaml:"uiColorPalette"`
}

