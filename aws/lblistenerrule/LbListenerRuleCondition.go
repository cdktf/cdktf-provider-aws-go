// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistenerrule


type LbListenerRuleCondition struct {
	// host_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/lb_listener_rule#host_header LbListenerRule#host_header}
	HostHeader *LbListenerRuleConditionHostHeader `field:"optional" json:"hostHeader" yaml:"hostHeader"`
	// http_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/lb_listener_rule#http_header LbListenerRule#http_header}
	HttpHeader *LbListenerRuleConditionHttpHeader `field:"optional" json:"httpHeader" yaml:"httpHeader"`
	// http_request_method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/lb_listener_rule#http_request_method LbListenerRule#http_request_method}
	HttpRequestMethod *LbListenerRuleConditionHttpRequestMethod `field:"optional" json:"httpRequestMethod" yaml:"httpRequestMethod"`
	// path_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/lb_listener_rule#path_pattern LbListenerRule#path_pattern}
	PathPattern *LbListenerRuleConditionPathPattern `field:"optional" json:"pathPattern" yaml:"pathPattern"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/lb_listener_rule#query_string LbListenerRule#query_string}
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// source_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/lb_listener_rule#source_ip LbListenerRule#source_ip}
	SourceIp *LbListenerRuleConditionSourceIp `field:"optional" json:"sourceIp" yaml:"sourceIp"`
}

