// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticachereplicationgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/elasticachereplicationgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElasticacheReplicationGroupNodeGroupConfigurationOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NodeGroupId() *string
	SetNodeGroupId(val *string)
	NodeGroupIdInput() *string
	PrimaryAvailabilityZone() *string
	SetPrimaryAvailabilityZone(val *string)
	PrimaryAvailabilityZoneInput() *string
	PrimaryOutpostArn() *string
	SetPrimaryOutpostArn(val *string)
	PrimaryOutpostArnInput() *string
	ReplicaAvailabilityZones() *[]*string
	SetReplicaAvailabilityZones(val *[]*string)
	ReplicaAvailabilityZonesInput() *[]*string
	ReplicaCount() *float64
	SetReplicaCount(val *float64)
	ReplicaCountInput() *float64
	ReplicaOutpostArns() *[]*string
	SetReplicaOutpostArns(val *[]*string)
	ReplicaOutpostArnsInput() *[]*string
	Slots() *string
	SetSlots(val *string)
	SlotsInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetNodeGroupId()
	ResetPrimaryAvailabilityZone()
	ResetPrimaryOutpostArn()
	ResetReplicaAvailabilityZones()
	ResetReplicaCount()
	ResetReplicaOutpostArns()
	ResetSlots()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElasticacheReplicationGroupNodeGroupConfigurationOutputReference
type jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) NodeGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) NodeGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) PrimaryAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) PrimaryAvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryAvailabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) PrimaryOutpostArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryOutpostArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) PrimaryOutpostArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryOutpostArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) ReplicaAvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicaAvailabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) ReplicaAvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicaAvailabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) ReplicaCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) ReplicaCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) ReplicaOutpostArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicaOutpostArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) ReplicaOutpostArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicaOutpostArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) Slots() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) SlotsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElasticacheReplicationGroupNodeGroupConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ElasticacheReplicationGroupNodeGroupConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewElasticacheReplicationGroupNodeGroupConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.elasticacheReplicationGroup.ElasticacheReplicationGroupNodeGroupConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewElasticacheReplicationGroupNodeGroupConfigurationOutputReference_Override(e ElasticacheReplicationGroupNodeGroupConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.elasticacheReplicationGroup.ElasticacheReplicationGroupNodeGroupConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference)SetNodeGroupId(val *string) {
	if err := j.validateSetNodeGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeGroupId",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference)SetPrimaryAvailabilityZone(val *string) {
	if err := j.validateSetPrimaryAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryAvailabilityZone",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference)SetPrimaryOutpostArn(val *string) {
	if err := j.validateSetPrimaryOutpostArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryOutpostArn",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference)SetReplicaAvailabilityZones(val *[]*string) {
	if err := j.validateSetReplicaAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaAvailabilityZones",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference)SetReplicaCount(val *float64) {
	if err := j.validateSetReplicaCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaCount",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference)SetReplicaOutpostArns(val *[]*string) {
	if err := j.validateSetReplicaOutpostArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaOutpostArns",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference)SetSlots(val *string) {
	if err := j.validateSetSlotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slots",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) ResetNodeGroupId() {
	_jsii_.InvokeVoid(
		e,
		"resetNodeGroupId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) ResetPrimaryAvailabilityZone() {
	_jsii_.InvokeVoid(
		e,
		"resetPrimaryAvailabilityZone",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) ResetPrimaryOutpostArn() {
	_jsii_.InvokeVoid(
		e,
		"resetPrimaryOutpostArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) ResetReplicaAvailabilityZones() {
	_jsii_.InvokeVoid(
		e,
		"resetReplicaAvailabilityZones",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) ResetReplicaCount() {
	_jsii_.InvokeVoid(
		e,
		"resetReplicaCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) ResetReplicaOutpostArns() {
	_jsii_.InvokeVoid(
		e,
		"resetReplicaOutpostArns",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) ResetSlots() {
	_jsii_.InvokeVoid(
		e,
		"resetSlots",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheReplicationGroupNodeGroupConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

