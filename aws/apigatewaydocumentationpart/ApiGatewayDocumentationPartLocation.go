// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewaydocumentationpart


type ApiGatewayDocumentationPartLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/api_gateway_documentation_part#type ApiGatewayDocumentationPart#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/api_gateway_documentation_part#method ApiGatewayDocumentationPart#method}.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/api_gateway_documentation_part#name ApiGatewayDocumentationPart#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/api_gateway_documentation_part#path ApiGatewayDocumentationPart#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/api_gateway_documentation_part#status_code ApiGatewayDocumentationPart#status_code}.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
}

