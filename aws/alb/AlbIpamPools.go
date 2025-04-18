// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alb


type AlbIpamPools struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/alb#ipv4_ipam_pool_id Alb#ipv4_ipam_pool_id}.
	Ipv4IpamPoolId *string `field:"required" json:"ipv4IpamPoolId" yaml:"ipv4IpamPoolId"`
}

