// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2clientvpnendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/dataawsec2clientvpnendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/data-sources/ec2_client_vpn_endpoint aws_ec2_client_vpn_endpoint}.
type DataAwsEc2ClientVpnEndpoint interface {
	cdktf.TerraformDataSource
	Arn() *string
	AuthenticationOptions() DataAwsEc2ClientVpnEndpointAuthenticationOptionsList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientCidrBlock() *string
	ClientConnectOptions() DataAwsEc2ClientVpnEndpointClientConnectOptionsList
	ClientLoginBannerOptions() DataAwsEc2ClientVpnEndpointClientLoginBannerOptionsList
	ClientVpnEndpointId() *string
	SetClientVpnEndpointId(val *string)
	ClientVpnEndpointIdInput() *string
	ConnectionLogOptions() DataAwsEc2ClientVpnEndpointConnectionLogOptionsList
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
	DnsName() *string
	DnsServers() *[]*string
	Filter() DataAwsEc2ClientVpnEndpointFilterList
	FilterInput() interface{}
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
	RawOverrides() interface{}
	SecurityGroupIds() *[]*string
	SelfServicePortal() *string
	ServerCertificateArn() *string
	SessionTimeoutHours() *float64
	SplitTunnel() cdktf.IResolvable
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataAwsEc2ClientVpnEndpointTimeoutsOutputReference
	TimeoutsInput() interface{}
	TransportProtocol() *string
	VpcId() *string
	VpnPort() *float64
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutFilter(value interface{})
	PutTimeouts(value *DataAwsEc2ClientVpnEndpointTimeouts)
	ResetClientVpnEndpointId()
	ResetFilter()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsEc2ClientVpnEndpoint
type jsiiProxy_DataAwsEc2ClientVpnEndpoint struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) AuthenticationOptions() DataAwsEc2ClientVpnEndpointAuthenticationOptionsList {
	var returns DataAwsEc2ClientVpnEndpointAuthenticationOptionsList
	_jsii_.Get(
		j,
		"authenticationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) ClientCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) ClientConnectOptions() DataAwsEc2ClientVpnEndpointClientConnectOptionsList {
	var returns DataAwsEc2ClientVpnEndpointClientConnectOptionsList
	_jsii_.Get(
		j,
		"clientConnectOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) ClientLoginBannerOptions() DataAwsEc2ClientVpnEndpointClientLoginBannerOptionsList {
	var returns DataAwsEc2ClientVpnEndpointClientLoginBannerOptionsList
	_jsii_.Get(
		j,
		"clientLoginBannerOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) ClientVpnEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientVpnEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) ClientVpnEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientVpnEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) ConnectionLogOptions() DataAwsEc2ClientVpnEndpointConnectionLogOptionsList {
	var returns DataAwsEc2ClientVpnEndpointConnectionLogOptionsList
	_jsii_.Get(
		j,
		"connectionLogOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) DnsServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) Filter() DataAwsEc2ClientVpnEndpointFilterList {
	var returns DataAwsEc2ClientVpnEndpointFilterList
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) SelfServicePortal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfServicePortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) ServerCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) SessionTimeoutHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) SplitTunnel() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"splitTunnel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) Timeouts() DataAwsEc2ClientVpnEndpointTimeoutsOutputReference {
	var returns DataAwsEc2ClientVpnEndpointTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) TransportProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transportProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint) VpnPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vpnPort",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/data-sources/ec2_client_vpn_endpoint aws_ec2_client_vpn_endpoint} Data Source.
func NewDataAwsEc2ClientVpnEndpoint(scope constructs.Construct, id *string, config *DataAwsEc2ClientVpnEndpointConfig) DataAwsEc2ClientVpnEndpoint {
	_init_.Initialize()

	if err := validateNewDataAwsEc2ClientVpnEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsEc2ClientVpnEndpoint{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsEc2ClientVpnEndpoint.DataAwsEc2ClientVpnEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/data-sources/ec2_client_vpn_endpoint aws_ec2_client_vpn_endpoint} Data Source.
func NewDataAwsEc2ClientVpnEndpoint_Override(d DataAwsEc2ClientVpnEndpoint, scope constructs.Construct, id *string, config *DataAwsEc2ClientVpnEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsEc2ClientVpnEndpoint.DataAwsEc2ClientVpnEndpoint",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint)SetClientVpnEndpointId(val *string) {
	if err := j.validateSetClientVpnEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientVpnEndpointId",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2ClientVpnEndpoint)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsEc2ClientVpnEndpoint resource upon running "cdktf plan <stack-name>".
func DataAwsEc2ClientVpnEndpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsEc2ClientVpnEndpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsEc2ClientVpnEndpoint.DataAwsEc2ClientVpnEndpoint",
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
func DataAwsEc2ClientVpnEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsEc2ClientVpnEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsEc2ClientVpnEndpoint.DataAwsEc2ClientVpnEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsEc2ClientVpnEndpoint_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsEc2ClientVpnEndpoint_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsEc2ClientVpnEndpoint.DataAwsEc2ClientVpnEndpoint",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsEc2ClientVpnEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsEc2ClientVpnEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsEc2ClientVpnEndpoint.DataAwsEc2ClientVpnEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsEc2ClientVpnEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsEc2ClientVpnEndpoint.DataAwsEc2ClientVpnEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) PutFilter(value interface{}) {
	if err := d.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFilter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) PutTimeouts(value *DataAwsEc2ClientVpnEndpointTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) ResetClientVpnEndpointId() {
	_jsii_.InvokeVoid(
		d,
		"resetClientVpnEndpointId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) ResetFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2ClientVpnEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

