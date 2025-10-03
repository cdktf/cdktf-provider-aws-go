// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lightsaildistribution

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LightsailDistributionConfig struct {
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
	// The bundle ID to use for the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/lightsail_distribution#bundle_id LightsailDistribution#bundle_id}
	BundleId *string `field:"required" json:"bundleId" yaml:"bundleId"`
	// default_cache_behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/lightsail_distribution#default_cache_behavior LightsailDistribution#default_cache_behavior}
	DefaultCacheBehavior *LightsailDistributionDefaultCacheBehavior `field:"required" json:"defaultCacheBehavior" yaml:"defaultCacheBehavior"`
	// The name of the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/lightsail_distribution#name LightsailDistribution#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// origin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/lightsail_distribution#origin LightsailDistribution#origin}
	Origin *LightsailDistributionOrigin `field:"required" json:"origin" yaml:"origin"`
	// cache_behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/lightsail_distribution#cache_behavior LightsailDistribution#cache_behavior}
	CacheBehavior interface{} `field:"optional" json:"cacheBehavior" yaml:"cacheBehavior"`
	// cache_behavior_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/lightsail_distribution#cache_behavior_settings LightsailDistribution#cache_behavior_settings}
	CacheBehaviorSettings *LightsailDistributionCacheBehaviorSettings `field:"optional" json:"cacheBehaviorSettings" yaml:"cacheBehaviorSettings"`
	// The name of the SSL/TLS certificate attached to the distribution, if any.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/lightsail_distribution#certificate_name LightsailDistribution#certificate_name}
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/lightsail_distribution#id LightsailDistribution#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The IP address type of the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/lightsail_distribution#ip_address_type LightsailDistribution#ip_address_type}
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Indicates whether the distribution is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/lightsail_distribution#is_enabled LightsailDistribution#is_enabled}
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/lightsail_distribution#region LightsailDistribution#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/lightsail_distribution#tags LightsailDistribution#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/lightsail_distribution#tags_all LightsailDistribution#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/lightsail_distribution#timeouts LightsailDistribution#timeouts}
	Timeouts *LightsailDistributionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

