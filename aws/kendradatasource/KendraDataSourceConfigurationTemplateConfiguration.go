// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kendradatasource


type KendraDataSourceConfigurationTemplateConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/kendra_data_source#template KendraDataSource#template}.
	Template *string `field:"required" json:"template" yaml:"template"`
}

