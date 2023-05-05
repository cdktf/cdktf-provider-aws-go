package cloudfrontdistribution


type CloudfrontDistributionOriginS3OriginConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/cloudfront_distribution#origin_access_identity CloudfrontDistribution#origin_access_identity}.
	OriginAccessIdentity *string `field:"required" json:"originAccessIdentity" yaml:"originAccessIdentity"`
}

