// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchdomain


type OpensearchDomainAutoTuneOptionsMaintenanceScheduleDuration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/opensearch_domain#unit OpensearchDomain#unit}.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/opensearch_domain#value OpensearchDomain#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

