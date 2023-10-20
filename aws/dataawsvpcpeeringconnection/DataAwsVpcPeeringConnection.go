// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsvpcpeeringconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/dataawsvpcpeeringconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/data-sources/vpc_peering_connection aws_vpc_peering_connection}.
type DataAwsVpcPeeringConnection interface {
	cdktf.TerraformDataSource
	Accepter() cdktf.BooleanMap
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CidrBlock() *string
	SetCidrBlock(val *string)
	CidrBlockInput() *string
	CidrBlockSet() DataAwsVpcPeeringConnectionCidrBlockSetList
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
	Filter() DataAwsVpcPeeringConnectionFilterList
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
	OwnerId() *string
	SetOwnerId(val *string)
	OwnerIdInput() *string
	PeerCidrBlock() *string
	SetPeerCidrBlock(val *string)
	PeerCidrBlockInput() *string
	PeerCidrBlockSet() DataAwsVpcPeeringConnectionPeerCidrBlockSetList
	PeerOwnerId() *string
	SetPeerOwnerId(val *string)
	PeerOwnerIdInput() *string
	PeerRegion() *string
	SetPeerRegion(val *string)
	PeerRegionInput() *string
	PeerVpcId() *string
	SetPeerVpcId(val *string)
	PeerVpcIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	Requester() cdktf.BooleanMap
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataAwsVpcPeeringConnectionTimeoutsOutputReference
	TimeoutsInput() interface{}
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
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
	PutTimeouts(value *DataAwsVpcPeeringConnectionTimeouts)
	ResetCidrBlock()
	ResetFilter()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwnerId()
	ResetPeerCidrBlock()
	ResetPeerOwnerId()
	ResetPeerRegion()
	ResetPeerVpcId()
	ResetRegion()
	ResetStatus()
	ResetTags()
	ResetTimeouts()
	ResetVpcId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsVpcPeeringConnection
type jsiiProxy_DataAwsVpcPeeringConnection struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) Accepter() cdktf.BooleanMap {
	var returns cdktf.BooleanMap
	_jsii_.Get(
		j,
		"accepter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) CidrBlockSet() DataAwsVpcPeeringConnectionCidrBlockSetList {
	var returns DataAwsVpcPeeringConnectionCidrBlockSetList
	_jsii_.Get(
		j,
		"cidrBlockSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) Filter() DataAwsVpcPeeringConnectionFilterList {
	var returns DataAwsVpcPeeringConnectionFilterList
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) OwnerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) PeerCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) PeerCidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerCidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) PeerCidrBlockSet() DataAwsVpcPeeringConnectionPeerCidrBlockSetList {
	var returns DataAwsVpcPeeringConnectionPeerCidrBlockSetList
	_jsii_.Get(
		j,
		"peerCidrBlockSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) PeerOwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) PeerOwnerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerOwnerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) PeerRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) PeerRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) PeerVpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) PeerVpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) Requester() cdktf.BooleanMap {
	var returns cdktf.BooleanMap
	_jsii_.Get(
		j,
		"requester",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) Timeouts() DataAwsVpcPeeringConnectionTimeoutsOutputReference {
	var returns DataAwsVpcPeeringConnectionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/data-sources/vpc_peering_connection aws_vpc_peering_connection} Data Source.
func NewDataAwsVpcPeeringConnection(scope constructs.Construct, id *string, config *DataAwsVpcPeeringConnectionConfig) DataAwsVpcPeeringConnection {
	_init_.Initialize()

	if err := validateNewDataAwsVpcPeeringConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsVpcPeeringConnection{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsVpcPeeringConnection.DataAwsVpcPeeringConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/data-sources/vpc_peering_connection aws_vpc_peering_connection} Data Source.
func NewDataAwsVpcPeeringConnection_Override(d DataAwsVpcPeeringConnection, scope constructs.Construct, id *string, config *DataAwsVpcPeeringConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsVpcPeeringConnection.DataAwsVpcPeeringConnection",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection)SetCidrBlock(val *string) {
	if err := j.validateSetCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrBlock",
		val,
	)
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection)SetOwnerId(val *string) {
	if err := j.validateSetOwnerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ownerId",
		val,
	)
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection)SetPeerCidrBlock(val *string) {
	if err := j.validateSetPeerCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerCidrBlock",
		val,
	)
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection)SetPeerOwnerId(val *string) {
	if err := j.validateSetPeerOwnerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerOwnerId",
		val,
	)
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection)SetPeerRegion(val *string) {
	if err := j.validateSetPeerRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerRegion",
		val,
	)
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection)SetPeerVpcId(val *string) {
	if err := j.validateSetPeerVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerVpcId",
		val,
	)
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DataAwsVpcPeeringConnection)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsVpcPeeringConnection resource upon running "cdktf plan <stack-name>".
func DataAwsVpcPeeringConnection_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsVpcPeeringConnection_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsVpcPeeringConnection.DataAwsVpcPeeringConnection",
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
func DataAwsVpcPeeringConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsVpcPeeringConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsVpcPeeringConnection.DataAwsVpcPeeringConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsVpcPeeringConnection_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsVpcPeeringConnection_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsVpcPeeringConnection.DataAwsVpcPeeringConnection",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsVpcPeeringConnection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsVpcPeeringConnection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsVpcPeeringConnection.DataAwsVpcPeeringConnection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsVpcPeeringConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsVpcPeeringConnection.DataAwsVpcPeeringConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsVpcPeeringConnection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsVpcPeeringConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsVpcPeeringConnection) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsVpcPeeringConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsVpcPeeringConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsVpcPeeringConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsVpcPeeringConnection) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsVpcPeeringConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsVpcPeeringConnection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsVpcPeeringConnection) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) PutFilter(value interface{}) {
	if err := d.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFilter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) PutTimeouts(value *DataAwsVpcPeeringConnectionTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) ResetCidrBlock() {
	_jsii_.InvokeVoid(
		d,
		"resetCidrBlock",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) ResetFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) ResetOwnerId() {
	_jsii_.InvokeVoid(
		d,
		"resetOwnerId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) ResetPeerCidrBlock() {
	_jsii_.InvokeVoid(
		d,
		"resetPeerCidrBlock",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) ResetPeerOwnerId() {
	_jsii_.InvokeVoid(
		d,
		"resetPeerOwnerId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) ResetPeerRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetPeerRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) ResetPeerVpcId() {
	_jsii_.InvokeVoid(
		d,
		"resetPeerVpcId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) ResetStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) ResetVpcId() {
	_jsii_.InvokeVoid(
		d,
		"resetVpcId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsVpcPeeringConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

