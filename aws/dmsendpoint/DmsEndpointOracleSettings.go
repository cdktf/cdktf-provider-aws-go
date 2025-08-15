// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint


type DmsEndpointOracleSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/dms_endpoint#authentication_method DmsEndpoint#authentication_method}.
	AuthenticationMethod *string `field:"optional" json:"authenticationMethod" yaml:"authenticationMethod"`
}

