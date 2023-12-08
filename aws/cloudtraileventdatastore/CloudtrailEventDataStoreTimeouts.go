// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudtraileventdatastore


type CloudtrailEventDataStoreTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/cloudtrail_event_data_store#create CloudtrailEventDataStore#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/cloudtrail_event_data_store#delete CloudtrailEventDataStore#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/cloudtrail_event_data_store#update CloudtrailEventDataStore#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

