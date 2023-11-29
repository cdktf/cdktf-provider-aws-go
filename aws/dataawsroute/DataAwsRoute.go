// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/dataawsroute/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/data-sources/route aws_route}.
type DataAwsRoute interface {
	cdktf.TerraformDataSource
	CarrierGatewayId() *string
	SetCarrierGatewayId(val *string)
	CarrierGatewayIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CoreNetworkArn() *string
	SetCoreNetworkArn(val *string)
	CoreNetworkArnInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DestinationCidrBlock() *string
	SetDestinationCidrBlock(val *string)
	DestinationCidrBlockInput() *string
	DestinationIpv6CidrBlock() *string
	SetDestinationIpv6CidrBlock(val *string)
	DestinationIpv6CidrBlockInput() *string
	DestinationPrefixListId() *string
	SetDestinationPrefixListId(val *string)
	DestinationPrefixListIdInput() *string
	EgressOnlyGatewayId() *string
	SetEgressOnlyGatewayId(val *string)
	EgressOnlyGatewayIdInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayId() *string
	SetGatewayId(val *string)
	GatewayIdInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceId() *string
	SetInstanceId(val *string)
	InstanceIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalGatewayId() *string
	SetLocalGatewayId(val *string)
	LocalGatewayIdInput() *string
	NatGatewayId() *string
	SetNatGatewayId(val *string)
	NatGatewayIdInput() *string
	NetworkInterfaceId() *string
	SetNetworkInterfaceId(val *string)
	NetworkInterfaceIdInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RouteTableId() *string
	SetRouteTableId(val *string)
	RouteTableIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataAwsRouteTimeoutsOutputReference
	TimeoutsInput() interface{}
	TransitGatewayId() *string
	SetTransitGatewayId(val *string)
	TransitGatewayIdInput() *string
	VpcPeeringConnectionId() *string
	SetVpcPeeringConnectionId(val *string)
	VpcPeeringConnectionIdInput() *string
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
	PutTimeouts(value *DataAwsRouteTimeouts)
	ResetCarrierGatewayId()
	ResetCoreNetworkArn()
	ResetDestinationCidrBlock()
	ResetDestinationIpv6CidrBlock()
	ResetDestinationPrefixListId()
	ResetEgressOnlyGatewayId()
	ResetGatewayId()
	ResetId()
	ResetInstanceId()
	ResetLocalGatewayId()
	ResetNatGatewayId()
	ResetNetworkInterfaceId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	ResetTransitGatewayId()
	ResetVpcPeeringConnectionId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsRoute
type jsiiProxy_DataAwsRoute struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsRoute) CarrierGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"carrierGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) CarrierGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"carrierGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) CoreNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) CoreNetworkArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) DestinationCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) DestinationCidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationCidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) DestinationIpv6CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationIpv6CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) DestinationIpv6CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationIpv6CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) DestinationPrefixListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPrefixListId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) DestinationPrefixListIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPrefixListIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) EgressOnlyGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egressOnlyGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) EgressOnlyGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egressOnlyGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) GatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) GatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) InstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) LocalGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) LocalGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) NatGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) NatGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) NetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) NetworkInterfaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) RouteTableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeTableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) RouteTableIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeTableIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) Timeouts() DataAwsRouteTimeoutsOutputReference {
	var returns DataAwsRouteTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) TransitGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) TransitGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) VpcPeeringConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcPeeringConnectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute) VpcPeeringConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcPeeringConnectionIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/data-sources/route aws_route} Data Source.
