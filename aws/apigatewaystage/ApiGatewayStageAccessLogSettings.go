// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewaystage


type ApiGatewayStageAccessLogSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/api_gateway_stage#destination_arn ApiGatewayStage#destination_arn}.
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/api_gateway_stage#format ApiGatewayStage#format}.
	Format *string `field:"required" json:"format" yaml:"format"`
}

