// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53recordsexclusive


type Route53RecordsExclusiveResourceRecordSetAliasTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/route53_records_exclusive#dns_name Route53RecordsExclusive#dns_name}.
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/route53_records_exclusive#evaluate_target_health Route53RecordsExclusive#evaluate_target_health}.
	EvaluateTargetHealth interface{} `field:"required" json:"evaluateTargetHealth" yaml:"evaluateTargetHealth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/route53_records_exclusive#hosted_zone_id Route53RecordsExclusive#hosted_zone_id}.
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
}

