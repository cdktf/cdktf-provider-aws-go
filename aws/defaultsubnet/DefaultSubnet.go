// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package defaultsubnet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/defaultsubnet/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/default_subnet aws_default_subnet}.
type DefaultSubnet interface {
	cdktf.TerraformResource
	Arn() *string
	AssignIpv6AddressOnCreation() interface{}
	SetAssignIpv6AddressOnCreation(val interface{})
	AssignIpv6AddressOnCreationInput() interface{}
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneId() *string
	AvailabilityZoneInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CidrBlock() *string
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
	CustomerOwnedIpv4Pool() *string
	SetCustomerOwnedIpv4Pool(val *string)
	CustomerOwnedIpv4PoolInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableDns64() interface{}
	SetEnableDns64(val interface{})
	EnableDns64Input() interface{}
	EnableLniAtDeviceIndex() *float64
	EnableResourceNameDnsAaaaRecordOnLaunch() interface{}
	SetEnableResourceNameDnsAaaaRecordOnLaunch(val interface{})
	EnableResourceNameDnsAaaaRecordOnLaunchInput() interface{}
	EnableResourceNameDnsARecordOnLaunch() interface{}
	SetEnableResourceNameDnsARecordOnLaunch(val interface{})
	EnableResourceNameDnsARecordOnLaunchInput() interface{}
	ExistingDefaultSubnet() cdktf.IResolvable
	ForceDestroy() interface{}
	SetForceDestroy(val interface{})
	ForceDestroyInput() interface{}
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
	Ipv6CidrBlock() *string
	SetIpv6CidrBlock(val *string)
	Ipv6CidrBlockAssociationId() *string
	Ipv6CidrBlockInput() *string
	Ipv6Native() interface{}
	SetIpv6Native(val interface{})
	Ipv6NativeInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MapCustomerOwnedIpOnLaunch() interface{}
	SetMapCustomerOwnedIpOnLaunch(val interface{})
	MapCustomerOwnedIpOnLaunchInput() interface{}
	MapPublicIpOnLaunch() interface{}
	SetMapPublicIpOnLaunch(val interface{})
	MapPublicIpOnLaunchInput() interface{}
	// The tree node.
	Node() constructs.Node
	OutpostArn() *string
	OwnerId() *string
	PrivateDnsHostnameTypeOnLaunch() *string
	SetPrivateDnsHostnameTypeOnLaunch(val *string)
	PrivateDnsHostnameTypeOnLaunchInput() *string
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
	Timeouts() DefaultSubnetTimeoutsOutputReference
	TimeoutsInput() interface{}
	VpcId() *string
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
	PutTimeouts(value *DefaultSubnetTimeouts)
	ResetAssignIpv6AddressOnCreation()
	ResetCustomerOwnedIpv4Pool()
	ResetEnableDns64()
	ResetEnableResourceNameDnsAaaaRecordOnLaunch()
	ResetEnableResourceNameDnsARecordOnLaunch()
	ResetForceDestroy()
	ResetId()
	ResetIpv6CidrBlock()
	ResetIpv6Native()
	ResetMapCustomerOwnedIpOnLaunch()
	ResetMapPublicIpOnLaunch()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateDnsHostnameTypeOnLaunch()
	ResetRegion()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
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

// The jsii proxy struct for DefaultSubnet
type jsiiProxy_DefaultSubnet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DefaultSubnet) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) AssignIpv6AddressOnCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignIpv6AddressOnCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) AssignIpv6AddressOnCreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignIpv6AddressOnCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) AvailabilityZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) CustomerOwnedIpv4Pool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerOwnedIpv4Pool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) CustomerOwnedIpv4PoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerOwnedIpv4PoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) EnableDns64() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDns64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) EnableDns64Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDns64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) EnableLniAtDeviceIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enableLniAtDeviceIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) EnableResourceNameDnsAaaaRecordOnLaunch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsAaaaRecordOnLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) EnableResourceNameDnsAaaaRecordOnLaunchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsAaaaRecordOnLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) EnableResourceNameDnsARecordOnLaunch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsARecordOnLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) EnableResourceNameDnsARecordOnLaunchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsARecordOnLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) ExistingDefaultSubnet() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"existingDefaultSubnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Ipv6CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Ipv6CidrBlockAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlockAssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Ipv6CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Ipv6Native() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6Native",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Ipv6NativeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6NativeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) MapCustomerOwnedIpOnLaunch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapCustomerOwnedIpOnLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) MapCustomerOwnedIpOnLaunchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapCustomerOwnedIpOnLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) MapPublicIpOnLaunch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapPublicIpOnLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) MapPublicIpOnLaunchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapPublicIpOnLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) OutpostArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) PrivateDnsHostnameTypeOnLaunch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsHostnameTypeOnLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) PrivateDnsHostnameTypeOnLaunchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsHostnameTypeOnLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) Timeouts() DefaultSubnetTimeoutsOutputReference {
	var returns DefaultSubnetTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultSubnet) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/default_subnet aws_default_subnet} Resource.
