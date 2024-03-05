// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontdistribution


type CloudfrontDistributionOrigin struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/cloudfront_distribution#domain_name CloudfrontDistribution#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/cloudfront_distribution#origin_id CloudfrontDistribution#origin_id}.
	OriginId *string `field:"required" json:"originId" yaml:"originId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/cloudfront_distribution#connection_attempts CloudfrontDistribution#connection_attempts}.
	ConnectionAttempts *float64 `field:"optional" json:"connectionAttempts" yaml:"connectionAttempts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/cloudfront_distribution#connection_timeout CloudfrontDistribution#connection_timeout}.
	ConnectionTimeout *float64 `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// custom_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/cloudfront_distribution#custom_header CloudfrontDistribution#custom_header}
	CustomHeader interface{} `field:"optional" json:"customHeader" yaml:"customHeader"`
	// custom_origin_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/cloudfront_distribution#custom_origin_config CloudfrontDistribution#custom_origin_config}
	CustomOriginConfig *CloudfrontDistributionOriginCustomOriginConfig `field:"optional" json:"customOriginConfig" yaml:"customOriginConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/cloudfront_distribution#origin_access_control_id CloudfrontDistribution#origin_access_control_id}.
	OriginAccessControlId *string `field:"optional" json:"originAccessControlId" yaml:"originAccessControlId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/cloudfront_distribution#origin_path CloudfrontDistribution#origin_path}.
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// origin_shield block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/cloudfront_distribution#origin_shield CloudfrontDistribution#origin_shield}
	OriginShield *CloudfrontDistributionOriginOriginShield `field:"optional" json:"originShield" yaml:"originShield"`
	// s3_origin_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/cloudfront_distribution#s3_origin_config CloudfrontDistribution#s3_origin_config}
	S3OriginConfig *CloudfrontDistributionOriginS3OriginConfig `field:"optional" json:"s3OriginConfig" yaml:"s3OriginConfig"`
}

