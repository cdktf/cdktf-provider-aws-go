// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagegatewaygateway


type StoragegatewayGatewaySmbActiveDirectorySettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/storagegateway_gateway#domain_name StoragegatewayGateway#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/storagegateway_gateway#password StoragegatewayGateway#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/storagegateway_gateway#username StoragegatewayGateway#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/storagegateway_gateway#domain_controllers StoragegatewayGateway#domain_controllers}.
	DomainControllers *[]*string `field:"optional" json:"domainControllers" yaml:"domainControllers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/storagegateway_gateway#organizational_unit StoragegatewayGateway#organizational_unit}.
	OrganizationalUnit *string `field:"optional" json:"organizationalUnit" yaml:"organizationalUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/storagegateway_gateway#timeout_in_seconds StoragegatewayGateway#timeout_in_seconds}.
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}

