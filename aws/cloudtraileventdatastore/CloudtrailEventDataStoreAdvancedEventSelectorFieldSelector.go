// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudtraileventdatastore


type CloudtrailEventDataStoreAdvancedEventSelectorFieldSelector struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/cloudtrail_event_data_store#ends_with CloudtrailEventDataStore#ends_with}.
	EndsWith *[]*string `field:"optional" json:"endsWith" yaml:"endsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/cloudtrail_event_data_store#equals CloudtrailEventDataStore#equals}.
	EqualTo *[]*string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/cloudtrail_event_data_store#field CloudtrailEventDataStore#field}.
	Field *string `field:"optional" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/cloudtrail_event_data_store#not_ends_with CloudtrailEventDataStore#not_ends_with}.
	NotEndsWith *[]*string `field:"optional" json:"notEndsWith" yaml:"notEndsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/cloudtrail_event_data_store#not_equals CloudtrailEventDataStore#not_equals}.
	NotEquals *[]*string `field:"optional" json:"notEquals" yaml:"notEquals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/cloudtrail_event_data_store#not_starts_with CloudtrailEventDataStore#not_starts_with}.
	NotStartsWith *[]*string `field:"optional" json:"notStartsWith" yaml:"notStartsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/cloudtrail_event_data_store#starts_with CloudtrailEventDataStore#starts_with}.
	StartsWith *[]*string `field:"optional" json:"startsWith" yaml:"startsWith"`
}

