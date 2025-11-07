// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightrefreshschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/quicksightrefreshschedule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference interface {
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
	Interval() *string
	SetInterval(val *string)
	IntervalInput() *string
	RefreshOnDay() QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList
	RefreshOnDayInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeOfTheDay() *string
	SetTimeOfTheDay(val *string)
	TimeOfTheDayInput() *string
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
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
	PutRefreshOnDay(value interface{})
	ResetRefreshOnDay()
	ResetTimeOfTheDay()
	ResetTimezone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference
type jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) Interval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) IntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) RefreshOnDay() QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList {
	var returns QuicksightRefreshScheduleScheduleScheduleFrequencyRefreshOnDayList
	_jsii_.Get(
		j,
		"refreshOnDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) RefreshOnDayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"refreshOnDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) TimeOfTheDay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfTheDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) TimeOfTheDayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfTheDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}


func NewQuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightRefreshScheduleScheduleScheduleFrequencyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightRefreshSchedule.QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewQuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference_Override(q QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.quicksightRefreshSchedule.QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		q,
	)
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference)SetInterval(val *string) {
	if err := j.validateSetIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference)SetTimeOfTheDay(val *string) {
	if err := j.validateSetTimeOfTheDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeOfTheDay",
		val,
	)
}

func (j *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) PutRefreshOnDay(value interface{}) {
	if err := q.validatePutRefreshOnDayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putRefreshOnDay",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) ResetRefreshOnDay() {
	_jsii_.InvokeVoid(
		q,
		"resetRefreshOnDay",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) ResetTimeOfTheDay() {
	_jsii_.InvokeVoid(
		q,
		"resetTimeOfTheDay",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) ResetTimezone() {
	_jsii_.InvokeVoid(
		q,
		"resetTimezone",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := q.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightRefreshScheduleScheduleScheduleFrequencyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

