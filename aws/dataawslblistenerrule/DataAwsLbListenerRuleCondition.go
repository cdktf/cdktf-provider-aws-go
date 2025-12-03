// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawslblistenerrule


type DataAwsLbListenerRuleCondition struct {
	// host_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/data-sources/lb_listener_rule#host_header DataAwsLbListenerRule#host_header}
	HostHeader interface{} `field:"optional" json:"hostHeader" yaml:"hostHeader"`
	// http_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/data-sources/lb_listener_rule#http_header DataAwsLbListenerRule#http_header}
	HttpHeader interface{} `field:"optional" json:"httpHeader" yaml:"httpHeader"`
	// http_request_method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/data-sources/lb_listener_rule#http_request_method DataAwsLbListenerRule#http_request_method}
	HttpRequestMethod interface{} `field:"optional" json:"httpRequestMethod" yaml:"httpRequestMethod"`
	// path_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/data-sources/lb_listener_rule#path_pattern DataAwsLbListenerRule#path_pattern}
	PathPattern interface{} `field:"optional" json:"pathPattern" yaml:"pathPattern"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/data-sources/lb_listener_rule#query_string DataAwsLbListenerRule#query_string}
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// source_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/data-sources/lb_listener_rule#source_ip DataAwsLbListenerRule#source_ip}
	SourceIp interface{} `field:"optional" json:"sourceIp" yaml:"sourceIp"`
}

