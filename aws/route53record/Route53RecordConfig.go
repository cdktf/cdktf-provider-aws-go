// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53record

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53RecordConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/route53_record#name Route53Record#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/route53_record#type Route53Record#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/route53_record#zone_id Route53Record#zone_id}.
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// alias block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/route53_record#alias Route53Record#alias}
	Alias *Route53RecordAlias `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/route53_record#allow_overwrite Route53Record#allow_overwrite}.
	AllowOverwrite interface{} `field:"optional" json:"allowOverwrite" yaml:"allowOverwrite"`
	// cidr_routing_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/route53_record#cidr_routing_policy Route53Record#cidr_routing_policy}
	CidrRoutingPolicy *Route53RecordCidrRoutingPolicy `field:"optional" json:"cidrRoutingPolicy" yaml:"cidrRoutingPolicy"`
	// failover_routing_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/route53_record#failover_routing_policy Route53Record#failover_routing_policy}
	FailoverRoutingPolicy *Route53RecordFailoverRoutingPolicy `field:"optional" json:"failoverRoutingPolicy" yaml:"failoverRoutingPolicy"`
	// geolocation_routing_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/route53_record#geolocation_routing_policy Route53Record#geolocation_routing_policy}
	GeolocationRoutingPolicy *Route53RecordGeolocationRoutingPolicy `field:"optional" json:"geolocationRoutingPolicy" yaml:"geolocationRoutingPolicy"`
	// geoproximity_routing_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/route53_record#geoproximity_routing_policy Route53Record#geoproximity_routing_policy}
	GeoproximityRoutingPolicy *Route53RecordGeoproximityRoutingPolicy `field:"optional" json:"geoproximityRoutingPolicy" yaml:"geoproximityRoutingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/route53_record#health_check_id Route53Record#health_check_id}.
	HealthCheckId *string `field:"optional" json:"healthCheckId" yaml:"healthCheckId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/route53_record#id Route53Record#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// latency_routing_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/route53_record#latency_routing_policy Route53Record#latency_routing_policy}
	LatencyRoutingPolicy *Route53RecordLatencyRoutingPolicy `field:"optional" json:"latencyRoutingPolicy" yaml:"latencyRoutingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/route53_record#multivalue_answer_routing_policy Route53Record#multivalue_answer_routing_policy}.
	MultivalueAnswerRoutingPolicy interface{} `field:"optional" json:"multivalueAnswerRoutingPolicy" yaml:"multivalueAnswerRoutingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/route53_record#records Route53Record#records}.
	Records *[]*string `field:"optional" json:"records" yaml:"records"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/route53_record#set_identifier Route53Record#set_identifier}.
	SetIdentifier *string `field:"optional" json:"setIdentifier" yaml:"setIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/route53_record#ttl Route53Record#ttl}.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
	// weighted_routing_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/route53_record#weighted_routing_policy Route53Record#weighted_routing_policy}
	WeightedRoutingPolicy *Route53RecordWeightedRoutingPolicy `field:"optional" json:"weightedRoutingPolicy" yaml:"weightedRoutingPolicy"`
}

