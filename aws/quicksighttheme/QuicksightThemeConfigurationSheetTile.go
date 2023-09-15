// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme


type QuicksightThemeConfigurationSheetTile struct {
	// border block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.17.0/docs/resources/quicksight_theme#border QuicksightTheme#border}
	Border *QuicksightThemeConfigurationSheetTileBorder `field:"optional" json:"border" yaml:"border"`
}

