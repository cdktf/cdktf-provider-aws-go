// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/vpcendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_endpoint aws_vpc_endpoint}.
type VpcEndpoint interface {
	cdktf.TerraformResource
	Arn() *string
	AutoAccept() interface{}
	SetAutoAccept(val interface{})
	AutoAcceptInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CidrBlocks() *[]*string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
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
	DnsEntry() VpcEndpointDnsEntryList
	DnsOptions() VpcEndpointDnsOptionsOutputReference
	DnsOptionsInput() *VpcEndpointDnsOptions
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
	IpAddressType() *string
	SetIpAddressType(val *string)
	IpAddressTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	NetworkInterfaceIds() *[]*string
	// The tree node.
	Node() constructs.Node
	OwnerId() *string
	Policy() *string
	SetPolicy(val *string)
	PolicyInput() *string
	PrefixListId() *string
	PrivateDnsEnabled() interface{}
	SetPrivateDnsEnabled(val interface{})
	PrivateDnsEnabledInput() interface{}
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RequesterManaged() cdktf.IResolvable
	ResourceConfigurationArn() *string
	SetResourceConfigurationArn(val *string)
	ResourceConfigurationArnInput() *string
	RouteTableIds() *[]*string
	SetRouteTableIds(val *[]*string)
	RouteTableIdsInput() *[]*string
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	ServiceName() *string
	SetServiceName(val *string)
	ServiceNameInput() *string
	ServiceNetworkArn() *string
	SetServiceNetworkArn(val *string)
	ServiceNetworkArnInput() *string
	ServiceRegion() *string
	SetServiceRegion(val *string)
	ServiceRegionInput() *string
	State() *string
	SubnetConfiguration() VpcEndpointSubnetConfigurationList
	SubnetConfigurationInput() interface{}
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
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
	Timeouts() VpcEndpointTimeoutsOutputReference
	TimeoutsInput() interface{}
	VpcEndpointType() *string
	SetVpcEndpointType(val *string)
	VpcEndpointTypeInput() *string
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutDnsOptions(value *VpcEndpointDnsOptions)
	PutSubnetConfiguration(value interface{})
	PutTimeouts(value *VpcEndpointTimeouts)
	ResetAutoAccept()
	ResetDnsOptions()
	ResetId()
	ResetIpAddressType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicy()
	ResetPrivateDnsEnabled()
	ResetRegion()
	ResetResourceConfigurationArn()
	ResetRouteTableIds()
	ResetSecurityGroupIds()
	ResetServiceName()
	ResetServiceNetworkArn()
	ResetServiceRegion()
	ResetSubnetConfiguration()
	ResetSubnetIds()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetVpcEndpointType()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VpcEndpoint
type jsiiProxy_VpcEndpoint struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VpcEndpoint) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) AutoAccept() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAccept",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) AutoAcceptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAcceptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) CidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) DnsEntry() VpcEndpointDnsEntryList {
	var returns VpcEndpointDnsEntryList
	_jsii_.Get(
		j,
		"dnsEntry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) DnsOptions() VpcEndpointDnsOptionsOutputReference {
	var returns VpcEndpointDnsOptionsOutputReference
	_jsii_.Get(
		j,
		"dnsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) DnsOptionsInput() *VpcEndpointDnsOptions {
	var returns *VpcEndpointDnsOptions
	_jsii_.Get(
		j,
		"dnsOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) IpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) NetworkInterfaceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkInterfaceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) PrefixListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixListId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) PrivateDnsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateDnsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) PrivateDnsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateDnsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) RequesterManaged() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"requesterManaged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) ResourceConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) ResourceConfigurationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceConfigurationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) RouteTableIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"routeTableIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) RouteTableIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"routeTableIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) ServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) ServiceNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) ServiceNetworkArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNetworkArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) ServiceRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) ServiceRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) SubnetConfiguration() VpcEndpointSubnetConfigurationList {
	var returns VpcEndpointSubnetConfigurationList
	_jsii_.Get(
		j,
		"subnetConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) SubnetConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subnetConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) Timeouts() VpcEndpointTimeoutsOutputReference {
	var returns VpcEndpointTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) VpcEndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) VpcEndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcEndpoint) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_endpoint aws_vpc_endpoint} Resource.
