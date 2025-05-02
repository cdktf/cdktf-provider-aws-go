// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatcheventconnection


type CloudwatchEventConnectionAuthParametersOauthOauthHttpParameters struct {
	// body block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/cloudwatch_event_connection#body CloudwatchEventConnection#body}
	Body interface{} `field:"optional" json:"body" yaml:"body"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/cloudwatch_event_connection#header CloudwatchEventConnection#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/cloudwatch_event_connection#query_string CloudwatchEventConnection#query_string}
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
}

