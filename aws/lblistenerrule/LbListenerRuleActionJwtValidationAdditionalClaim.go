// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistenerrule


type LbListenerRuleActionJwtValidationAdditionalClaim struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/lb_listener_rule#format LbListenerRule#format}.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/lb_listener_rule#name LbListenerRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/lb_listener_rule#values LbListenerRule#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

