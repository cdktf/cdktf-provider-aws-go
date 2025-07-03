// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistener


type LbListenerDefaultActionForwardStickiness struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/lb_listener#duration LbListener#duration}.
	Duration *float64 `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/lb_listener#enabled LbListener#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

