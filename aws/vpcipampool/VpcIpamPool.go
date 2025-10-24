// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcipampool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/vpcipampool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/vpc_ipam_pool aws_vpc_ipam_pool}.
type VpcIpamPool interface {
	cdktf.TerraformResource
	AddressFamily() *string
	SetAddressFamily(val *string)
	AddressFamilyInput() *string
	AllocationDefaultNetmaskLength() *float64
	SetAllocationDefaultNetmaskLength(val *float64)
	AllocationDefaultNetmaskLengthInput() *float64
	AllocationMaxNetmaskLength() *float64
	SetAllocationMaxNetmaskLength(val *float64)
	AllocationMaxNetmaskLengthInput() *float64
	AllocationMinNetmaskLength() *float64
	SetAllocationMinNetmaskLength(val *float64)
	AllocationMinNetmaskLengthInput() *float64
	AllocationResourceTags() *map[string]*string
	SetAllocationResourceTags(val *map[string]*string)
	AllocationResourceTagsInput() *map[string]*string
	Arn() *string
	AutoImport() interface{}
	SetAutoImport(val interface{})
	AutoImportInput() interface{}
	AwsService() *string
	SetAwsService(val *string)
	AwsServiceInput() *string
	Cascade() interface{}
	SetCascade(val interface{})
	CascadeInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	IpamScopeId() *string
	SetIpamScopeId(val *string)
	IpamScopeIdInput() *string
	IpamScopeType() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Locale() *string
	SetLocale(val *string)
	LocaleInput() *string
	// The tree node.
	Node() constructs.Node
	PoolDepth() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicIpSource() *string
	SetPublicIpSource(val *string)
	PublicIpSourceInput() *string
	PubliclyAdvertisable() interface{}
	SetPubliclyAdvertisable(val interface{})
	PubliclyAdvertisableInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SourceIpamPoolId() *string
	SetSourceIpamPoolId(val *string)
	SourceIpamPoolIdInput() *string
	State() *string
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
	Timeouts() VpcIpamPoolTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutTimeouts(value *VpcIpamPoolTimeouts)
	ResetAllocationDefaultNetmaskLength()
	ResetAllocationMaxNetmaskLength()
	ResetAllocationMinNetmaskLength()
	ResetAllocationResourceTags()
	ResetAutoImport()
	ResetAwsService()
	ResetCascade()
	ResetDescription()
	ResetId()
	ResetLocale()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicIpSource()
	ResetPubliclyAdvertisable()
	ResetRegion()
	ResetSourceIpamPoolId()
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

// The jsii proxy struct for VpcIpamPool
type jsiiProxy_VpcIpamPool struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VpcIpamPool) AddressFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AddressFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AllocationDefaultNetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationDefaultNetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AllocationDefaultNetmaskLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationDefaultNetmaskLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AllocationMaxNetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationMaxNetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AllocationMaxNetmaskLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationMaxNetmaskLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AllocationMinNetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationMinNetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AllocationMinNetmaskLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationMinNetmaskLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AllocationResourceTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"allocationResourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AllocationResourceTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"allocationResourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AutoImport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoImport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AutoImportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoImportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AwsService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AwsServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Cascade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cascade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) CascadeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cascadeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) IpamScopeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamScopeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) IpamScopeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamScopeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) IpamScopeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamScopeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Locale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) LocaleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) PoolDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"poolDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) PublicIpSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) PublicIpSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) PubliclyAdvertisable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAdvertisable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) PubliclyAdvertisableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAdvertisableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) SourceIpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIpamPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) SourceIpamPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIpamPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Timeouts() VpcIpamPoolTimeoutsOutputReference {
	var returns VpcIpamPoolTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/vpc_ipam_pool aws_vpc_ipam_pool} Resource.
