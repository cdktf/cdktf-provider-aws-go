// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchdomain


type OpensearchDomainOffPeakWindowOptionsOffPeakWindowWindowStartTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opensearch_domain#hours OpensearchDomain#hours}.
	Hours *float64 `field:"optional" json:"hours" yaml:"hours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/opensearch_domain#minutes OpensearchDomain#minutes}.
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
}

