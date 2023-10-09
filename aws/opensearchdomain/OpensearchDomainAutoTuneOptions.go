// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchdomain


type OpensearchDomainAutoTuneOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/opensearch_domain#desired_state OpensearchDomain#desired_state}.
	DesiredState *string `field:"required" json:"desiredState" yaml:"desiredState"`
	// maintenance_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/opensearch_domain#maintenance_schedule OpensearchDomain#maintenance_schedule}
	MaintenanceSchedule interface{} `field:"optional" json:"maintenanceSchedule" yaml:"maintenanceSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/opensearch_domain#rollback_on_disable OpensearchDomain#rollback_on_disable}.
	RollbackOnDisable *string `field:"optional" json:"rollbackOnDisable" yaml:"rollbackOnDisable"`
}

