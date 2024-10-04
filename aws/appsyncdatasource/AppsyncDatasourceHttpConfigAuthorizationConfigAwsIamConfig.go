// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncdatasource


type AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/appsync_datasource#signing_region AppsyncDatasource#signing_region}.
	SigningRegion *string `field:"optional" json:"signingRegion" yaml:"signingRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/appsync_datasource#signing_service_name AppsyncDatasource#signing_service_name}.
	SigningServiceName *string `field:"optional" json:"signingServiceName" yaml:"signingServiceName"`
}

