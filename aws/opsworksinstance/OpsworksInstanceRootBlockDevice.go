// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksinstance


type OpsworksInstanceRootBlockDevice struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/opsworks_instance#delete_on_termination OpsworksInstance#delete_on_termination}.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/opsworks_instance#iops OpsworksInstance#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/opsworks_instance#volume_size OpsworksInstance#volume_size}.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/opsworks_instance#volume_type OpsworksInstance#volume_type}.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

