// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsfsxontapstoragevirtualmachine


type DataAwsFsxOntapStorageVirtualMachineFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/fsx_ontap_storage_virtual_machine#name DataAwsFsxOntapStorageVirtualMachine#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/fsx_ontap_storage_virtual_machine#values DataAwsFsxOntapStorageVirtualMachine#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

