// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworkshaproxylayer


type OpsworksHaproxyLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_haproxy_layer#mount_point OpsworksHaproxyLayer#mount_point}.
	MountPoint *string `field:"required" json:"mountPoint" yaml:"mountPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_haproxy_layer#number_of_disks OpsworksHaproxyLayer#number_of_disks}.
	NumberOfDisks *float64 `field:"required" json:"numberOfDisks" yaml:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_haproxy_layer#size OpsworksHaproxyLayer#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_haproxy_layer#encrypted OpsworksHaproxyLayer#encrypted}.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_haproxy_layer#iops OpsworksHaproxyLayer#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_haproxy_layer#raid_level OpsworksHaproxyLayer#raid_level}.
	RaidLevel *string `field:"optional" json:"raidLevel" yaml:"raidLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_haproxy_layer#type OpsworksHaproxyLayer#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

