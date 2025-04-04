// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53recordsexclusive


type Route53RecordsExclusiveResourceRecordSetCidrRoutingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/route53_records_exclusive#collection_id Route53RecordsExclusive#collection_id}.
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/route53_records_exclusive#location_name Route53RecordsExclusive#location_name}.
	LocationName *string `field:"required" json:"locationName" yaml:"locationName"`
}

