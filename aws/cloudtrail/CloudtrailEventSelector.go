// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudtrail


type CloudtrailEventSelector struct {
	// data_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/cloudtrail#data_resource Cloudtrail#data_resource}
	DataResource interface{} `field:"optional" json:"dataResource" yaml:"dataResource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/cloudtrail#exclude_management_event_sources Cloudtrail#exclude_management_event_sources}.
	ExcludeManagementEventSources *[]*string `field:"optional" json:"excludeManagementEventSources" yaml:"excludeManagementEventSources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/cloudtrail#include_management_events Cloudtrail#include_management_events}.
	IncludeManagementEvents interface{} `field:"optional" json:"includeManagementEvents" yaml:"includeManagementEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/cloudtrail#read_write_type Cloudtrail#read_write_type}.
	ReadWriteType *string `field:"optional" json:"readWriteType" yaml:"readWriteType"`
}

