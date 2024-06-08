// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alb

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbConfig struct {
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
	// access_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#access_logs Alb#access_logs}
	AccessLogs *AlbAccessLogs `field:"optional" json:"accessLogs" yaml:"accessLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#client_keep_alive Alb#client_keep_alive}.
	ClientKeepAlive *float64 `field:"optional" json:"clientKeepAlive" yaml:"clientKeepAlive"`
	// connection_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#connection_logs Alb#connection_logs}
	ConnectionLogs *AlbConnectionLogs `field:"optional" json:"connectionLogs" yaml:"connectionLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#customer_owned_ipv4_pool Alb#customer_owned_ipv4_pool}.
	CustomerOwnedIpv4Pool *string `field:"optional" json:"customerOwnedIpv4Pool" yaml:"customerOwnedIpv4Pool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#desync_mitigation_mode Alb#desync_mitigation_mode}.
	DesyncMitigationMode *string `field:"optional" json:"desyncMitigationMode" yaml:"desyncMitigationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#dns_record_client_routing_policy Alb#dns_record_client_routing_policy}.
	DnsRecordClientRoutingPolicy *string `field:"optional" json:"dnsRecordClientRoutingPolicy" yaml:"dnsRecordClientRoutingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#drop_invalid_header_fields Alb#drop_invalid_header_fields}.
	DropInvalidHeaderFields interface{} `field:"optional" json:"dropInvalidHeaderFields" yaml:"dropInvalidHeaderFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#enable_cross_zone_load_balancing Alb#enable_cross_zone_load_balancing}.
	EnableCrossZoneLoadBalancing interface{} `field:"optional" json:"enableCrossZoneLoadBalancing" yaml:"enableCrossZoneLoadBalancing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#enable_deletion_protection Alb#enable_deletion_protection}.
	EnableDeletionProtection interface{} `field:"optional" json:"enableDeletionProtection" yaml:"enableDeletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#enable_http2 Alb#enable_http2}.
	EnableHttp2 interface{} `field:"optional" json:"enableHttp2" yaml:"enableHttp2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#enable_tls_version_and_cipher_suite_headers Alb#enable_tls_version_and_cipher_suite_headers}.
	EnableTlsVersionAndCipherSuiteHeaders interface{} `field:"optional" json:"enableTlsVersionAndCipherSuiteHeaders" yaml:"enableTlsVersionAndCipherSuiteHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#enable_waf_fail_open Alb#enable_waf_fail_open}.
	EnableWafFailOpen interface{} `field:"optional" json:"enableWafFailOpen" yaml:"enableWafFailOpen"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#enable_xff_client_port Alb#enable_xff_client_port}.
	EnableXffClientPort interface{} `field:"optional" json:"enableXffClientPort" yaml:"enableXffClientPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#enforce_security_group_inbound_rules_on_private_link_traffic Alb#enforce_security_group_inbound_rules_on_private_link_traffic}.
	EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic *string `field:"optional" json:"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic" yaml:"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#id Alb#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#idle_timeout Alb#idle_timeout}.
	IdleTimeout *float64 `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#internal Alb#internal}.
	Internal interface{} `field:"optional" json:"internal" yaml:"internal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#ip_address_type Alb#ip_address_type}.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#load_balancer_type Alb#load_balancer_type}.
	LoadBalancerType *string `field:"optional" json:"loadBalancerType" yaml:"loadBalancerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#name Alb#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#name_prefix Alb#name_prefix}.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#preserve_host_header Alb#preserve_host_header}.
	PreserveHostHeader interface{} `field:"optional" json:"preserveHostHeader" yaml:"preserveHostHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#security_groups Alb#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// subnet_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#subnet_mapping Alb#subnet_mapping}
	SubnetMapping interface{} `field:"optional" json:"subnetMapping" yaml:"subnetMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#subnets Alb#subnets}.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#tags Alb#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#tags_all Alb#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#timeouts Alb#timeouts}
	Timeouts *AlbTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/alb#xff_header_processing_mode Alb#xff_header_processing_mode}.
	XffHeaderProcessingMode *string `field:"optional" json:"xffHeaderProcessingMode" yaml:"xffHeaderProcessingMode"`
}

