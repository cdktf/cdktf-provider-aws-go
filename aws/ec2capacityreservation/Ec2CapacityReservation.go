// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2capacityreservation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/ec2capacityreservation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/ec2_capacity_reservation aws_ec2_capacity_reservation}.
type Ec2CapacityReservation interface {
	cdktf.TerraformResource
	Arn() *string
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
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
	EbsOptimized() interface{}
	SetEbsOptimized(val interface{})
	EbsOptimizedInput() interface{}
	EndDate() *string
	SetEndDate(val *string)
	EndDateInput() *string
	EndDateType() *string
	SetEndDateType(val *string)
	EndDateTypeInput() *string
	EphemeralStorage() interface{}
	SetEphemeralStorage(val interface{})
	EphemeralStorageInput() interface{}
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
	InstanceCount() *float64
	SetInstanceCount(val *float64)
	InstanceCountInput() *float64
	InstanceMatchCriteria() *string
	SetInstanceMatchCriteria(val *string)
	InstanceMatchCriteriaInput() *string
	InstancePlatform() *string
	SetInstancePlatform(val *string)
	InstancePlatformInput() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OutpostArn() *string
	SetOutpostArn(val *string)
	OutpostArnInput() *string
	OwnerId() *string
	PlacementGroupArn() *string
	SetPlacementGroupArn(val *string)
	PlacementGroupArnInput() *string
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	Tenancy() *string
	SetTenancy(val *string)
	TenancyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() Ec2CapacityReservationTimeoutsOutputReference
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
	PutTimeouts(value *Ec2CapacityReservationTimeouts)
	ResetEbsOptimized()
	ResetEndDate()
	ResetEndDateType()
	ResetEphemeralStorage()
	ResetId()
	ResetInstanceMatchCriteria()
	ResetOutpostArn()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlacementGroupArn()
	ResetTags()
	ResetTagsAll()
	ResetTenancy()
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

// The jsii proxy struct for Ec2CapacityReservation
type jsiiProxy_Ec2CapacityReservation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Ec2CapacityReservation) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) EbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) EndDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) EndDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) EndDateType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) EndDateTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDateTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) EphemeralStorage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) EphemeralStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) InstanceMatchCriteria() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceMatchCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) InstanceMatchCriteriaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceMatchCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) InstancePlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) InstancePlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) OutpostArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) OutpostArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) PlacementGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) PlacementGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) Tenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) TenancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) Timeouts() Ec2CapacityReservationTimeoutsOutputReference {
	var returns Ec2CapacityReservationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservation) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/ec2_capacity_reservation aws_ec2_capacity_reservation} Resource.
func NewEc2CapacityReservation(scope constructs.Construct, id *string, config *Ec2CapacityReservationConfig) Ec2CapacityReservation {
	_init_.Initialize()

	if err := validateNewEc2CapacityReservationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2CapacityReservation{}

	_jsii_.Create(
		"@cdktf/provider-aws.ec2CapacityReservation.Ec2CapacityReservation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/ec2_capacity_reservation aws_ec2_capacity_reservation} Resource.
func NewEc2CapacityReservation_Override(e Ec2CapacityReservation, scope constructs.Construct, id *string, config *Ec2CapacityReservationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ec2CapacityReservation.Ec2CapacityReservation",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetEbsOptimized(val interface{}) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetEndDate(val *string) {
	if err := j.validateSetEndDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endDate",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetEndDateType(val *string) {
	if err := j.validateSetEndDateTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endDateType",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetEphemeralStorage(val interface{}) {
	if err := j.validateSetEphemeralStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ephemeralStorage",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetInstanceCount(val *float64) {
	if err := j.validateSetInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetInstanceMatchCriteria(val *string) {
	if err := j.validateSetInstanceMatchCriteriaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceMatchCriteria",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetInstancePlatform(val *string) {
	if err := j.validateSetInstancePlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instancePlatform",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetOutpostArn(val *string) {
	if err := j.validateSetOutpostArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outpostArn",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetPlacementGroupArn(val *string) {
	if err := j.validateSetPlacementGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementGroupArn",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservation)SetTenancy(val *string) {
	if err := j.validateSetTenancyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenancy",
		val,
	)
}

// Generates CDKTF code for importing a Ec2CapacityReservation resource upon running "cdktf plan <stack-name>".
func Ec2CapacityReservation_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEc2CapacityReservation_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ec2CapacityReservation.Ec2CapacityReservation",
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
func Ec2CapacityReservation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2CapacityReservation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ec2CapacityReservation.Ec2CapacityReservation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2CapacityReservation_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2CapacityReservation_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ec2CapacityReservation.Ec2CapacityReservation",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2CapacityReservation_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2CapacityReservation_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ec2CapacityReservation.Ec2CapacityReservation",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Ec2CapacityReservation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.ec2CapacityReservation.Ec2CapacityReservation",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Ec2CapacityReservation) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2CapacityReservation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2CapacityReservation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2CapacityReservation) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2CapacityReservation) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2CapacityReservation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2CapacityReservation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2CapacityReservation) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2CapacityReservation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2CapacityReservation) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2CapacityReservation) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2CapacityReservation) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) PutTimeouts(value *Ec2CapacityReservationTimeouts) {
	if err := e.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		e,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) ResetEndDate() {
	_jsii_.InvokeVoid(
		e,
		"resetEndDate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) ResetEndDateType() {
	_jsii_.InvokeVoid(
		e,
		"resetEndDateType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) ResetEphemeralStorage() {
	_jsii_.InvokeVoid(
		e,
		"resetEphemeralStorage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) ResetInstanceMatchCriteria() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceMatchCriteria",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) ResetOutpostArn() {
	_jsii_.InvokeVoid(
		e,
		"resetOutpostArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) ResetPlacementGroupArn() {
	_jsii_.InvokeVoid(
		e,
		"resetPlacementGroupArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) ResetTagsAll() {
	_jsii_.InvokeVoid(
		e,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) ResetTenancy() {
	_jsii_.InvokeVoid(
		e,
		"resetTenancy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) ResetTimeouts() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2CapacityReservation) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2CapacityReservation) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2CapacityReservation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2CapacityReservation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2CapacityReservation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

