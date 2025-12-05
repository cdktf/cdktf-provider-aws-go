// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package shielddrtaccesslogbucketassociation


type ShieldDrtAccessLogBucketAssociationTimeouts struct {
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/shield_drt_access_log_bucket_association#create ShieldDrtAccessLogBucketAssociation#create}
	Create *string `field:"optional" json:"create" yaml:"create"`
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/shield_drt_access_log_bucket_association#delete ShieldDrtAccessLogBucketAssociation#delete}
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

