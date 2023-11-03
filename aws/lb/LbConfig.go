// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lb

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#access_logs Lb#access_logs}
	AccessLogs *LbAccessLogs `field:"optional" json:"accessLogs" yaml:"accessLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#customer_owned_ipv4_pool Lb#customer_owned_ipv4_pool}.
	CustomerOwnedIpv4Pool *string `field:"optional" json:"customerOwnedIpv4Pool" yaml:"customerOwnedIpv4Pool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#desync_mitigation_mode Lb#desync_mitigation_mode}.
	DesyncMitigationMode *string `field:"optional" json:"desyncMitigationMode" yaml:"desyncMitigationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#dns_record_client_routing_policy Lb#dns_record_client_routing_policy}.
	DnsRecordClientRoutingPolicy *string `field:"optional" json:"dnsRecordClientRoutingPolicy" yaml:"dnsRecordClientRoutingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#drop_invalid_header_fields Lb#drop_invalid_header_fields}.
	DropInvalidHeaderFields interface{} `field:"optional" json:"dropInvalidHeaderFields" yaml:"dropInvalidHeaderFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#enable_cross_zone_load_balancing Lb#enable_cross_zone_load_balancing}.
	EnableCrossZoneLoadBalancing interface{} `field:"optional" json:"enableCrossZoneLoadBalancing" yaml:"enableCrossZoneLoadBalancing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#enable_deletion_protection Lb#enable_deletion_protection}.
	EnableDeletionProtection interface{} `field:"optional" json:"enableDeletionProtection" yaml:"enableDeletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#enable_http2 Lb#enable_http2}.
	EnableHttp2 interface{} `field:"optional" json:"enableHttp2" yaml:"enableHttp2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#enable_tls_version_and_cipher_suite_headers Lb#enable_tls_version_and_cipher_suite_headers}.
	EnableTlsVersionAndCipherSuiteHeaders interface{} `field:"optional" json:"enableTlsVersionAndCipherSuiteHeaders" yaml:"enableTlsVersionAndCipherSuiteHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#enable_waf_fail_open Lb#enable_waf_fail_open}.
	EnableWafFailOpen interface{} `field:"optional" json:"enableWafFailOpen" yaml:"enableWafFailOpen"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#enable_xff_client_port Lb#enable_xff_client_port}.
	EnableXffClientPort interface{} `field:"optional" json:"enableXffClientPort" yaml:"enableXffClientPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#id Lb#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#idle_timeout Lb#idle_timeout}.
	IdleTimeout *float64 `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#internal Lb#internal}.
	Internal interface{} `field:"optional" json:"internal" yaml:"internal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#ip_address_type Lb#ip_address_type}.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#load_balancer_type Lb#load_balancer_type}.
	LoadBalancerType *string `field:"optional" json:"loadBalancerType" yaml:"loadBalancerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#name Lb#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#name_prefix Lb#name_prefix}.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#preserve_host_header Lb#preserve_host_header}.
	PreserveHostHeader interface{} `field:"optional" json:"preserveHostHeader" yaml:"preserveHostHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#security_groups Lb#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// subnet_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#subnet_mapping Lb#subnet_mapping}
	SubnetMapping interface{} `field:"optional" json:"subnetMapping" yaml:"subnetMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#subnets Lb#subnets}.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#tags Lb#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#tags_all Lb#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#timeouts Lb#timeouts}
	Timeouts *LbTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/lb#xff_header_processing_mode Lb#xff_header_processing_mode}.
	XffHeaderProcessingMode *string `field:"optional" json:"xffHeaderProcessingMode" yaml:"xffHeaderProcessingMode"`
}

