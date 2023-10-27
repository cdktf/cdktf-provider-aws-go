// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dbproxydefaulttargetgroup


type DbProxyDefaultTargetGroupConnectionPoolConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/db_proxy_default_target_group#connection_borrow_timeout DbProxyDefaultTargetGroup#connection_borrow_timeout}.
	ConnectionBorrowTimeout *float64 `field:"optional" json:"connectionBorrowTimeout" yaml:"connectionBorrowTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/db_proxy_default_target_group#init_query DbProxyDefaultTargetGroup#init_query}.
	InitQuery *string `field:"optional" json:"initQuery" yaml:"initQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/db_proxy_default_target_group#max_connections_percent DbProxyDefaultTargetGroup#max_connections_percent}.
	MaxConnectionsPercent *float64 `field:"optional" json:"maxConnectionsPercent" yaml:"maxConnectionsPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/db_proxy_default_target_group#max_idle_connections_percent DbProxyDefaultTargetGroup#max_idle_connections_percent}.
	MaxIdleConnectionsPercent *float64 `field:"optional" json:"maxIdleConnectionsPercent" yaml:"maxIdleConnectionsPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/db_proxy_default_target_group#session_pinning_filters DbProxyDefaultTargetGroup#session_pinning_filters}.
	SessionPinningFilters *[]*string `field:"optional" json:"sessionPinningFilters" yaml:"sessionPinningFilters"`
}

