// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2clientvpnendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2ClientVpnEndpointConfig struct {
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
	// authentication_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ec2_client_vpn_endpoint#authentication_options Ec2ClientVpnEndpoint#authentication_options}
	AuthenticationOptions interface{} `field:"required" json:"authenticationOptions" yaml:"authenticationOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ec2_client_vpn_endpoint#client_cidr_block Ec2ClientVpnEndpoint#client_cidr_block}.
	ClientCidrBlock *string `field:"required" json:"clientCidrBlock" yaml:"clientCidrBlock"`
	// connection_log_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ec2_client_vpn_endpoint#connection_log_options Ec2ClientVpnEndpoint#connection_log_options}
	ConnectionLogOptions *Ec2ClientVpnEndpointConnectionLogOptions `field:"required" json:"connectionLogOptions" yaml:"connectionLogOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ec2_client_vpn_endpoint#server_certificate_arn Ec2ClientVpnEndpoint#server_certificate_arn}.
	ServerCertificateArn *string `field:"required" json:"serverCertificateArn" yaml:"serverCertificateArn"`
	// client_connect_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ec2_client_vpn_endpoint#client_connect_options Ec2ClientVpnEndpoint#client_connect_options}
	ClientConnectOptions *Ec2ClientVpnEndpointClientConnectOptions `field:"optional" json:"clientConnectOptions" yaml:"clientConnectOptions"`
	// client_login_banner_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ec2_client_vpn_endpoint#client_login_banner_options Ec2ClientVpnEndpoint#client_login_banner_options}
	ClientLoginBannerOptions *Ec2ClientVpnEndpointClientLoginBannerOptions `field:"optional" json:"clientLoginBannerOptions" yaml:"clientLoginBannerOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ec2_client_vpn_endpoint#description Ec2ClientVpnEndpoint#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ec2_client_vpn_endpoint#dns_servers Ec2ClientVpnEndpoint#dns_servers}.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ec2_client_vpn_endpoint#id Ec2ClientVpnEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ec2_client_vpn_endpoint#security_group_ids Ec2ClientVpnEndpoint#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ec2_client_vpn_endpoint#self_service_portal Ec2ClientVpnEndpoint#self_service_portal}.
	SelfServicePortal *string `field:"optional" json:"selfServicePortal" yaml:"selfServicePortal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ec2_client_vpn_endpoint#session_timeout_hours Ec2ClientVpnEndpoint#session_timeout_hours}.
	SessionTimeoutHours *float64 `field:"optional" json:"sessionTimeoutHours" yaml:"sessionTimeoutHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ec2_client_vpn_endpoint#split_tunnel Ec2ClientVpnEndpoint#split_tunnel}.
	SplitTunnel interface{} `field:"optional" json:"splitTunnel" yaml:"splitTunnel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ec2_client_vpn_endpoint#tags Ec2ClientVpnEndpoint#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ec2_client_vpn_endpoint#tags_all Ec2ClientVpnEndpoint#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ec2_client_vpn_endpoint#transport_protocol Ec2ClientVpnEndpoint#transport_protocol}.
	TransportProtocol *string `field:"optional" json:"transportProtocol" yaml:"transportProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ec2_client_vpn_endpoint#vpc_id Ec2ClientVpnEndpoint#vpc_id}.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/ec2_client_vpn_endpoint#vpn_port Ec2ClientVpnEndpoint#vpn_port}.
	VpnPort *float64 `field:"optional" json:"vpnPort" yaml:"vpnPort"`
}

