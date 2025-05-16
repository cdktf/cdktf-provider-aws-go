// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsavailabilityzones

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsAvailabilityZonesConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/data-sources/availability_zones#all_availability_zones DataAwsAvailabilityZones#all_availability_zones}.
	AllAvailabilityZones interface{} `field:"optional" json:"allAvailabilityZones" yaml:"allAvailabilityZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/data-sources/availability_zones#exclude_names DataAwsAvailabilityZones#exclude_names}.
	ExcludeNames *[]*string `field:"optional" json:"excludeNames" yaml:"excludeNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/data-sources/availability_zones#exclude_zone_ids DataAwsAvailabilityZones#exclude_zone_ids}.
	ExcludeZoneIds *[]*string `field:"optional" json:"excludeZoneIds" yaml:"excludeZoneIds"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/data-sources/availability_zones#filter DataAwsAvailabilityZones#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/data-sources/availability_zones#id DataAwsAvailabilityZones#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/data-sources/availability_zones#state DataAwsAvailabilityZones#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/data-sources/availability_zones#timeouts DataAwsAvailabilityZones#timeouts}
	Timeouts *DataAwsAvailabilityZonesTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

