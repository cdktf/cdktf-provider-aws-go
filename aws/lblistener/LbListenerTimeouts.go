// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistener


type LbListenerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/lb_listener#read LbListener#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

