// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchdomain


type OpensearchDomainOffPeakWindowOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/opensearch_domain#enabled OpensearchDomain#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// off_peak_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/opensearch_domain#off_peak_window OpensearchDomain#off_peak_window}
	OffPeakWindow *OpensearchDomainOffPeakWindowOptionsOffPeakWindow `field:"optional" json:"offPeakWindow" yaml:"offPeakWindow"`
}

