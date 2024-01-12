// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksganglialayer


type OpsworksGangliaLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_ganglia_layer#mount_point OpsworksGangliaLayer#mount_point}.
	MountPoint *string `field:"required" json:"mountPoint" yaml:"mountPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_ganglia_layer#number_of_disks OpsworksGangliaLayer#number_of_disks}.
	NumberOfDisks *float64 `field:"required" json:"numberOfDisks" yaml:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_ganglia_layer#size OpsworksGangliaLayer#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_ganglia_layer#encrypted OpsworksGangliaLayer#encrypted}.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_ganglia_layer#iops OpsworksGangliaLayer#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_ganglia_layer#raid_level OpsworksGangliaLayer#raid_level}.
	RaidLevel *string `field:"optional" json:"raidLevel" yaml:"raidLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_ganglia_layer#type OpsworksGangliaLayer#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

