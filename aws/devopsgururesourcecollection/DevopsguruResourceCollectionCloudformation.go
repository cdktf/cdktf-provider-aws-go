// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devopsgururesourcecollection


type DevopsguruResourceCollectionCloudformation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/devopsguru_resource_collection#stack_names DevopsguruResourceCollection#stack_names}.
	StackNames *[]*string `field:"required" json:"stackNames" yaml:"stackNames"`
}

