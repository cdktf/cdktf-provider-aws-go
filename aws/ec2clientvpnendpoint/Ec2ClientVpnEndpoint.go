// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2clientvpnendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ec2clientvpnendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/ec2_client_vpn_endpoint aws_ec2_client_vpn_endpoint}.
type Ec2ClientVpnEndpoint interface {
	cdktf.TerraformResource
	Arn() *string
	AuthenticationOptions() Ec2ClientVpnEndpointAuthenticationOptionsList
	AuthenticationOptionsInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientCidrBlock() *string
	SetClientCidrBlock(val *string)
	ClientCidrBlockInput() *string
	ClientConnectOptions() Ec2ClientVpnEndpointClientConnectOptionsOutputReference
	ClientConnectOptionsInput() *Ec2ClientVpnEndpointClientConnectOptions
	ClientLoginBannerOptions() Ec2ClientVpnEndpointClientLoginBannerOptionsOutputReference
	ClientLoginBannerOptionsInput() *Ec2ClientVpnEndpointClientLoginBannerOptions
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionLogOptions() Ec2ClientVpnEndpointConnectionLogOptionsOutputReference
	ConnectionLogOptionsInput() *Ec2ClientVpnEndpointConnectionLogOptions
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DnsName() *string
	DnsServers() *[]*string
	SetDnsServers(val *[]*string)
	DnsServersInput() *[]*string
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
	// The tree node.
	Node() constructs.Node
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
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SelfServicePortal() *string
	SetSelfServicePortal(val *string)
	SelfServicePortalInput() *string
	SelfServicePortalUrl() *string
	ServerCertificateArn() *string
	SetServerCertificateArn(val *string)
	ServerCertificateArnInput() *string
	SessionTimeoutHours() *float64
	SetSessionTimeoutHours(val *float64)
	SessionTimeoutHoursInput() *float64
	SplitTunnel() interface{}
	SetSplitTunnel(val interface{})
	SplitTunnelInput() interface{}
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
	TransportProtocol() *string
	SetTransportProtocol(val *string)
	TransportProtocolInput() *string
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
	VpnPort() *float64
	SetVpnPort(val *float64)
	VpnPortInput() *float64
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
	PutAuthenticationOptions(value interface{})
	PutClientConnectOptions(value *Ec2ClientVpnEndpointClientConnectOptions)
	PutClientLoginBannerOptions(value *Ec2ClientVpnEndpointClientLoginBannerOptions)
	PutConnectionLogOptions(value *Ec2ClientVpnEndpointConnectionLogOptions)
	ResetClientConnectOptions()
	ResetClientLoginBannerOptions()
	ResetDescription()
	ResetDnsServers()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSecurityGroupIds()
	ResetSelfServicePortal()
	ResetSessionTimeoutHours()
	ResetSplitTunnel()
	ResetTags()
	ResetTagsAll()
	ResetTransportProtocol()
	ResetVpcId()
	ResetVpnPort()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Ec2ClientVpnEndpoint
type jsiiProxy_Ec2ClientVpnEndpoint struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) AuthenticationOptions() Ec2ClientVpnEndpointAuthenticationOptionsList {
	var returns Ec2ClientVpnEndpointAuthenticationOptionsList
	_jsii_.Get(
		j,
		"authenticationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) AuthenticationOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) ClientCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) ClientCidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) ClientConnectOptions() Ec2ClientVpnEndpointClientConnectOptionsOutputReference {
	var returns Ec2ClientVpnEndpointClientConnectOptionsOutputReference
	_jsii_.Get(
		j,
		"clientConnectOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) ClientConnectOptionsInput() *Ec2ClientVpnEndpointClientConnectOptions {
	var returns *Ec2ClientVpnEndpointClientConnectOptions
	_jsii_.Get(
		j,
		"clientConnectOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) ClientLoginBannerOptions() Ec2ClientVpnEndpointClientLoginBannerOptionsOutputReference {
	var returns Ec2ClientVpnEndpointClientLoginBannerOptionsOutputReference
	_jsii_.Get(
		j,
		"clientLoginBannerOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) ClientLoginBannerOptionsInput() *Ec2ClientVpnEndpointClientLoginBannerOptions {
	var returns *Ec2ClientVpnEndpointClientLoginBannerOptions
	_jsii_.Get(
		j,
		"clientLoginBannerOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) ConnectionLogOptions() Ec2ClientVpnEndpointConnectionLogOptionsOutputReference {
	var returns Ec2ClientVpnEndpointConnectionLogOptionsOutputReference
	_jsii_.Get(
		j,
		"connectionLogOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) ConnectionLogOptionsInput() *Ec2ClientVpnEndpointConnectionLogOptions {
	var returns *Ec2ClientVpnEndpointConnectionLogOptions
	_jsii_.Get(
		j,
		"connectionLogOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) DnsServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) DnsServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) SelfServicePortal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfServicePortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) SelfServicePortalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfServicePortalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) SelfServicePortalUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfServicePortalUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) ServerCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) ServerCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) SessionTimeoutHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) SessionTimeoutHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) SplitTunnel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splitTunnel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) SplitTunnelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splitTunnelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) TransportProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transportProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) TransportProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transportProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) VpnPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vpnPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint) VpnPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vpnPortInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/ec2_client_vpn_endpoint aws_ec2_client_vpn_endpoint} Resource.
func NewEc2ClientVpnEndpoint(scope constructs.Construct, id *string, config *Ec2ClientVpnEndpointConfig) Ec2ClientVpnEndpoint {
	_init_.Initialize()

	if err := validateNewEc2ClientVpnEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2ClientVpnEndpoint{}

	_jsii_.Create(
		"@cdktf/provider-aws.ec2ClientVpnEndpoint.Ec2ClientVpnEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/ec2_client_vpn_endpoint aws_ec2_client_vpn_endpoint} Resource.
func NewEc2ClientVpnEndpoint_Override(e Ec2ClientVpnEndpoint, scope constructs.Construct, id *string, config *Ec2ClientVpnEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ec2ClientVpnEndpoint.Ec2ClientVpnEndpoint",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetClientCidrBlock(val *string) {
	if err := j.validateSetClientCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCidrBlock",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetDnsServers(val *[]*string) {
	if err := j.validateSetDnsServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServers",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetSelfServicePortal(val *string) {
	if err := j.validateSetSelfServicePortalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfServicePortal",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetServerCertificateArn(val *string) {
	if err := j.validateSetServerCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverCertificateArn",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetSessionTimeoutHours(val *float64) {
	if err := j.validateSetSessionTimeoutHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionTimeoutHours",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetSplitTunnel(val interface{}) {
	if err := j.validateSetSplitTunnelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"splitTunnel",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetTransportProtocol(val *string) {
	if err := j.validateSetTransportProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transportProtocol",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

func (j *jsiiProxy_Ec2ClientVpnEndpoint)SetVpnPort(val *float64) {
	if err := j.validateSetVpnPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnPort",
		val,
	)
}

// Generates CDKTF code for importing a Ec2ClientVpnEndpoint resource upon running "cdktf plan <stack-name>".
func Ec2ClientVpnEndpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEc2ClientVpnEndpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ec2ClientVpnEndpoint.Ec2ClientVpnEndpoint",
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
func Ec2ClientVpnEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2ClientVpnEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ec2ClientVpnEndpoint.Ec2ClientVpnEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2ClientVpnEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2ClientVpnEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ec2ClientVpnEndpoint.Ec2ClientVpnEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2ClientVpnEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2ClientVpnEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ec2ClientVpnEndpoint.Ec2ClientVpnEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Ec2ClientVpnEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.ec2ClientVpnEndpoint.Ec2ClientVpnEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) PutAuthenticationOptions(value interface{}) {
	if err := e.validatePutAuthenticationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAuthenticationOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) PutClientConnectOptions(value *Ec2ClientVpnEndpointClientConnectOptions) {
	if err := e.validatePutClientConnectOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putClientConnectOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) PutClientLoginBannerOptions(value *Ec2ClientVpnEndpointClientLoginBannerOptions) {
	if err := e.validatePutClientLoginBannerOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putClientLoginBannerOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) PutConnectionLogOptions(value *Ec2ClientVpnEndpointConnectionLogOptions) {
	if err := e.validatePutConnectionLogOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putConnectionLogOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) ResetClientConnectOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetClientConnectOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) ResetClientLoginBannerOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetClientLoginBannerOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) ResetDnsServers() {
	_jsii_.InvokeVoid(
		e,
		"resetDnsServers",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		e,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) ResetSelfServicePortal() {
	_jsii_.InvokeVoid(
		e,
		"resetSelfServicePortal",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) ResetSessionTimeoutHours() {
	_jsii_.InvokeVoid(
		e,
		"resetSessionTimeoutHours",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) ResetSplitTunnel() {
	_jsii_.InvokeVoid(
		e,
		"resetSplitTunnel",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) ResetTagsAll() {
	_jsii_.InvokeVoid(
		e,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) ResetTransportProtocol() {
	_jsii_.InvokeVoid(
		e,
		"resetTransportProtocol",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) ResetVpcId() {
	_jsii_.InvokeVoid(
		e,
		"resetVpcId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) ResetVpnPort() {
	_jsii_.InvokeVoid(
		e,
		"resetVpnPort",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ClientVpnEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

