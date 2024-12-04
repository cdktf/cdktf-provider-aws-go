// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchdomain


type OpensearchDomainSoftwareUpdateOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/opensearch_domain#auto_software_update_enabled OpensearchDomain#auto_software_update_enabled}.
	AutoSoftwareUpdateEnabled interface{} `field:"optional" json:"autoSoftwareUpdateEnabled" yaml:"autoSoftwareUpdateEnabled"`
}

