// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmcontactsrotation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ssmcontactsrotation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsmcontactsRotationRecurrenceOutputReference interface {
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
	DailySettings() SsmcontactsRotationRecurrenceDailySettingsList
	DailySettingsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MonthlySettings() SsmcontactsRotationRecurrenceMonthlySettingsList
	MonthlySettingsInput() interface{}
	NumberOfOnCalls() *float64
	SetNumberOfOnCalls(val *float64)
	NumberOfOnCallsInput() *float64
	RecurrenceMultiplier() *float64
	SetRecurrenceMultiplier(val *float64)
	RecurrenceMultiplierInput() *float64
	ShiftCoverages() SsmcontactsRotationRecurrenceShiftCoveragesList
	ShiftCoveragesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeeklySettings() SsmcontactsRotationRecurrenceWeeklySettingsList
	WeeklySettingsInput() interface{}
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
	PutDailySettings(value interface{})
	PutMonthlySettings(value interface{})
	PutShiftCoverages(value interface{})
	PutWeeklySettings(value interface{})
	ResetDailySettings()
	ResetMonthlySettings()
	ResetShiftCoverages()
	ResetWeeklySettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SsmcontactsRotationRecurrenceOutputReference
type jsiiProxy_SsmcontactsRotationRecurrenceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) DailySettings() SsmcontactsRotationRecurrenceDailySettingsList {
	var returns SsmcontactsRotationRecurrenceDailySettingsList
	_jsii_.Get(
		j,
		"dailySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) DailySettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dailySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) MonthlySettings() SsmcontactsRotationRecurrenceMonthlySettingsList {
	var returns SsmcontactsRotationRecurrenceMonthlySettingsList
	_jsii_.Get(
		j,
		"monthlySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) MonthlySettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monthlySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) NumberOfOnCalls() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfOnCalls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) NumberOfOnCallsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfOnCallsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) RecurrenceMultiplier() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recurrenceMultiplier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) RecurrenceMultiplierInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recurrenceMultiplierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) ShiftCoverages() SsmcontactsRotationRecurrenceShiftCoveragesList {
	var returns SsmcontactsRotationRecurrenceShiftCoveragesList
	_jsii_.Get(
		j,
		"shiftCoverages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) ShiftCoveragesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shiftCoveragesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) WeeklySettings() SsmcontactsRotationRecurrenceWeeklySettingsList {
	var returns SsmcontactsRotationRecurrenceWeeklySettingsList
	_jsii_.Get(
		j,
		"weeklySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) WeeklySettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"weeklySettingsInput",
		&returns,
	)
	return returns
}


func NewSsmcontactsRotationRecurrenceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SsmcontactsRotationRecurrenceOutputReference {
	_init_.Initialize()

	if err := validateNewSsmcontactsRotationRecurrenceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SsmcontactsRotationRecurrenceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ssmcontactsRotation.SsmcontactsRotationRecurrenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSsmcontactsRotationRecurrenceOutputReference_Override(s SsmcontactsRotationRecurrenceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ssmcontactsRotation.SsmcontactsRotationRecurrenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference)SetNumberOfOnCalls(val *float64) {
	if err := j.validateSetNumberOfOnCallsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfOnCalls",
		val,
	)
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference)SetRecurrenceMultiplier(val *float64) {
	if err := j.validateSetRecurrenceMultiplierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recurrenceMultiplier",
		val,
	)
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) PutDailySettings(value interface{}) {
	if err := s.validatePutDailySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDailySettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) PutMonthlySettings(value interface{}) {
	if err := s.validatePutMonthlySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMonthlySettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) PutShiftCoverages(value interface{}) {
	if err := s.validatePutShiftCoveragesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putShiftCoverages",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) PutWeeklySettings(value interface{}) {
	if err := s.validatePutWeeklySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putWeeklySettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) ResetDailySettings() {
	_jsii_.InvokeVoid(
		s,
		"resetDailySettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) ResetMonthlySettings() {
	_jsii_.InvokeVoid(
		s,
		"resetMonthlySettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) ResetShiftCoverages() {
	_jsii_.InvokeVoid(
		s,
		"resetShiftCoverages",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) ResetWeeklySettings() {
	_jsii_.InvokeVoid(
		s,
		"resetWeeklySettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmcontactsRotationRecurrenceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

