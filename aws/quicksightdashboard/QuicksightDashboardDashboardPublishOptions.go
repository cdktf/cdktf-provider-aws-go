// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdashboard


type QuicksightDashboardDashboardPublishOptions struct {
	// ad_hoc_filtering_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/quicksight_dashboard#ad_hoc_filtering_option QuicksightDashboard#ad_hoc_filtering_option}
	AdHocFilteringOption *QuicksightDashboardDashboardPublishOptionsAdHocFilteringOption `field:"optional" json:"adHocFilteringOption" yaml:"adHocFilteringOption"`
	// data_point_drill_up_down_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/quicksight_dashboard#data_point_drill_up_down_option QuicksightDashboard#data_point_drill_up_down_option}
	DataPointDrillUpDownOption *QuicksightDashboardDashboardPublishOptionsDataPointDrillUpDownOption `field:"optional" json:"dataPointDrillUpDownOption" yaml:"dataPointDrillUpDownOption"`
	// data_point_menu_label_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/quicksight_dashboard#data_point_menu_label_option QuicksightDashboard#data_point_menu_label_option}
	DataPointMenuLabelOption *QuicksightDashboardDashboardPublishOptionsDataPointMenuLabelOption `field:"optional" json:"dataPointMenuLabelOption" yaml:"dataPointMenuLabelOption"`
	// data_point_tooltip_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/quicksight_dashboard#data_point_tooltip_option QuicksightDashboard#data_point_tooltip_option}
	DataPointTooltipOption *QuicksightDashboardDashboardPublishOptionsDataPointTooltipOption `field:"optional" json:"dataPointTooltipOption" yaml:"dataPointTooltipOption"`
	// export_to_csv_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/quicksight_dashboard#export_to_csv_option QuicksightDashboard#export_to_csv_option}
	ExportToCsvOption *QuicksightDashboardDashboardPublishOptionsExportToCsvOption `field:"optional" json:"exportToCsvOption" yaml:"exportToCsvOption"`
	// export_with_hidden_fields_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/quicksight_dashboard#export_with_hidden_fields_option QuicksightDashboard#export_with_hidden_fields_option}
	ExportWithHiddenFieldsOption *QuicksightDashboardDashboardPublishOptionsExportWithHiddenFieldsOption `field:"optional" json:"exportWithHiddenFieldsOption" yaml:"exportWithHiddenFieldsOption"`
	// sheet_controls_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/quicksight_dashboard#sheet_controls_option QuicksightDashboard#sheet_controls_option}
	SheetControlsOption *QuicksightDashboardDashboardPublishOptionsSheetControlsOption `field:"optional" json:"sheetControlsOption" yaml:"sheetControlsOption"`
	// sheet_layout_element_maximization_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/quicksight_dashboard#sheet_layout_element_maximization_option QuicksightDashboard#sheet_layout_element_maximization_option}
	SheetLayoutElementMaximizationOption *QuicksightDashboardDashboardPublishOptionsSheetLayoutElementMaximizationOption `field:"optional" json:"sheetLayoutElementMaximizationOption" yaml:"sheetLayoutElementMaximizationOption"`
	// visual_axis_sort_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/quicksight_dashboard#visual_axis_sort_option QuicksightDashboard#visual_axis_sort_option}
	VisualAxisSortOption *QuicksightDashboardDashboardPublishOptionsVisualAxisSortOption `field:"optional" json:"visualAxisSortOption" yaml:"visualAxisSortOption"`
	// visual_menu_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/quicksight_dashboard#visual_menu_option QuicksightDashboard#visual_menu_option}
	VisualMenuOption *QuicksightDashboardDashboardPublishOptionsVisualMenuOption `field:"optional" json:"visualMenuOption" yaml:"visualMenuOption"`
}

