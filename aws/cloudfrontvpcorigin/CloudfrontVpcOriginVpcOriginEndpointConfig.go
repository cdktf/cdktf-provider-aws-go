// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontvpcorigin


type CloudfrontVpcOriginVpcOriginEndpointConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudfront_vpc_origin#arn CloudfrontVpcOrigin#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudfront_vpc_origin#http_port CloudfrontVpcOrigin#http_port}.
	HttpPort *float64 `field:"required" json:"httpPort" yaml:"httpPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudfront_vpc_origin#https_port CloudfrontVpcOrigin#https_port}.
	HttpsPort *float64 `field:"required" json:"httpsPort" yaml:"httpsPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudfront_vpc_origin#name CloudfrontVpcOrigin#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudfront_vpc_origin#origin_protocol_policy CloudfrontVpcOrigin#origin_protocol_policy}.
	OriginProtocolPolicy *string `field:"required" json:"originProtocolPolicy" yaml:"originProtocolPolicy"`
	// origin_ssl_protocols block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudfront_vpc_origin#origin_ssl_protocols CloudfrontVpcOrigin#origin_ssl_protocols}
	OriginSslProtocols interface{} `field:"optional" json:"originSslProtocols" yaml:"originSslProtocols"`
}

