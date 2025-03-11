// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appfabricappauthorizationconnection


type AppfabricAppAuthorizationConnectionTimeouts struct {
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/appfabric_app_authorization_connection#create AppfabricAppAuthorizationConnection#create}
	Create *string `field:"optional" json:"create" yaml:"create"`
}