func NewDataAwsRoute(scope constructs.Construct, id *string, config *DataAwsRouteConfig) DataAwsRoute {
	_init_.Initialize()

	if err := validateNewDataAwsRouteParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsRoute{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsRoute.DataAwsRoute",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/data-sources/route aws_route} Data Source.
func NewDataAwsRoute_Override(d DataAwsRoute, scope constructs.Construct, id *string, config *DataAwsRouteConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsRoute.DataAwsRoute",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsRoute)SetCarrierGatewayId(val *string) {
	if err := j.validateSetCarrierGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"carrierGatewayId",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute)SetCoreNetworkArn(val *string) {
	if err := j.validateSetCoreNetworkArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreNetworkArn",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute)SetDestinationCidrBlock(val *string) {
	if err := j.validateSetDestinationCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationCidrBlock",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute)SetDestinationIpv6CidrBlock(val *string) {
	if err := j.validateSetDestinationIpv6CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationIpv6CidrBlock",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute)SetDestinationPrefixListId(val *string) {
	if err := j.validateSetDestinationPrefixListIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPrefixListId",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute)SetEgressOnlyGatewayId(val *string) {
	if err := j.validateSetEgressOnlyGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egressOnlyGatewayId",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute)SetGatewayId(val *string) {
	if err := j.validateSetGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayId",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute)SetInstanceId(val *string) {
	if err := j.validateSetInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute)SetLocalGatewayId(val *string) {
	if err := j.validateSetLocalGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localGatewayId",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute)SetNatGatewayId(val *string) {
	if err := j.validateSetNatGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"natGatewayId",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute)SetNetworkInterfaceId(val *string) {
	if err := j.validateSetNetworkInterfaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkInterfaceId",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute)SetRouteTableId(val *string) {
	if err := j.validateSetRouteTableIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeTableId",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute)SetTransitGatewayId(val *string) {
	if err := j.validateSetTransitGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitGatewayId",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute)SetVpcPeeringConnectionId(val *string) {
	if err := j.validateSetVpcPeeringConnectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcPeeringConnectionId",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsRoute resource upon running "cdktf plan <stack-name>".
func DataAwsRoute_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsRoute_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsRoute.DataAwsRoute",
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
func DataAwsRoute_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsRoute_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsRoute.DataAwsRoute",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsRoute_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsRoute_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsRoute.DataAwsRoute",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsRoute_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsRoute_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsRoute.DataAwsRoute",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsRoute_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsRoute.DataAwsRoute",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsRoute) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsRoute) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsRoute) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsRoute) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsRoute) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsRoute) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsRoute) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsRoute) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsRoute) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsRoute) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsRoute) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsRoute) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsRoute) PutTimeouts(value *DataAwsRouteTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsRoute) ResetCarrierGatewayId() {
	_jsii_.InvokeVoid(
		d,
		"resetCarrierGatewayId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute) ResetCoreNetworkArn() {
	_jsii_.InvokeVoid(
		d,
		"resetCoreNetworkArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute) ResetDestinationCidrBlock() {
	_jsii_.InvokeVoid(
		d,
		"resetDestinationCidrBlock",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute) ResetDestinationIpv6CidrBlock() {
	_jsii_.InvokeVoid(
		d,
		"resetDestinationIpv6CidrBlock",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute) ResetDestinationPrefixListId() {
	_jsii_.InvokeVoid(
		d,
		"resetDestinationPrefixListId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute) ResetEgressOnlyGatewayId() {
	_jsii_.InvokeVoid(
		d,
		"resetEgressOnlyGatewayId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute) ResetGatewayId() {
	_jsii_.InvokeVoid(
		d,
		"resetGatewayId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute) ResetInstanceId() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute) ResetLocalGatewayId() {
	_jsii_.InvokeVoid(
		d,
		"resetLocalGatewayId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute) ResetNatGatewayId() {
	_jsii_.InvokeVoid(
		d,
		"resetNatGatewayId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute) ResetNetworkInterfaceId() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkInterfaceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute) ResetTransitGatewayId() {
	_jsii_.InvokeVoid(
		d,
		"resetTransitGatewayId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute) ResetVpcPeeringConnectionId() {
	_jsii_.InvokeVoid(
		d,
		"resetVpcPeeringConnectionId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

