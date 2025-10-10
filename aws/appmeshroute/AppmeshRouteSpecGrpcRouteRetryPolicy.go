// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshroute


type AppmeshRouteSpecGrpcRouteRetryPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appmesh_route#max_retries AppmeshRoute#max_retries}.
	MaxRetries *float64 `field:"required" json:"maxRetries" yaml:"maxRetries"`
	// per_retry_timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appmesh_route#per_retry_timeout AppmeshRoute#per_retry_timeout}
	PerRetryTimeout *AppmeshRouteSpecGrpcRouteRetryPolicyPerRetryTimeout `field:"required" json:"perRetryTimeout" yaml:"perRetryTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appmesh_route#grpc_retry_events AppmeshRoute#grpc_retry_events}.
	GrpcRetryEvents *[]*string `field:"optional" json:"grpcRetryEvents" yaml:"grpcRetryEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appmesh_route#http_retry_events AppmeshRoute#http_retry_events}.
	HttpRetryEvents *[]*string `field:"optional" json:"httpRetryEvents" yaml:"httpRetryEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appmesh_route#tcp_retry_events AppmeshRoute#tcp_retry_events}.
	TcpRetryEvents *[]*string `field:"optional" json:"tcpRetryEvents" yaml:"tcpRetryEvents"`
}

