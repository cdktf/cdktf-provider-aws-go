// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentKendraConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/lexv2models_intent#kendra_index Lexv2ModelsIntent#kendra_index}.
	KendraIndex *string `field:"required" json:"kendraIndex" yaml:"kendraIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/lexv2models_intent#query_filter_string Lexv2ModelsIntent#query_filter_string}.
	QueryFilterString *string `field:"optional" json:"queryFilterString" yaml:"queryFilterString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/lexv2models_intent#query_filter_string_enabled Lexv2ModelsIntent#query_filter_string_enabled}.
	QueryFilterStringEnabled interface{} `field:"optional" json:"queryFilterStringEnabled" yaml:"queryFilterStringEnabled"`
}

