// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dbproxydefaulttargetgroup


type DbProxyDefaultTargetGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.0/docs/resources/db_proxy_default_target_group#create DbProxyDefaultTargetGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.0/docs/resources/db_proxy_default_target_group#update DbProxyDefaultTargetGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

