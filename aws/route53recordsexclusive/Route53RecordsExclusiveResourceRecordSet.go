// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53recordsexclusive


type Route53RecordsExclusiveResourceRecordSet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/route53_records_exclusive#name Route53RecordsExclusive#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// alias_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/route53_records_exclusive#alias_target Route53RecordsExclusive#alias_target}
	AliasTarget interface{} `field:"optional" json:"aliasTarget" yaml:"aliasTarget"`
	// cidr_routing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/route53_records_exclusive#cidr_routing_config Route53RecordsExclusive#cidr_routing_config}
	CidrRoutingConfig interface{} `field:"optional" json:"cidrRoutingConfig" yaml:"cidrRoutingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/route53_records_exclusive#failover Route53RecordsExclusive#failover}.
	Failover *string `field:"optional" json:"failover" yaml:"failover"`
	// geolocation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/route53_records_exclusive#geolocation Route53RecordsExclusive#geolocation}
	Geolocation interface{} `field:"optional" json:"geolocation" yaml:"geolocation"`
	// geoproximity_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/route53_records_exclusive#geoproximity_location Route53RecordsExclusive#geoproximity_location}
	GeoproximityLocation interface{} `field:"optional" json:"geoproximityLocation" yaml:"geoproximityLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/route53_records_exclusive#health_check_id Route53RecordsExclusive#health_check_id}.
	HealthCheckId *string `field:"optional" json:"healthCheckId" yaml:"healthCheckId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/route53_records_exclusive#multi_value_answer Route53RecordsExclusive#multi_value_answer}.
	MultiValueAnswer interface{} `field:"optional" json:"multiValueAnswer" yaml:"multiValueAnswer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/route53_records_exclusive#region Route53RecordsExclusive#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// resource_records block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/route53_records_exclusive#resource_records Route53RecordsExclusive#resource_records}
	ResourceRecords interface{} `field:"optional" json:"resourceRecords" yaml:"resourceRecords"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/route53_records_exclusive#set_identifier Route53RecordsExclusive#set_identifier}.
	SetIdentifier *string `field:"optional" json:"setIdentifier" yaml:"setIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/route53_records_exclusive#traffic_policy_instance_id Route53RecordsExclusive#traffic_policy_instance_id}.
	TrafficPolicyInstanceId *string `field:"optional" json:"trafficPolicyInstanceId" yaml:"trafficPolicyInstanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/route53_records_exclusive#ttl Route53RecordsExclusive#ttl}.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/route53_records_exclusive#type Route53RecordsExclusive#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/route53_records_exclusive#weight Route53RecordsExclusive#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

