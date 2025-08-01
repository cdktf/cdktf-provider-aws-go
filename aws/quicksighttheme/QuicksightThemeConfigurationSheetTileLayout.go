// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme


type QuicksightThemeConfigurationSheetTileLayout struct {
	// gutter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/quicksight_theme#gutter QuicksightTheme#gutter}
	Gutter *QuicksightThemeConfigurationSheetTileLayoutGutter `field:"optional" json:"gutter" yaml:"gutter"`
	// margin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/quicksight_theme#margin QuicksightTheme#margin}
	Margin *QuicksightThemeConfigurationSheetTileLayoutMargin `field:"optional" json:"margin" yaml:"margin"`
}