func NewVpcEndpoint(scope constructs.Construct, id *string, config *VpcEndpointConfig) VpcEndpoint {
	_init_.Initialize()

	if err := validateNewVpcEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpcEndpoint{}

	_jsii_.Create(
		"@cdktf/provider-aws.vpcEndpoint.VpcEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_endpoint aws_vpc_endpoint} Resource.
func NewVpcEndpoint_Override(v VpcEndpoint, scope constructs.Construct, id *string, config *VpcEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.vpcEndpoint.VpcEndpoint",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetAutoAccept(val interface{}) {
	if err := j.validateSetAutoAcceptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoAccept",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetIpAddressType(val *string) {
	if err := j.validateSetIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetPolicy(val *string) {
	if err := j.validateSetPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetPrivateDnsEnabled(val interface{}) {
	if err := j.validateSetPrivateDnsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateDnsEnabled",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetResourceConfigurationArn(val *string) {
	if err := j.validateSetResourceConfigurationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceConfigurationArn",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetRouteTableIds(val *[]*string) {
	if err := j.validateSetRouteTableIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeTableIds",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetServiceName(val *string) {
	if err := j.validateSetServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceName",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetServiceNetworkArn(val *string) {
	if err := j.validateSetServiceNetworkArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceNetworkArn",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetServiceRegion(val *string) {
	if err := j.validateSetServiceRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRegion",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetVpcEndpointType(val *string) {
	if err := j.validateSetVpcEndpointTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcEndpointType",
		val,
	)
}

func (j *jsiiProxy_VpcEndpoint)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a VpcEndpoint resource upon running "cdktf plan <stack-name>".
func VpcEndpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVpcEndpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpcEndpoint.VpcEndpoint",
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
func VpcEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpcEndpoint.VpcEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpcEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpcEndpoint.VpcEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpcEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpcEndpoint.VpcEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VpcEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.vpcEndpoint.VpcEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VpcEndpoint) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VpcEndpoint) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VpcEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VpcEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpcEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VpcEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VpcEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VpcEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VpcEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VpcEndpoint) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VpcEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VpcEndpoint) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEndpoint) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VpcEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpcEndpoint) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VpcEndpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VpcEndpoint) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VpcEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VpcEndpoint) PutDnsOptions(value *VpcEndpointDnsOptions) {
	if err := v.validatePutDnsOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putDnsOptions",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpcEndpoint) PutSubnetConfiguration(value interface{}) {
	if err := v.validatePutSubnetConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putSubnetConfiguration",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpcEndpoint) PutTimeouts(value *VpcEndpointTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpcEndpoint) ResetAutoAccept() {
	_jsii_.InvokeVoid(
		v,
		"resetAutoAccept",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpoint) ResetDnsOptions() {
	_jsii_.InvokeVoid(
		v,
		"resetDnsOptions",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpoint) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpoint) ResetIpAddressType() {
	_jsii_.InvokeVoid(
		v,
		"resetIpAddressType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpoint) ResetPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpoint) ResetPrivateDnsEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetPrivateDnsEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpoint) ResetRegion() {
	_jsii_.InvokeVoid(
		v,
		"resetRegion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpoint) ResetResourceConfigurationArn() {
	_jsii_.InvokeVoid(
		v,
		"resetResourceConfigurationArn",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpoint) ResetRouteTableIds() {
	_jsii_.InvokeVoid(
		v,
		"resetRouteTableIds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpoint) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		v,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpoint) ResetServiceName() {
	_jsii_.InvokeVoid(
		v,
		"resetServiceName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpoint) ResetServiceNetworkArn() {
	_jsii_.InvokeVoid(
		v,
		"resetServiceNetworkArn",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpoint) ResetServiceRegion() {
	_jsii_.InvokeVoid(
		v,
		"resetServiceRegion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpoint) ResetSubnetConfiguration() {
	_jsii_.InvokeVoid(
		v,
		"resetSubnetConfiguration",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpoint) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		v,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpoint) ResetTagsAll() {
	_jsii_.InvokeVoid(
		v,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpoint) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpoint) ResetVpcEndpointType() {
	_jsii_.InvokeVoid(
		v,
		"resetVpcEndpointType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEndpoint) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEndpoint) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

