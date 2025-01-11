// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRuleRuleDefinitionMatchAttributes struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/networkfirewall_rule_group#destination NetworkfirewallRuleGroup#destination}
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// destination_port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/networkfirewall_rule_group#destination_port NetworkfirewallRuleGroup#destination_port}
	DestinationPort interface{} `field:"optional" json:"destinationPort" yaml:"destinationPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/networkfirewall_rule_group#protocols NetworkfirewallRuleGroup#protocols}.
	Protocols *[]*float64 `field:"optional" json:"protocols" yaml:"protocols"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/networkfirewall_rule_group#source NetworkfirewallRuleGroup#source}
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// source_port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/networkfirewall_rule_group#source_port NetworkfirewallRuleGroup#source_port}
	SourcePort interface{} `field:"optional" json:"sourcePort" yaml:"sourcePort"`
	// tcp_flag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/networkfirewall_rule_group#tcp_flag NetworkfirewallRuleGroup#tcp_flag}
	TcpFlag interface{} `field:"optional" json:"tcpFlag" yaml:"tcpFlag"`
}

