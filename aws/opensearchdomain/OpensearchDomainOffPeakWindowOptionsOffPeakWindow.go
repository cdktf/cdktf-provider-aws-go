// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchdomain


type OpensearchDomainOffPeakWindowOptionsOffPeakWindow struct {
	// window_start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/opensearch_domain#window_start_time OpensearchDomain#window_start_time}
	WindowStartTime *OpensearchDomainOffPeakWindowOptionsOffPeakWindowWindowStartTime `field:"optional" json:"windowStartTime" yaml:"windowStartTime"`
}

