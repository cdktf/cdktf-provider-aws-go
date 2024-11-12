// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticsearchdomain


type ElasticsearchDomainAutoTuneOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/elasticsearch_domain#desired_state ElasticsearchDomain#desired_state}.
	DesiredState *string `field:"required" json:"desiredState" yaml:"desiredState"`
	// maintenance_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/elasticsearch_domain#maintenance_schedule ElasticsearchDomain#maintenance_schedule}
	MaintenanceSchedule interface{} `field:"optional" json:"maintenanceSchedule" yaml:"maintenanceSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/elasticsearch_domain#rollback_on_disable ElasticsearchDomain#rollback_on_disable}.
	RollbackOnDisable *string `field:"optional" json:"rollbackOnDisable" yaml:"rollbackOnDisable"`
}

