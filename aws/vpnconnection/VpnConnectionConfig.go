// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#customer_gateway_id VpnConnection#customer_gateway_id}.
	CustomerGatewayId *string `field:"required" json:"customerGatewayId" yaml:"customerGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#type VpnConnection#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#enable_acceleration VpnConnection#enable_acceleration}.
	EnableAcceleration interface{} `field:"optional" json:"enableAcceleration" yaml:"enableAcceleration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#id VpnConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#local_ipv4_network_cidr VpnConnection#local_ipv4_network_cidr}.
	LocalIpv4NetworkCidr *string `field:"optional" json:"localIpv4NetworkCidr" yaml:"localIpv4NetworkCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#local_ipv6_network_cidr VpnConnection#local_ipv6_network_cidr}.
	LocalIpv6NetworkCidr *string `field:"optional" json:"localIpv6NetworkCidr" yaml:"localIpv6NetworkCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#outside_ip_address_type VpnConnection#outside_ip_address_type}.
	OutsideIpAddressType *string `field:"optional" json:"outsideIpAddressType" yaml:"outsideIpAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#remote_ipv4_network_cidr VpnConnection#remote_ipv4_network_cidr}.
	RemoteIpv4NetworkCidr *string `field:"optional" json:"remoteIpv4NetworkCidr" yaml:"remoteIpv4NetworkCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#remote_ipv6_network_cidr VpnConnection#remote_ipv6_network_cidr}.
	RemoteIpv6NetworkCidr *string `field:"optional" json:"remoteIpv6NetworkCidr" yaml:"remoteIpv6NetworkCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#static_routes_only VpnConnection#static_routes_only}.
	StaticRoutesOnly interface{} `field:"optional" json:"staticRoutesOnly" yaml:"staticRoutesOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tags VpnConnection#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tags_all VpnConnection#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#transit_gateway_id VpnConnection#transit_gateway_id}.
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#transport_transit_gateway_attachment_id VpnConnection#transport_transit_gateway_attachment_id}.
	TransportTransitGatewayAttachmentId *string `field:"optional" json:"transportTransitGatewayAttachmentId" yaml:"transportTransitGatewayAttachmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel1_dpd_timeout_action VpnConnection#tunnel1_dpd_timeout_action}.
	Tunnel1DpdTimeoutAction *string `field:"optional" json:"tunnel1DpdTimeoutAction" yaml:"tunnel1DpdTimeoutAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel1_dpd_timeout_seconds VpnConnection#tunnel1_dpd_timeout_seconds}.
	Tunnel1DpdTimeoutSeconds *float64 `field:"optional" json:"tunnel1DpdTimeoutSeconds" yaml:"tunnel1DpdTimeoutSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel1_enable_tunnel_lifecycle_control VpnConnection#tunnel1_enable_tunnel_lifecycle_control}.
	Tunnel1EnableTunnelLifecycleControl interface{} `field:"optional" json:"tunnel1EnableTunnelLifecycleControl" yaml:"tunnel1EnableTunnelLifecycleControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel1_ike_versions VpnConnection#tunnel1_ike_versions}.
	Tunnel1IkeVersions *[]*string `field:"optional" json:"tunnel1IkeVersions" yaml:"tunnel1IkeVersions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel1_inside_cidr VpnConnection#tunnel1_inside_cidr}.
	Tunnel1InsideCidr *string `field:"optional" json:"tunnel1InsideCidr" yaml:"tunnel1InsideCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel1_inside_ipv6_cidr VpnConnection#tunnel1_inside_ipv6_cidr}.
	Tunnel1InsideIpv6Cidr *string `field:"optional" json:"tunnel1InsideIpv6Cidr" yaml:"tunnel1InsideIpv6Cidr"`
	// tunnel1_log_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel1_log_options VpnConnection#tunnel1_log_options}
	Tunnel1LogOptions *VpnConnectionTunnel1LogOptions `field:"optional" json:"tunnel1LogOptions" yaml:"tunnel1LogOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel1_phase1_dh_group_numbers VpnConnection#tunnel1_phase1_dh_group_numbers}.
	Tunnel1Phase1DhGroupNumbers *[]*float64 `field:"optional" json:"tunnel1Phase1DhGroupNumbers" yaml:"tunnel1Phase1DhGroupNumbers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel1_phase1_encryption_algorithms VpnConnection#tunnel1_phase1_encryption_algorithms}.
	Tunnel1Phase1EncryptionAlgorithms *[]*string `field:"optional" json:"tunnel1Phase1EncryptionAlgorithms" yaml:"tunnel1Phase1EncryptionAlgorithms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel1_phase1_integrity_algorithms VpnConnection#tunnel1_phase1_integrity_algorithms}.
	Tunnel1Phase1IntegrityAlgorithms *[]*string `field:"optional" json:"tunnel1Phase1IntegrityAlgorithms" yaml:"tunnel1Phase1IntegrityAlgorithms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel1_phase1_lifetime_seconds VpnConnection#tunnel1_phase1_lifetime_seconds}.
	Tunnel1Phase1LifetimeSeconds *float64 `field:"optional" json:"tunnel1Phase1LifetimeSeconds" yaml:"tunnel1Phase1LifetimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel1_phase2_dh_group_numbers VpnConnection#tunnel1_phase2_dh_group_numbers}.
	Tunnel1Phase2DhGroupNumbers *[]*float64 `field:"optional" json:"tunnel1Phase2DhGroupNumbers" yaml:"tunnel1Phase2DhGroupNumbers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel1_phase2_encryption_algorithms VpnConnection#tunnel1_phase2_encryption_algorithms}.
	Tunnel1Phase2EncryptionAlgorithms *[]*string `field:"optional" json:"tunnel1Phase2EncryptionAlgorithms" yaml:"tunnel1Phase2EncryptionAlgorithms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel1_phase2_integrity_algorithms VpnConnection#tunnel1_phase2_integrity_algorithms}.
	Tunnel1Phase2IntegrityAlgorithms *[]*string `field:"optional" json:"tunnel1Phase2IntegrityAlgorithms" yaml:"tunnel1Phase2IntegrityAlgorithms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel1_phase2_lifetime_seconds VpnConnection#tunnel1_phase2_lifetime_seconds}.
	Tunnel1Phase2LifetimeSeconds *float64 `field:"optional" json:"tunnel1Phase2LifetimeSeconds" yaml:"tunnel1Phase2LifetimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel1_preshared_key VpnConnection#tunnel1_preshared_key}.
	Tunnel1PresharedKey *string `field:"optional" json:"tunnel1PresharedKey" yaml:"tunnel1PresharedKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel1_rekey_fuzz_percentage VpnConnection#tunnel1_rekey_fuzz_percentage}.
	Tunnel1RekeyFuzzPercentage *float64 `field:"optional" json:"tunnel1RekeyFuzzPercentage" yaml:"tunnel1RekeyFuzzPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel1_rekey_margin_time_seconds VpnConnection#tunnel1_rekey_margin_time_seconds}.
	Tunnel1RekeyMarginTimeSeconds *float64 `field:"optional" json:"tunnel1RekeyMarginTimeSeconds" yaml:"tunnel1RekeyMarginTimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel1_replay_window_size VpnConnection#tunnel1_replay_window_size}.
	Tunnel1ReplayWindowSize *float64 `field:"optional" json:"tunnel1ReplayWindowSize" yaml:"tunnel1ReplayWindowSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel1_startup_action VpnConnection#tunnel1_startup_action}.
	Tunnel1StartupAction *string `field:"optional" json:"tunnel1StartupAction" yaml:"tunnel1StartupAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel2_dpd_timeout_action VpnConnection#tunnel2_dpd_timeout_action}.
	Tunnel2DpdTimeoutAction *string `field:"optional" json:"tunnel2DpdTimeoutAction" yaml:"tunnel2DpdTimeoutAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel2_dpd_timeout_seconds VpnConnection#tunnel2_dpd_timeout_seconds}.
	Tunnel2DpdTimeoutSeconds *float64 `field:"optional" json:"tunnel2DpdTimeoutSeconds" yaml:"tunnel2DpdTimeoutSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel2_enable_tunnel_lifecycle_control VpnConnection#tunnel2_enable_tunnel_lifecycle_control}.
	Tunnel2EnableTunnelLifecycleControl interface{} `field:"optional" json:"tunnel2EnableTunnelLifecycleControl" yaml:"tunnel2EnableTunnelLifecycleControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel2_ike_versions VpnConnection#tunnel2_ike_versions}.
	Tunnel2IkeVersions *[]*string `field:"optional" json:"tunnel2IkeVersions" yaml:"tunnel2IkeVersions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel2_inside_cidr VpnConnection#tunnel2_inside_cidr}.
	Tunnel2InsideCidr *string `field:"optional" json:"tunnel2InsideCidr" yaml:"tunnel2InsideCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel2_inside_ipv6_cidr VpnConnection#tunnel2_inside_ipv6_cidr}.
	Tunnel2InsideIpv6Cidr *string `field:"optional" json:"tunnel2InsideIpv6Cidr" yaml:"tunnel2InsideIpv6Cidr"`
	// tunnel2_log_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel2_log_options VpnConnection#tunnel2_log_options}
	Tunnel2LogOptions *VpnConnectionTunnel2LogOptions `field:"optional" json:"tunnel2LogOptions" yaml:"tunnel2LogOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel2_phase1_dh_group_numbers VpnConnection#tunnel2_phase1_dh_group_numbers}.
	Tunnel2Phase1DhGroupNumbers *[]*float64 `field:"optional" json:"tunnel2Phase1DhGroupNumbers" yaml:"tunnel2Phase1DhGroupNumbers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel2_phase1_encryption_algorithms VpnConnection#tunnel2_phase1_encryption_algorithms}.
	Tunnel2Phase1EncryptionAlgorithms *[]*string `field:"optional" json:"tunnel2Phase1EncryptionAlgorithms" yaml:"tunnel2Phase1EncryptionAlgorithms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel2_phase1_integrity_algorithms VpnConnection#tunnel2_phase1_integrity_algorithms}.
	Tunnel2Phase1IntegrityAlgorithms *[]*string `field:"optional" json:"tunnel2Phase1IntegrityAlgorithms" yaml:"tunnel2Phase1IntegrityAlgorithms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel2_phase1_lifetime_seconds VpnConnection#tunnel2_phase1_lifetime_seconds}.
	Tunnel2Phase1LifetimeSeconds *float64 `field:"optional" json:"tunnel2Phase1LifetimeSeconds" yaml:"tunnel2Phase1LifetimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel2_phase2_dh_group_numbers VpnConnection#tunnel2_phase2_dh_group_numbers}.
	Tunnel2Phase2DhGroupNumbers *[]*float64 `field:"optional" json:"tunnel2Phase2DhGroupNumbers" yaml:"tunnel2Phase2DhGroupNumbers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel2_phase2_encryption_algorithms VpnConnection#tunnel2_phase2_encryption_algorithms}.
	Tunnel2Phase2EncryptionAlgorithms *[]*string `field:"optional" json:"tunnel2Phase2EncryptionAlgorithms" yaml:"tunnel2Phase2EncryptionAlgorithms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel2_phase2_integrity_algorithms VpnConnection#tunnel2_phase2_integrity_algorithms}.
	Tunnel2Phase2IntegrityAlgorithms *[]*string `field:"optional" json:"tunnel2Phase2IntegrityAlgorithms" yaml:"tunnel2Phase2IntegrityAlgorithms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel2_phase2_lifetime_seconds VpnConnection#tunnel2_phase2_lifetime_seconds}.
	Tunnel2Phase2LifetimeSeconds *float64 `field:"optional" json:"tunnel2Phase2LifetimeSeconds" yaml:"tunnel2Phase2LifetimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel2_preshared_key VpnConnection#tunnel2_preshared_key}.
	Tunnel2PresharedKey *string `field:"optional" json:"tunnel2PresharedKey" yaml:"tunnel2PresharedKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel2_rekey_fuzz_percentage VpnConnection#tunnel2_rekey_fuzz_percentage}.
	Tunnel2RekeyFuzzPercentage *float64 `field:"optional" json:"tunnel2RekeyFuzzPercentage" yaml:"tunnel2RekeyFuzzPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel2_rekey_margin_time_seconds VpnConnection#tunnel2_rekey_margin_time_seconds}.
	Tunnel2RekeyMarginTimeSeconds *float64 `field:"optional" json:"tunnel2RekeyMarginTimeSeconds" yaml:"tunnel2RekeyMarginTimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel2_replay_window_size VpnConnection#tunnel2_replay_window_size}.
	Tunnel2ReplayWindowSize *float64 `field:"optional" json:"tunnel2ReplayWindowSize" yaml:"tunnel2ReplayWindowSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel2_startup_action VpnConnection#tunnel2_startup_action}.
	Tunnel2StartupAction *string `field:"optional" json:"tunnel2StartupAction" yaml:"tunnel2StartupAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#tunnel_inside_ip_version VpnConnection#tunnel_inside_ip_version}.
	TunnelInsideIpVersion *string `field:"optional" json:"tunnelInsideIpVersion" yaml:"tunnelInsideIpVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/vpn_connection#vpn_gateway_id VpnConnection#vpn_gateway_id}.
	VpnGatewayId *string `field:"optional" json:"vpnGatewayId" yaml:"vpnGatewayId"`
}

