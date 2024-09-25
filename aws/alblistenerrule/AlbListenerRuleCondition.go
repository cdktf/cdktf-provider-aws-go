// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alblistenerrule


type AlbListenerRuleCondition struct {
	// host_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/alb_listener_rule#host_header AlbListenerRule#host_header}
	HostHeader *AlbListenerRuleConditionHostHeader `field:"optional" json:"hostHeader" yaml:"hostHeader"`
	// http_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/alb_listener_rule#http_header AlbListenerRule#http_header}
	HttpHeader *AlbListenerRuleConditionHttpHeader `field:"optional" json:"httpHeader" yaml:"httpHeader"`
	// http_request_method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/alb_listener_rule#http_request_method AlbListenerRule#http_request_method}
	HttpRequestMethod *AlbListenerRuleConditionHttpRequestMethod `field:"optional" json:"httpRequestMethod" yaml:"httpRequestMethod"`
	// path_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/alb_listener_rule#path_pattern AlbListenerRule#path_pattern}
	PathPattern *AlbListenerRuleConditionPathPattern `field:"optional" json:"pathPattern" yaml:"pathPattern"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/alb_listener_rule#query_string AlbListenerRule#query_string}
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// source_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/alb_listener_rule#source_ip AlbListenerRule#source_ip}
	SourceIp *AlbListenerRuleConditionSourceIp `field:"optional" json:"sourceIp" yaml:"sourceIp"`
}

