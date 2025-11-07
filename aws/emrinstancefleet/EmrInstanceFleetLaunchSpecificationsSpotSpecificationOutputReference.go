// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrinstancefleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/emrinstancefleet/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference interface {
	cdktf.ComplexObject
	AllocationStrategy() *string
	SetAllocationStrategy(val *string)
	AllocationStrategyInput() *string
	BlockDurationMinutes() *float64
	SetBlockDurationMinutes(val *float64)
	BlockDurationMinutesInput() *float64
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
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutAction() *string
	SetTimeoutAction(val *string)
	TimeoutActionInput() *string
	TimeoutDurationMinutes() *float64
	SetTimeoutDurationMinutes(val *float64)
	TimeoutDurationMinutesInput() *float64
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
	ResetBlockDurationMinutes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference
type jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) AllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) AllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) BlockDurationMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockDurationMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) BlockDurationMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockDurationMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) TimeoutAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) TimeoutActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) TimeoutDurationMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutDurationMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) TimeoutDurationMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutDurationMinutesInput",
		&returns,
	)
	return returns
}


func NewEmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewEmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.emrInstanceFleet.EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference_Override(e EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.emrInstanceFleet.EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference)SetAllocationStrategy(val *string) {
	if err := j.validateSetAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationStrategy",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference)SetBlockDurationMinutes(val *float64) {
	if err := j.validateSetBlockDurationMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockDurationMinutes",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference)SetTimeoutAction(val *string) {
	if err := j.validateSetTimeoutActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutAction",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference)SetTimeoutDurationMinutes(val *float64) {
	if err := j.validateSetTimeoutDurationMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutDurationMinutes",
		val,
	)
}

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) ResetBlockDurationMinutes() {
	_jsii_.InvokeVoid(
		e,
		"resetBlockDurationMinutes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsSpotSpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

