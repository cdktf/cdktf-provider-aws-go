// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasyncagent


type DatasyncAgentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/datasync_agent#create DatasyncAgent#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