func NewDefaultSubnet(scope constructs.Construct, id *string, config *DefaultSubnetConfig) DefaultSubnet {
	_init_.Initialize()

	if err := validateNewDefaultSubnetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DefaultSubnet{}

	_jsii_.Create(
		"@cdktf/provider-aws.defaultSubnet.DefaultSubnet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/default_subnet aws_default_subnet} Resource.
func NewDefaultSubnet_Override(d DefaultSubnet, scope constructs.Construct, id *string, config *DefaultSubnetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.defaultSubnet.DefaultSubnet",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetAssignIpv6AddressOnCreation(val interface{}) {
	if err := j.validateSetAssignIpv6AddressOnCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assignIpv6AddressOnCreation",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetCustomerOwnedIpv4Pool(val *string) {
	if err := j.validateSetCustomerOwnedIpv4PoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerOwnedIpv4Pool",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetEnableDns64(val interface{}) {
	if err := j.validateSetEnableDns64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDns64",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetEnableResourceNameDnsAaaaRecordOnLaunch(val interface{}) {
	if err := j.validateSetEnableResourceNameDnsAaaaRecordOnLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourceNameDnsAaaaRecordOnLaunch",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetEnableResourceNameDnsARecordOnLaunch(val interface{}) {
	if err := j.validateSetEnableResourceNameDnsARecordOnLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourceNameDnsARecordOnLaunch",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetIpv6CidrBlock(val *string) {
	if err := j.validateSetIpv6CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6CidrBlock",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetIpv6Native(val interface{}) {
	if err := j.validateSetIpv6NativeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Native",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetMapCustomerOwnedIpOnLaunch(val interface{}) {
	if err := j.validateSetMapCustomerOwnedIpOnLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapCustomerOwnedIpOnLaunch",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetMapPublicIpOnLaunch(val interface{}) {
	if err := j.validateSetMapPublicIpOnLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapPublicIpOnLaunch",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetPrivateDnsHostnameTypeOnLaunch(val *string) {
	if err := j.validateSetPrivateDnsHostnameTypeOnLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateDnsHostnameTypeOnLaunch",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DefaultSubnet)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a DefaultSubnet resource upon running "cdktf plan <stack-name>".
func DefaultSubnet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDefaultSubnet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.defaultSubnet.DefaultSubnet",
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
func DefaultSubnet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDefaultSubnet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.defaultSubnet.DefaultSubnet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DefaultSubnet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDefaultSubnet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.defaultSubnet.DefaultSubnet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DefaultSubnet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDefaultSubnet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.defaultSubnet.DefaultSubnet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DefaultSubnet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.defaultSubnet.DefaultSubnet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DefaultSubnet) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DefaultSubnet) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DefaultSubnet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DefaultSubnet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DefaultSubnet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DefaultSubnet) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DefaultSubnet) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DefaultSubnet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DefaultSubnet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DefaultSubnet) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DefaultSubnet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DefaultSubnet) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultSubnet) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DefaultSubnet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DefaultSubnet) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DefaultSubnet) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DefaultSubnet) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DefaultSubnet) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DefaultSubnet) PutTimeouts(value *DefaultSubnetTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetAssignIpv6AddressOnCreation() {
	_jsii_.InvokeVoid(
		d,
		"resetAssignIpv6AddressOnCreation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetCustomerOwnedIpv4Pool() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomerOwnedIpv4Pool",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetEnableDns64() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableDns64",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetEnableResourceNameDnsAaaaRecordOnLaunch() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableResourceNameDnsAaaaRecordOnLaunch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetEnableResourceNameDnsARecordOnLaunch() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableResourceNameDnsARecordOnLaunch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		d,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetIpv6CidrBlock() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv6CidrBlock",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetIpv6Native() {
	_jsii_.InvokeVoid(
		d,
		"resetIpv6Native",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetMapCustomerOwnedIpOnLaunch() {
	_jsii_.InvokeVoid(
		d,
		"resetMapCustomerOwnedIpOnLaunch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetMapPublicIpOnLaunch() {
	_jsii_.InvokeVoid(
		d,
		"resetMapPublicIpOnLaunch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetPrivateDnsHostnameTypeOnLaunch() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivateDnsHostnameTypeOnLaunch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetTagsAll() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultSubnet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultSubnet) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultSubnet) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultSubnet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultSubnet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultSubnet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

