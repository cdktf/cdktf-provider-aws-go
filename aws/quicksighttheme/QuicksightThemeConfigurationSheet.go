// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme


type QuicksightThemeConfigurationSheet struct {
	// tile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/quicksight_theme#tile QuicksightTheme#tile}
	Tile *QuicksightThemeConfigurationSheetTile `field:"optional" json:"tile" yaml:"tile"`
	// tile_layout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/quicksight_theme#tile_layout QuicksightTheme#tile_layout}
	TileLayout *QuicksightThemeConfigurationSheetTileLayout `field:"optional" json:"tileLayout" yaml:"tileLayout"`
}

