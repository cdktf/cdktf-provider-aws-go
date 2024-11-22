// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitylakecustomlogsource


type SecuritylakeCustomLogSourceConfigurationProviderIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.77.0/docs/resources/securitylake_custom_log_source#external_id SecuritylakeCustomLogSource#external_id}.
	ExternalId *string `field:"required" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.77.0/docs/resources/securitylake_custom_log_source#principal SecuritylakeCustomLogSource#principal}.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

