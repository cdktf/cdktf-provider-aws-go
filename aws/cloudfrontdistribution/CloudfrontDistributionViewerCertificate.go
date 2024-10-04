// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontdistribution


type CloudfrontDistributionViewerCertificate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/cloudfront_distribution#acm_certificate_arn CloudfrontDistribution#acm_certificate_arn}.
	AcmCertificateArn *string `field:"optional" json:"acmCertificateArn" yaml:"acmCertificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/cloudfront_distribution#cloudfront_default_certificate CloudfrontDistribution#cloudfront_default_certificate}.
	CloudfrontDefaultCertificate interface{} `field:"optional" json:"cloudfrontDefaultCertificate" yaml:"cloudfrontDefaultCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/cloudfront_distribution#iam_certificate_id CloudfrontDistribution#iam_certificate_id}.
	IamCertificateId *string `field:"optional" json:"iamCertificateId" yaml:"iamCertificateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/cloudfront_distribution#minimum_protocol_version CloudfrontDistribution#minimum_protocol_version}.
	MinimumProtocolVersion *string `field:"optional" json:"minimumProtocolVersion" yaml:"minimumProtocolVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/cloudfront_distribution#ssl_support_method CloudfrontDistribution#ssl_support_method}.
	SslSupportMethod *string `field:"optional" json:"sslSupportMethod" yaml:"sslSupportMethod"`
}

