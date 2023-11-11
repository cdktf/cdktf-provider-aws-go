// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/vpnconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/vpn_connection aws_vpn_connection}.
type VpnConnection interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CoreNetworkArn() *string
	CoreNetworkAttachmentArn() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomerGatewayConfiguration() *string
	CustomerGatewayId() *string
	SetCustomerGatewayId(val *string)
	CustomerGatewayIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableAcceleration() interface{}
	SetEnableAcceleration(val interface{})
	EnableAccelerationInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalIpv4NetworkCidr() *string
	SetLocalIpv4NetworkCidr(val *string)
	LocalIpv4NetworkCidrInput() *string
	LocalIpv6NetworkCidr() *string
	SetLocalIpv6NetworkCidr(val *string)
	LocalIpv6NetworkCidrInput() *string
	// The tree node.
	Node() constructs.Node
	OutsideIpAddressType() *string
	SetOutsideIpAddressType(val *string)
	OutsideIpAddressTypeInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RemoteIpv4NetworkCidr() *string
	SetRemoteIpv4NetworkCidr(val *string)
	RemoteIpv4NetworkCidrInput() *string
	RemoteIpv6NetworkCidr() *string
	SetRemoteIpv6NetworkCidr(val *string)
	RemoteIpv6NetworkCidrInput() *string
	Routes() VpnConnectionRoutesList
	StaticRoutesOnly() interface{}
	SetStaticRoutesOnly(val interface{})
	StaticRoutesOnlyInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TransitGatewayAttachmentId() *string
	TransitGatewayId() *string
	SetTransitGatewayId(val *string)
	TransitGatewayIdInput() *string
	TransportTransitGatewayAttachmentId() *string
	SetTransportTransitGatewayAttachmentId(val *string)
	TransportTransitGatewayAttachmentIdInput() *string
	Tunnel1Address() *string
	Tunnel1BgpAsn() *string
	Tunnel1BgpHoldtime() *float64
	Tunnel1CgwInsideAddress() *string
	Tunnel1DpdTimeoutAction() *string
	SetTunnel1DpdTimeoutAction(val *string)
	Tunnel1DpdTimeoutActionInput() *string
	Tunnel1DpdTimeoutSeconds() *float64
	SetTunnel1DpdTimeoutSeconds(val *float64)
	Tunnel1DpdTimeoutSecondsInput() *float64
	Tunnel1EnableTunnelLifecycleControl() interface{}
	SetTunnel1EnableTunnelLifecycleControl(val interface{})
	Tunnel1EnableTunnelLifecycleControlInput() interface{}
	Tunnel1IkeVersions() *[]*string
	SetTunnel1IkeVersions(val *[]*string)
	Tunnel1IkeVersionsInput() *[]*string
	Tunnel1InsideCidr() *string
	SetTunnel1InsideCidr(val *string)
	Tunnel1InsideCidrInput() *string
	Tunnel1InsideIpv6Cidr() *string
	SetTunnel1InsideIpv6Cidr(val *string)
	Tunnel1InsideIpv6CidrInput() *string
	Tunnel1LogOptions() VpnConnectionTunnel1LogOptionsOutputReference
	Tunnel1LogOptionsInput() *VpnConnectionTunnel1LogOptions
	Tunnel1Phase1DhGroupNumbers() *[]*float64
	SetTunnel1Phase1DhGroupNumbers(val *[]*float64)
	Tunnel1Phase1DhGroupNumbersInput() *[]*float64
	Tunnel1Phase1EncryptionAlgorithms() *[]*string
	SetTunnel1Phase1EncryptionAlgorithms(val *[]*string)
	Tunnel1Phase1EncryptionAlgorithmsInput() *[]*string
	Tunnel1Phase1IntegrityAlgorithms() *[]*string
	SetTunnel1Phase1IntegrityAlgorithms(val *[]*string)
	Tunnel1Phase1IntegrityAlgorithmsInput() *[]*string
	Tunnel1Phase1LifetimeSeconds() *float64
	SetTunnel1Phase1LifetimeSeconds(val *float64)
	Tunnel1Phase1LifetimeSecondsInput() *float64
	Tunnel1Phase2DhGroupNumbers() *[]*float64
	SetTunnel1Phase2DhGroupNumbers(val *[]*float64)
	Tunnel1Phase2DhGroupNumbersInput() *[]*float64
	Tunnel1Phase2EncryptionAlgorithms() *[]*string
	SetTunnel1Phase2EncryptionAlgorithms(val *[]*string)
	Tunnel1Phase2EncryptionAlgorithmsInput() *[]*string
	Tunnel1Phase2IntegrityAlgorithms() *[]*string
	SetTunnel1Phase2IntegrityAlgorithms(val *[]*string)
	Tunnel1Phase2IntegrityAlgorithmsInput() *[]*string
	Tunnel1Phase2LifetimeSeconds() *float64
	SetTunnel1Phase2LifetimeSeconds(val *float64)
	Tunnel1Phase2LifetimeSecondsInput() *float64
	Tunnel1PresharedKey() *string
	SetTunnel1PresharedKey(val *string)
	Tunnel1PresharedKeyInput() *string
	Tunnel1RekeyFuzzPercentage() *float64
	SetTunnel1RekeyFuzzPercentage(val *float64)
	Tunnel1RekeyFuzzPercentageInput() *float64
	Tunnel1RekeyMarginTimeSeconds() *float64
	SetTunnel1RekeyMarginTimeSeconds(val *float64)
	Tunnel1RekeyMarginTimeSecondsInput() *float64
	Tunnel1ReplayWindowSize() *float64
	SetTunnel1ReplayWindowSize(val *float64)
	Tunnel1ReplayWindowSizeInput() *float64
	Tunnel1StartupAction() *string
	SetTunnel1StartupAction(val *string)
	Tunnel1StartupActionInput() *string
	Tunnel1VgwInsideAddress() *string
	Tunnel2Address() *string
	Tunnel2BgpAsn() *string
	Tunnel2BgpHoldtime() *float64
	Tunnel2CgwInsideAddress() *string
	Tunnel2DpdTimeoutAction() *string
	SetTunnel2DpdTimeoutAction(val *string)
	Tunnel2DpdTimeoutActionInput() *string
	Tunnel2DpdTimeoutSeconds() *float64
	SetTunnel2DpdTimeoutSeconds(val *float64)
	Tunnel2DpdTimeoutSecondsInput() *float64
	Tunnel2EnableTunnelLifecycleControl() interface{}
	SetTunnel2EnableTunnelLifecycleControl(val interface{})
	Tunnel2EnableTunnelLifecycleControlInput() interface{}
	Tunnel2IkeVersions() *[]*string
	SetTunnel2IkeVersions(val *[]*string)
	Tunnel2IkeVersionsInput() *[]*string
	Tunnel2InsideCidr() *string
	SetTunnel2InsideCidr(val *string)
	Tunnel2InsideCidrInput() *string
	Tunnel2InsideIpv6Cidr() *string
	SetTunnel2InsideIpv6Cidr(val *string)
	Tunnel2InsideIpv6CidrInput() *string
	Tunnel2LogOptions() VpnConnectionTunnel2LogOptionsOutputReference
	Tunnel2LogOptionsInput() *VpnConnectionTunnel2LogOptions
	Tunnel2Phase1DhGroupNumbers() *[]*float64
	SetTunnel2Phase1DhGroupNumbers(val *[]*float64)
	Tunnel2Phase1DhGroupNumbersInput() *[]*float64
	Tunnel2Phase1EncryptionAlgorithms() *[]*string
	SetTunnel2Phase1EncryptionAlgorithms(val *[]*string)
	Tunnel2Phase1EncryptionAlgorithmsInput() *[]*string
	Tunnel2Phase1IntegrityAlgorithms() *[]*string
	SetTunnel2Phase1IntegrityAlgorithms(val *[]*string)
	Tunnel2Phase1IntegrityAlgorithmsInput() *[]*string
	Tunnel2Phase1LifetimeSeconds() *float64
	SetTunnel2Phase1LifetimeSeconds(val *float64)
	Tunnel2Phase1LifetimeSecondsInput() *float64
	Tunnel2Phase2DhGroupNumbers() *[]*float64
	SetTunnel2Phase2DhGroupNumbers(val *[]*float64)
	Tunnel2Phase2DhGroupNumbersInput() *[]*float64
	Tunnel2Phase2EncryptionAlgorithms() *[]*string
	SetTunnel2Phase2EncryptionAlgorithms(val *[]*string)
	Tunnel2Phase2EncryptionAlgorithmsInput() *[]*string
	Tunnel2Phase2IntegrityAlgorithms() *[]*string
	SetTunnel2Phase2IntegrityAlgorithms(val *[]*string)
	Tunnel2Phase2IntegrityAlgorithmsInput() *[]*string
	Tunnel2Phase2LifetimeSeconds() *float64
	SetTunnel2Phase2LifetimeSeconds(val *float64)
	Tunnel2Phase2LifetimeSecondsInput() *float64
	Tunnel2PresharedKey() *string
	SetTunnel2PresharedKey(val *string)
	Tunnel2PresharedKeyInput() *string
	Tunnel2RekeyFuzzPercentage() *float64
	SetTunnel2RekeyFuzzPercentage(val *float64)
	Tunnel2RekeyFuzzPercentageInput() *float64
	Tunnel2RekeyMarginTimeSeconds() *float64
	SetTunnel2RekeyMarginTimeSeconds(val *float64)
	Tunnel2RekeyMarginTimeSecondsInput() *float64
	Tunnel2ReplayWindowSize() *float64
	SetTunnel2ReplayWindowSize(val *float64)
	Tunnel2ReplayWindowSizeInput() *float64
	Tunnel2StartupAction() *string
	SetTunnel2StartupAction(val *string)
	Tunnel2StartupActionInput() *string
	Tunnel2VgwInsideAddress() *string
	TunnelInsideIpVersion() *string
	SetTunnelInsideIpVersion(val *string)
	TunnelInsideIpVersionInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	VgwTelemetry() VpnConnectionVgwTelemetryList
	VpnGatewayId() *string
	SetVpnGatewayId(val *string)
	VpnGatewayIdInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTunnel1LogOptions(value *VpnConnectionTunnel1LogOptions)
	PutTunnel2LogOptions(value *VpnConnectionTunnel2LogOptions)
	ResetEnableAcceleration()
	ResetId()
	ResetLocalIpv4NetworkCidr()
	ResetLocalIpv6NetworkCidr()
	ResetOutsideIpAddressType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRemoteIpv4NetworkCidr()
	ResetRemoteIpv6NetworkCidr()
	ResetStaticRoutesOnly()
	ResetTags()
	ResetTagsAll()
	ResetTransitGatewayId()
	ResetTransportTransitGatewayAttachmentId()
	ResetTunnel1DpdTimeoutAction()
	ResetTunnel1DpdTimeoutSeconds()
	ResetTunnel1EnableTunnelLifecycleControl()
	ResetTunnel1IkeVersions()
	ResetTunnel1InsideCidr()
	ResetTunnel1InsideIpv6Cidr()
	ResetTunnel1LogOptions()
	ResetTunnel1Phase1DhGroupNumbers()
	ResetTunnel1Phase1EncryptionAlgorithms()
	ResetTunnel1Phase1IntegrityAlgorithms()
	ResetTunnel1Phase1LifetimeSeconds()
	ResetTunnel1Phase2DhGroupNumbers()
	ResetTunnel1Phase2EncryptionAlgorithms()
	ResetTunnel1Phase2IntegrityAlgorithms()
	ResetTunnel1Phase2LifetimeSeconds()
	ResetTunnel1PresharedKey()
	ResetTunnel1RekeyFuzzPercentage()
	ResetTunnel1RekeyMarginTimeSeconds()
	ResetTunnel1ReplayWindowSize()
	ResetTunnel1StartupAction()
	ResetTunnel2DpdTimeoutAction()
	ResetTunnel2DpdTimeoutSeconds()
	ResetTunnel2EnableTunnelLifecycleControl()
	ResetTunnel2IkeVersions()
	ResetTunnel2InsideCidr()
	ResetTunnel2InsideIpv6Cidr()
	ResetTunnel2LogOptions()
	ResetTunnel2Phase1DhGroupNumbers()
	ResetTunnel2Phase1EncryptionAlgorithms()
	ResetTunnel2Phase1IntegrityAlgorithms()
	ResetTunnel2Phase1LifetimeSeconds()
	ResetTunnel2Phase2DhGroupNumbers()
	ResetTunnel2Phase2EncryptionAlgorithms()
	ResetTunnel2Phase2IntegrityAlgorithms()
	ResetTunnel2Phase2LifetimeSeconds()
	ResetTunnel2PresharedKey()
	ResetTunnel2RekeyFuzzPercentage()
	ResetTunnel2RekeyMarginTimeSeconds()
	ResetTunnel2ReplayWindowSize()
	ResetTunnel2StartupAction()
	ResetTunnelInsideIpVersion()
	ResetVpnGatewayId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VpnConnection
