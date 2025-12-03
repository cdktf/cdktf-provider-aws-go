// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionTenancyConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/lambda_function#tenant_isolation_mode LambdaFunction#tenant_isolation_mode}.
	TenantIsolationMode *string `field:"required" json:"tenantIsolationMode" yaml:"tenantIsolationMode"`
}

