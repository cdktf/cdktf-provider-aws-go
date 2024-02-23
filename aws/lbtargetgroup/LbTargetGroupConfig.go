// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbtargetgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbTargetGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#connection_termination LbTargetGroup#connection_termination}.
	ConnectionTermination interface{} `field:"optional" json:"connectionTermination" yaml:"connectionTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#deregistration_delay LbTargetGroup#deregistration_delay}.
	DeregistrationDelay *string `field:"optional" json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#health_check LbTargetGroup#health_check}
	HealthCheck *LbTargetGroupHealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#id LbTargetGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#ip_address_type LbTargetGroup#ip_address_type}.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#lambda_multi_value_headers_enabled LbTargetGroup#lambda_multi_value_headers_enabled}.
	LambdaMultiValueHeadersEnabled interface{} `field:"optional" json:"lambdaMultiValueHeadersEnabled" yaml:"lambdaMultiValueHeadersEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#load_balancing_algorithm_type LbTargetGroup#load_balancing_algorithm_type}.
	LoadBalancingAlgorithmType *string `field:"optional" json:"loadBalancingAlgorithmType" yaml:"loadBalancingAlgorithmType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#load_balancing_anomaly_mitigation LbTargetGroup#load_balancing_anomaly_mitigation}.
	LoadBalancingAnomalyMitigation *string `field:"optional" json:"loadBalancingAnomalyMitigation" yaml:"loadBalancingAnomalyMitigation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#load_balancing_cross_zone_enabled LbTargetGroup#load_balancing_cross_zone_enabled}.
	LoadBalancingCrossZoneEnabled *string `field:"optional" json:"loadBalancingCrossZoneEnabled" yaml:"loadBalancingCrossZoneEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#name LbTargetGroup#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#name_prefix LbTargetGroup#name_prefix}.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#port LbTargetGroup#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#preserve_client_ip LbTargetGroup#preserve_client_ip}.
	PreserveClientIp *string `field:"optional" json:"preserveClientIp" yaml:"preserveClientIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#protocol LbTargetGroup#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#protocol_version LbTargetGroup#protocol_version}.
	ProtocolVersion *string `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#proxy_protocol_v2 LbTargetGroup#proxy_protocol_v2}.
	ProxyProtocolV2 interface{} `field:"optional" json:"proxyProtocolV2" yaml:"proxyProtocolV2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#slow_start LbTargetGroup#slow_start}.
	SlowStart *float64 `field:"optional" json:"slowStart" yaml:"slowStart"`
	// stickiness block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#stickiness LbTargetGroup#stickiness}
	Stickiness *LbTargetGroupStickiness `field:"optional" json:"stickiness" yaml:"stickiness"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#tags LbTargetGroup#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#tags_all LbTargetGroup#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// target_failover block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#target_failover LbTargetGroup#target_failover}
	TargetFailover interface{} `field:"optional" json:"targetFailover" yaml:"targetFailover"`
	// target_health_state block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#target_health_state LbTargetGroup#target_health_state}
	TargetHealthState interface{} `field:"optional" json:"targetHealthState" yaml:"targetHealthState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#target_type LbTargetGroup#target_type}.
	TargetType *string `field:"optional" json:"targetType" yaml:"targetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/lb_target_group#vpc_id LbTargetGroup#vpc_id}.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

