// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchdomain


type OpensearchDomainSnapshotOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/opensearch_domain#automated_snapshot_start_hour OpensearchDomain#automated_snapshot_start_hour}.
	AutomatedSnapshotStartHour *float64 `field:"required" json:"automatedSnapshotStartHour" yaml:"automatedSnapshotStartHour"`
}

