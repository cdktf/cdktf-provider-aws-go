// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontdistribution

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontDistributionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// default_cache_behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#default_cache_behavior CloudfrontDistribution#default_cache_behavior}
	DefaultCacheBehavior *CloudfrontDistributionDefaultCacheBehavior `field:"required" json:"defaultCacheBehavior" yaml:"defaultCacheBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#enabled CloudfrontDistribution#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// origin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#origin CloudfrontDistribution#origin}
	Origin interface{} `field:"required" json:"origin" yaml:"origin"`
	// restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#restrictions CloudfrontDistribution#restrictions}
	Restrictions *CloudfrontDistributionRestrictions `field:"required" json:"restrictions" yaml:"restrictions"`
	// viewer_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#viewer_certificate CloudfrontDistribution#viewer_certificate}
	ViewerCertificate *CloudfrontDistributionViewerCertificate `field:"required" json:"viewerCertificate" yaml:"viewerCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#aliases CloudfrontDistribution#aliases}.
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#anycast_ip_list_id CloudfrontDistribution#anycast_ip_list_id}.
	AnycastIpListId *string `field:"optional" json:"anycastIpListId" yaml:"anycastIpListId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#comment CloudfrontDistribution#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#continuous_deployment_policy_id CloudfrontDistribution#continuous_deployment_policy_id}.
	ContinuousDeploymentPolicyId *string `field:"optional" json:"continuousDeploymentPolicyId" yaml:"continuousDeploymentPolicyId"`
	// custom_error_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#custom_error_response CloudfrontDistribution#custom_error_response}
	CustomErrorResponse interface{} `field:"optional" json:"customErrorResponse" yaml:"customErrorResponse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#default_root_object CloudfrontDistribution#default_root_object}.
	DefaultRootObject *string `field:"optional" json:"defaultRootObject" yaml:"defaultRootObject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#http_version CloudfrontDistribution#http_version}.
	HttpVersion *string `field:"optional" json:"httpVersion" yaml:"httpVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#id CloudfrontDistribution#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#is_ipv6_enabled CloudfrontDistribution#is_ipv6_enabled}.
	IsIpv6Enabled interface{} `field:"optional" json:"isIpv6Enabled" yaml:"isIpv6Enabled"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#logging_config CloudfrontDistribution#logging_config}
	LoggingConfig *CloudfrontDistributionLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// ordered_cache_behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#ordered_cache_behavior CloudfrontDistribution#ordered_cache_behavior}
	OrderedCacheBehavior interface{} `field:"optional" json:"orderedCacheBehavior" yaml:"orderedCacheBehavior"`
	// origin_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#origin_group CloudfrontDistribution#origin_group}
	OriginGroup interface{} `field:"optional" json:"originGroup" yaml:"originGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#price_class CloudfrontDistribution#price_class}.
	PriceClass *string `field:"optional" json:"priceClass" yaml:"priceClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#retain_on_delete CloudfrontDistribution#retain_on_delete}.
	RetainOnDelete interface{} `field:"optional" json:"retainOnDelete" yaml:"retainOnDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#staging CloudfrontDistribution#staging}.
	Staging interface{} `field:"optional" json:"staging" yaml:"staging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#tags CloudfrontDistribution#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#tags_all CloudfrontDistribution#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#wait_for_deployment CloudfrontDistribution#wait_for_deployment}.
	WaitForDeployment interface{} `field:"optional" json:"waitForDeployment" yaml:"waitForDeployment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_distribution#web_acl_id CloudfrontDistribution#web_acl_id}.
	WebAclId *string `field:"optional" json:"webAclId" yaml:"webAclId"`
}

