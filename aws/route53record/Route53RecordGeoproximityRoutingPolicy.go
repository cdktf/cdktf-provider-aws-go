// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53record


type Route53RecordGeoproximityRoutingPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/route53_record#aws_region Route53Record#aws_region}.
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/route53_record#bias Route53Record#bias}.
	Bias *float64 `field:"optional" json:"bias" yaml:"bias"`
	// coordinates block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/route53_record#coordinates Route53Record#coordinates}
	Coordinates interface{} `field:"optional" json:"coordinates" yaml:"coordinates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/route53_record#local_zone_group Route53Record#local_zone_group}.
	LocalZoneGroup *string `field:"optional" json:"localZoneGroup" yaml:"localZoneGroup"`
}