func NewVpcIpamPool(scope constructs.Construct, id *string, config *VpcIpamPoolConfig) VpcIpamPool {
	_init_.Initialize()

	if err := validateNewVpcIpamPoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpcIpamPool{}

	_jsii_.Create(
		"@cdktf/provider-aws.vpcIpamPool.VpcIpamPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/vpc_ipam_pool aws_vpc_ipam_pool} Resource.
func NewVpcIpamPool_Override(v VpcIpamPool, scope constructs.Construct, id *string, config *VpcIpamPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.vpcIpamPool.VpcIpamPool",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetAddressFamily(val *string) {
	if err := j.validateSetAddressFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressFamily",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetAllocationDefaultNetmaskLength(val *float64) {
	if err := j.validateSetAllocationDefaultNetmaskLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationDefaultNetmaskLength",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetAllocationMaxNetmaskLength(val *float64) {
	if err := j.validateSetAllocationMaxNetmaskLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationMaxNetmaskLength",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetAllocationMinNetmaskLength(val *float64) {
	if err := j.validateSetAllocationMinNetmaskLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationMinNetmaskLength",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetAllocationResourceTags(val *map[string]*string) {
	if err := j.validateSetAllocationResourceTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationResourceTags",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetAutoImport(val interface{}) {
	if err := j.validateSetAutoImportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoImport",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetAwsService(val *string) {
	if err := j.validateSetAwsServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsService",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetCascade(val interface{}) {
	if err := j.validateSetCascadeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cascade",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetIpamScopeId(val *string) {
	if err := j.validateSetIpamScopeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipamScopeId",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetLocale(val *string) {
	if err := j.validateSetLocaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locale",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetPublicIpSource(val *string) {
	if err := j.validateSetPublicIpSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIpSource",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetPubliclyAdvertisable(val interface{}) {
	if err := j.validateSetPubliclyAdvertisableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAdvertisable",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetSourceIpamPoolId(val *string) {
	if err := j.validateSetSourceIpamPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIpamPoolId",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a VpcIpamPool resource upon running "cdktf plan <stack-name>".
func VpcIpamPool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVpcIpamPool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpcIpamPool.VpcIpamPool",
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
func VpcIpamPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcIpamPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpcIpamPool.VpcIpamPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpcIpamPool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcIpamPool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpcIpamPool.VpcIpamPool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpcIpamPool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcIpamPool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.vpcIpamPool.VpcIpamPool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VpcIpamPool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.vpcIpamPool.VpcIpamPool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VpcIpamPool) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VpcIpamPool) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VpcIpamPool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VpcIpamPool) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpcIpamPool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VpcIpamPool) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VpcIpamPool) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VpcIpamPool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VpcIpamPool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VpcIpamPool) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VpcIpamPool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VpcIpamPool) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VpcIpamPool) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpcIpamPool) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VpcIpamPool) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VpcIpamPool) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VpcIpamPool) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VpcIpamPool) PutTimeouts(value *VpcIpamPoolTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetAllocationDefaultNetmaskLength() {
	_jsii_.InvokeVoid(
		v,
		"resetAllocationDefaultNetmaskLength",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetAllocationMaxNetmaskLength() {
	_jsii_.InvokeVoid(
		v,
		"resetAllocationMaxNetmaskLength",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetAllocationMinNetmaskLength() {
	_jsii_.InvokeVoid(
		v,
		"resetAllocationMinNetmaskLength",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetAllocationResourceTags() {
	_jsii_.InvokeVoid(
		v,
		"resetAllocationResourceTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetAutoImport() {
	_jsii_.InvokeVoid(
		v,
		"resetAutoImport",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetAwsService() {
	_jsii_.InvokeVoid(
		v,
		"resetAwsService",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetCascade() {
	_jsii_.InvokeVoid(
		v,
		"resetCascade",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetDescription() {
	_jsii_.InvokeVoid(
		v,
		"resetDescription",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetLocale() {
	_jsii_.InvokeVoid(
		v,
		"resetLocale",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetPublicIpSource() {
	_jsii_.InvokeVoid(
		v,
		"resetPublicIpSource",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetPubliclyAdvertisable() {
	_jsii_.InvokeVoid(
		v,
		"resetPubliclyAdvertisable",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetRegion() {
	_jsii_.InvokeVoid(
		v,
		"resetRegion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetSourceIpamPoolId() {
	_jsii_.InvokeVoid(
		v,
		"resetSourceIpamPoolId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetTagsAll() {
	_jsii_.InvokeVoid(
		v,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

