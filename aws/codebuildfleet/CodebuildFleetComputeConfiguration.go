// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codebuildfleet


type CodebuildFleetComputeConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/codebuild_fleet#disk CodebuildFleet#disk}.
	Disk *float64 `field:"optional" json:"disk" yaml:"disk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/codebuild_fleet#instance_type CodebuildFleet#instance_type}.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/codebuild_fleet#machine_type CodebuildFleet#machine_type}.
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/codebuild_fleet#memory CodebuildFleet#memory}.
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/codebuild_fleet#vcpu CodebuildFleet#vcpu}.
	Vcpu *float64 `field:"optional" json:"vcpu" yaml:"vcpu"`
}