type jsiiProxy_VpnConnection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VpnConnection) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) CoreNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) CoreNetworkAttachmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkAttachmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) CustomerGatewayConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerGatewayConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) CustomerGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) CustomerGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) EnableAcceleration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceleration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) EnableAccelerationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAccelerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) LocalIpv4NetworkCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpv4NetworkCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) LocalIpv4NetworkCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpv4NetworkCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) LocalIpv6NetworkCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpv6NetworkCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) LocalIpv6NetworkCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpv6NetworkCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) OutsideIpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outsideIpAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) OutsideIpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outsideIpAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) RemoteIpv4NetworkCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpv4NetworkCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) RemoteIpv4NetworkCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpv4NetworkCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) RemoteIpv6NetworkCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpv6NetworkCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) RemoteIpv6NetworkCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpv6NetworkCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Routes() VpnConnectionRoutesList {
	var returns VpnConnectionRoutesList
	_jsii_.Get(
		j,
		"routes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) StaticRoutesOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticRoutesOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) StaticRoutesOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticRoutesOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) TransitGatewayAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) TransitGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) TransitGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) TransportTransitGatewayAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transportTransitGatewayAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) TransportTransitGatewayAttachmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transportTransitGatewayAttachmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1BgpAsn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1BgpAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1BgpHoldtime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1BgpHoldtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1CgwInsideAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1CgwInsideAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1DpdTimeoutAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1DpdTimeoutAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1DpdTimeoutActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1DpdTimeoutActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1DpdTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1DpdTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1DpdTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1DpdTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1EnableTunnelLifecycleControl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tunnel1EnableTunnelLifecycleControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1EnableTunnelLifecycleControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tunnel1EnableTunnelLifecycleControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1IkeVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1IkeVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1IkeVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1IkeVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1InsideCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1InsideCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1InsideCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1InsideCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1InsideIpv6Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1InsideIpv6Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1InsideIpv6CidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1InsideIpv6CidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1LogOptions() VpnConnectionTunnel1LogOptionsOutputReference {
	var returns VpnConnectionTunnel1LogOptionsOutputReference
	_jsii_.Get(
		j,
		"tunnel1LogOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1LogOptionsInput() *VpnConnectionTunnel1LogOptions {
	var returns *VpnConnectionTunnel1LogOptions
	_jsii_.Get(
		j,
		"tunnel1LogOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1Phase1DhGroupNumbers() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel1Phase1DhGroupNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1Phase1DhGroupNumbersInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel1Phase1DhGroupNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1Phase1EncryptionAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase1EncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1Phase1EncryptionAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase1EncryptionAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1Phase1IntegrityAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase1IntegrityAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1Phase1IntegrityAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase1IntegrityAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1Phase1LifetimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1Phase1LifetimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1Phase1LifetimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1Phase1LifetimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1Phase2DhGroupNumbers() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel1Phase2DhGroupNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1Phase2DhGroupNumbersInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel1Phase2DhGroupNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1Phase2EncryptionAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase2EncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1Phase2EncryptionAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase2EncryptionAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1Phase2IntegrityAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase2IntegrityAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1Phase2IntegrityAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase2IntegrityAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1Phase2LifetimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1Phase2LifetimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1Phase2LifetimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1Phase2LifetimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1PresharedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1PresharedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1PresharedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1PresharedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1RekeyFuzzPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1RekeyFuzzPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1RekeyFuzzPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1RekeyFuzzPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1RekeyMarginTimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1RekeyMarginTimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1RekeyMarginTimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1RekeyMarginTimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1ReplayWindowSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1ReplayWindowSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1ReplayWindowSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1ReplayWindowSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1StartupAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1StartupAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1StartupActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1StartupActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel1VgwInsideAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1VgwInsideAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2BgpAsn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2BgpAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2BgpHoldtime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2BgpHoldtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2CgwInsideAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2CgwInsideAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2DpdTimeoutAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2DpdTimeoutAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2DpdTimeoutActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2DpdTimeoutActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2DpdTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2DpdTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2DpdTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2DpdTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2EnableTunnelLifecycleControl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tunnel2EnableTunnelLifecycleControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2EnableTunnelLifecycleControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tunnel2EnableTunnelLifecycleControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2IkeVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2IkeVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2IkeVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2IkeVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2InsideCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2InsideCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2InsideCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2InsideCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2InsideIpv6Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2InsideIpv6Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2InsideIpv6CidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2InsideIpv6CidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2LogOptions() VpnConnectionTunnel2LogOptionsOutputReference {
	var returns VpnConnectionTunnel2LogOptionsOutputReference
	_jsii_.Get(
		j,
		"tunnel2LogOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2LogOptionsInput() *VpnConnectionTunnel2LogOptions {
	var returns *VpnConnectionTunnel2LogOptions
	_jsii_.Get(
		j,
		"tunnel2LogOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2Phase1DhGroupNumbers() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel2Phase1DhGroupNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2Phase1DhGroupNumbersInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel2Phase1DhGroupNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2Phase1EncryptionAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase1EncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2Phase1EncryptionAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase1EncryptionAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2Phase1IntegrityAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase1IntegrityAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2Phase1IntegrityAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase1IntegrityAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2Phase1LifetimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2Phase1LifetimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2Phase1LifetimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2Phase1LifetimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2Phase2DhGroupNumbers() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel2Phase2DhGroupNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2Phase2DhGroupNumbersInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel2Phase2DhGroupNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2Phase2EncryptionAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase2EncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2Phase2EncryptionAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase2EncryptionAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2Phase2IntegrityAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase2IntegrityAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2Phase2IntegrityAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase2IntegrityAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2Phase2LifetimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2Phase2LifetimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2Phase2LifetimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2Phase2LifetimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2PresharedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2PresharedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2PresharedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2PresharedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2RekeyFuzzPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2RekeyFuzzPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2RekeyFuzzPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2RekeyFuzzPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2RekeyMarginTimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2RekeyMarginTimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2RekeyMarginTimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2RekeyMarginTimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2ReplayWindowSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2ReplayWindowSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2ReplayWindowSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2ReplayWindowSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2StartupAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2StartupAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2StartupActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2StartupActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Tunnel2VgwInsideAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2VgwInsideAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) TunnelInsideIpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelInsideIpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) TunnelInsideIpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelInsideIpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) VgwTelemetry() VpnConnectionVgwTelemetryList {
	var returns VpnConnectionVgwTelemetryList
	_jsii_.Get(
		j,
		"vgwTelemetry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) VpnGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnConnection) VpnGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnGatewayIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/vpn_connection aws_vpn_connection} Resource.
func NewVpnConnection(scope constructs.Construct, id *string, config *VpnConnectionConfig) VpnConnection {
	_init_.Initialize()

	if err := validateNewVpnConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpnConnection{}

	_jsii_.Create(
		"@cdktf/provider-aws.vpnConnection.VpnConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/vpn_connection aws_vpn_connection} Resource.
func NewVpnConnection_Override(v VpnConnection, scope constructs.Construct, id *string, config *VpnConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.vpnConnection.VpnConnection",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VpnConnection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetCustomerGatewayId(val *string) {
	if err := j.validateSetCustomerGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerGatewayId",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetEnableAcceleration(val interface{}) {
	if err := j.validateSetEnableAccelerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAcceleration",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetLocalIpv4NetworkCidr(val *string) {
	if err := j.validateSetLocalIpv4NetworkCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localIpv4NetworkCidr",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetLocalIpv6NetworkCidr(val *string) {
	if err := j.validateSetLocalIpv6NetworkCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localIpv6NetworkCidr",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetOutsideIpAddressType(val *string) {
	if err := j.validateSetOutsideIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outsideIpAddressType",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetRemoteIpv4NetworkCidr(val *string) {
	if err := j.validateSetRemoteIpv4NetworkCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteIpv4NetworkCidr",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetRemoteIpv6NetworkCidr(val *string) {
	if err := j.validateSetRemoteIpv6NetworkCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteIpv6NetworkCidr",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetStaticRoutesOnly(val interface{}) {
	if err := j.validateSetStaticRoutesOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staticRoutesOnly",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTransitGatewayId(val *string) {
	if err := j.validateSetTransitGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitGatewayId",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTransportTransitGatewayAttachmentId(val *string) {
	if err := j.validateSetTransportTransitGatewayAttachmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transportTransitGatewayAttachmentId",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel1DpdTimeoutAction(val *string) {
	if err := j.validateSetTunnel1DpdTimeoutActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1DpdTimeoutAction",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel1DpdTimeoutSeconds(val *float64) {
	if err := j.validateSetTunnel1DpdTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1DpdTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel1EnableTunnelLifecycleControl(val interface{}) {
	if err := j.validateSetTunnel1EnableTunnelLifecycleControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1EnableTunnelLifecycleControl",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel1IkeVersions(val *[]*string) {
	if err := j.validateSetTunnel1IkeVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1IkeVersions",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel1InsideCidr(val *string) {
	if err := j.validateSetTunnel1InsideCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1InsideCidr",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel1InsideIpv6Cidr(val *string) {
	if err := j.validateSetTunnel1InsideIpv6CidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1InsideIpv6Cidr",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel1Phase1DhGroupNumbers(val *[]*float64) {
	if err := j.validateSetTunnel1Phase1DhGroupNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase1DhGroupNumbers",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel1Phase1EncryptionAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel1Phase1EncryptionAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase1EncryptionAlgorithms",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel1Phase1IntegrityAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel1Phase1IntegrityAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase1IntegrityAlgorithms",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel1Phase1LifetimeSeconds(val *float64) {
	if err := j.validateSetTunnel1Phase1LifetimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase1LifetimeSeconds",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel1Phase2DhGroupNumbers(val *[]*float64) {
	if err := j.validateSetTunnel1Phase2DhGroupNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase2DhGroupNumbers",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel1Phase2EncryptionAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel1Phase2EncryptionAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase2EncryptionAlgorithms",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel1Phase2IntegrityAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel1Phase2IntegrityAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase2IntegrityAlgorithms",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel1Phase2LifetimeSeconds(val *float64) {
	if err := j.validateSetTunnel1Phase2LifetimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase2LifetimeSeconds",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel1PresharedKey(val *string) {
	if err := j.validateSetTunnel1PresharedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1PresharedKey",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel1RekeyFuzzPercentage(val *float64) {
	if err := j.validateSetTunnel1RekeyFuzzPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1RekeyFuzzPercentage",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel1RekeyMarginTimeSeconds(val *float64) {
	if err := j.validateSetTunnel1RekeyMarginTimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1RekeyMarginTimeSeconds",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel1ReplayWindowSize(val *float64) {
	if err := j.validateSetTunnel1ReplayWindowSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1ReplayWindowSize",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel1StartupAction(val *string) {
	if err := j.validateSetTunnel1StartupActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1StartupAction",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel2DpdTimeoutAction(val *string) {
	if err := j.validateSetTunnel2DpdTimeoutActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2DpdTimeoutAction",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel2DpdTimeoutSeconds(val *float64) {
	if err := j.validateSetTunnel2DpdTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2DpdTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel2EnableTunnelLifecycleControl(val interface{}) {
	if err := j.validateSetTunnel2EnableTunnelLifecycleControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2EnableTunnelLifecycleControl",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel2IkeVersions(val *[]*string) {
	if err := j.validateSetTunnel2IkeVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2IkeVersions",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel2InsideCidr(val *string) {
	if err := j.validateSetTunnel2InsideCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2InsideCidr",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel2InsideIpv6Cidr(val *string) {
	if err := j.validateSetTunnel2InsideIpv6CidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2InsideIpv6Cidr",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel2Phase1DhGroupNumbers(val *[]*float64) {
	if err := j.validateSetTunnel2Phase1DhGroupNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase1DhGroupNumbers",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel2Phase1EncryptionAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel2Phase1EncryptionAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase1EncryptionAlgorithms",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel2Phase1IntegrityAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel2Phase1IntegrityAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase1IntegrityAlgorithms",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel2Phase1LifetimeSeconds(val *float64) {
	if err := j.validateSetTunnel2Phase1LifetimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase1LifetimeSeconds",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel2Phase2DhGroupNumbers(val *[]*float64) {
	if err := j.validateSetTunnel2Phase2DhGroupNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase2DhGroupNumbers",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel2Phase2EncryptionAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel2Phase2EncryptionAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase2EncryptionAlgorithms",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel2Phase2IntegrityAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel2Phase2IntegrityAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase2IntegrityAlgorithms",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel2Phase2LifetimeSeconds(val *float64) {
	if err := j.validateSetTunnel2Phase2LifetimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase2LifetimeSeconds",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel2PresharedKey(val *string) {
	if err := j.validateSetTunnel2PresharedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2PresharedKey",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel2RekeyFuzzPercentage(val *float64) {
	if err := j.validateSetTunnel2RekeyFuzzPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2RekeyFuzzPercentage",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel2RekeyMarginTimeSeconds(val *float64) {
	if err := j.validateSetTunnel2RekeyMarginTimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2RekeyMarginTimeSeconds",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel2ReplayWindowSize(val *float64) {
	if err := j.validateSetTunnel2ReplayWindowSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2ReplayWindowSize",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnel2StartupAction(val *string) {
	if err := j.validateSetTunnel2StartupActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2StartupAction",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetTunnelInsideIpVersion(val *string) {
	if err := j.validateSetTunnelInsideIpVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelInsideIpVersion",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_VpnConnection)SetVpnGatewayId(val *string) {
	if err := j.validateSetVpnGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnGatewayId",
		val,
	)
}

// Generates CDKTF code for importing a VpnConnection resource upon running "cdktf plan <stack-name>".
func VpnConnection_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVpnConnection_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpnConnection.VpnConnection",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func VpnConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpnConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpnConnection.VpnConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpnConnection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpnConnection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpnConnection.VpnConnection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpnConnection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpnConnection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpnConnection.VpnConnection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VpnConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.vpnConnection.VpnConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VpnConnection) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VpnConnection) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VpnConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnection) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnection) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnection) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VpnConnection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnection) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VpnConnection) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VpnConnection) PutTunnel1LogOptions(value *VpnConnectionTunnel1LogOptions) {
	if err := v.validatePutTunnel1LogOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTunnel1LogOptions",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnConnection) PutTunnel2LogOptions(value *VpnConnectionTunnel2LogOptions) {
	if err := v.validatePutTunnel2LogOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTunnel2LogOptions",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnConnection) ResetEnableAcceleration() {
	_jsii_.InvokeVoid(
		v,
		"resetEnableAcceleration",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetLocalIpv4NetworkCidr() {
	_jsii_.InvokeVoid(
		v,
		"resetLocalIpv4NetworkCidr",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetLocalIpv6NetworkCidr() {
	_jsii_.InvokeVoid(
		v,
		"resetLocalIpv6NetworkCidr",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetOutsideIpAddressType() {
	_jsii_.InvokeVoid(
		v,
		"resetOutsideIpAddressType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetRemoteIpv4NetworkCidr() {
	_jsii_.InvokeVoid(
		v,
		"resetRemoteIpv4NetworkCidr",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetRemoteIpv6NetworkCidr() {
	_jsii_.InvokeVoid(
		v,
		"resetRemoteIpv6NetworkCidr",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetStaticRoutesOnly() {
	_jsii_.InvokeVoid(
		v,
		"resetStaticRoutesOnly",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTagsAll() {
	_jsii_.InvokeVoid(
		v,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTransitGatewayId() {
	_jsii_.InvokeVoid(
		v,
		"resetTransitGatewayId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTransportTransitGatewayAttachmentId() {
	_jsii_.InvokeVoid(
		v,
		"resetTransportTransitGatewayAttachmentId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel1DpdTimeoutAction() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel1DpdTimeoutAction",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel1DpdTimeoutSeconds() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel1DpdTimeoutSeconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel1EnableTunnelLifecycleControl() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel1EnableTunnelLifecycleControl",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel1IkeVersions() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel1IkeVersions",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel1InsideCidr() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel1InsideCidr",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel1InsideIpv6Cidr() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel1InsideIpv6Cidr",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel1LogOptions() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel1LogOptions",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel1Phase1DhGroupNumbers() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel1Phase1DhGroupNumbers",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel1Phase1EncryptionAlgorithms() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel1Phase1EncryptionAlgorithms",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel1Phase1IntegrityAlgorithms() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel1Phase1IntegrityAlgorithms",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel1Phase1LifetimeSeconds() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel1Phase1LifetimeSeconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel1Phase2DhGroupNumbers() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel1Phase2DhGroupNumbers",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel1Phase2EncryptionAlgorithms() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel1Phase2EncryptionAlgorithms",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel1Phase2IntegrityAlgorithms() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel1Phase2IntegrityAlgorithms",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel1Phase2LifetimeSeconds() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel1Phase2LifetimeSeconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel1PresharedKey() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel1PresharedKey",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel1RekeyFuzzPercentage() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel1RekeyFuzzPercentage",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel1RekeyMarginTimeSeconds() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel1RekeyMarginTimeSeconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel1ReplayWindowSize() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel1ReplayWindowSize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel1StartupAction() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel1StartupAction",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel2DpdTimeoutAction() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel2DpdTimeoutAction",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel2DpdTimeoutSeconds() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel2DpdTimeoutSeconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel2EnableTunnelLifecycleControl() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel2EnableTunnelLifecycleControl",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel2IkeVersions() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel2IkeVersions",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel2InsideCidr() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel2InsideCidr",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel2InsideIpv6Cidr() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel2InsideIpv6Cidr",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel2LogOptions() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel2LogOptions",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel2Phase1DhGroupNumbers() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel2Phase1DhGroupNumbers",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel2Phase1EncryptionAlgorithms() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel2Phase1EncryptionAlgorithms",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel2Phase1IntegrityAlgorithms() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel2Phase1IntegrityAlgorithms",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel2Phase1LifetimeSeconds() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel2Phase1LifetimeSeconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel2Phase2DhGroupNumbers() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel2Phase2DhGroupNumbers",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel2Phase2EncryptionAlgorithms() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel2Phase2EncryptionAlgorithms",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel2Phase2IntegrityAlgorithms() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel2Phase2IntegrityAlgorithms",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel2Phase2LifetimeSeconds() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel2Phase2LifetimeSeconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel2PresharedKey() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel2PresharedKey",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel2RekeyFuzzPercentage() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel2RekeyFuzzPercentage",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel2RekeyMarginTimeSeconds() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel2RekeyMarginTimeSeconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel2ReplayWindowSize() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel2ReplayWindowSize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnel2StartupAction() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnel2StartupAction",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetTunnelInsideIpVersion() {
	_jsii_.InvokeVoid(
		v,
		"resetTunnelInsideIpVersion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) ResetVpnGatewayId() {
	_jsii_.InvokeVoid(
		v,
		"resetVpnGatewayId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

