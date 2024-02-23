// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksnodejsapplayer


type OpsworksNodejsAppLayerEbsVolume struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/opsworks_nodejs_app_layer#mount_point OpsworksNodejsAppLayer#mount_point}.
	MountPoint *string `field:"required" json:"mountPoint" yaml:"mountPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/opsworks_nodejs_app_layer#number_of_disks OpsworksNodejsAppLayer#number_of_disks}.
	NumberOfDisks *float64 `field:"required" json:"numberOfDisks" yaml:"numberOfDisks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/opsworks_nodejs_app_layer#size OpsworksNodejsAppLayer#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/opsworks_nodejs_app_layer#encrypted OpsworksNodejsAppLayer#encrypted}.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/opsworks_nodejs_app_layer#iops OpsworksNodejsAppLayer#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/opsworks_nodejs_app_layer#raid_level OpsworksNodejsAppLayer#raid_level}.
	RaidLevel *string `field:"optional" json:"raidLevel" yaml:"raidLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/opsworks_nodejs_app_layer#type OpsworksNodejsAppLayer#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

