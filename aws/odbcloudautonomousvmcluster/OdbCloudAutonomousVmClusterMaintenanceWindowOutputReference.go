// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package odbcloudautonomousvmcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/odbcloudautonomousvmcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference interface {
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
	DaysOfWeek() OdbCloudAutonomousVmClusterMaintenanceWindowDaysOfWeekList
	DaysOfWeekInput() interface{}
	// Experimental.
	Fqn() *string
	HoursOfDay() *[]*float64
	SetHoursOfDay(val *[]*float64)
	HoursOfDayInput() *[]*float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LeadTimeInWeeks() *float64
	SetLeadTimeInWeeks(val *float64)
	LeadTimeInWeeksInput() *float64
	Months() OdbCloudAutonomousVmClusterMaintenanceWindowMonthsList
	MonthsInput() interface{}
	Preference() *string
	SetPreference(val *string)
	PreferenceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeeksOfMonth() *[]*float64
	SetWeeksOfMonth(val *[]*float64)
	WeeksOfMonthInput() *[]*float64
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutDaysOfWeek(value interface{})
	PutMonths(value interface{})
	ResetDaysOfWeek()
	ResetHoursOfDay()
	ResetLeadTimeInWeeks()
	ResetMonths()
	ResetWeeksOfMonth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference
type jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) DaysOfWeek() OdbCloudAutonomousVmClusterMaintenanceWindowDaysOfWeekList {
	var returns OdbCloudAutonomousVmClusterMaintenanceWindowDaysOfWeekList
	_jsii_.Get(
		j,
		"daysOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) DaysOfWeekInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"daysOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) HoursOfDay() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"hoursOfDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) HoursOfDayInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"hoursOfDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) LeadTimeInWeeks() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leadTimeInWeeks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) LeadTimeInWeeksInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leadTimeInWeeksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) Months() OdbCloudAutonomousVmClusterMaintenanceWindowMonthsList {
	var returns OdbCloudAutonomousVmClusterMaintenanceWindowMonthsList
	_jsii_.Get(
		j,
		"months",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) MonthsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monthsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) Preference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) PreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) WeeksOfMonth() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"weeksOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) WeeksOfMonthInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"weeksOfMonthInput",
		&returns,
	)
	return returns
}


func NewOdbCloudAutonomousVmClusterMaintenanceWindowOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference {
	_init_.Initialize()

	if err := validateNewOdbCloudAutonomousVmClusterMaintenanceWindowOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOdbCloudAutonomousVmClusterMaintenanceWindowOutputReference_Override(o OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference)SetHoursOfDay(val *[]*float64) {
	if err := j.validateSetHoursOfDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hoursOfDay",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference)SetLeadTimeInWeeks(val *float64) {
	if err := j.validateSetLeadTimeInWeeksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"leadTimeInWeeks",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference)SetPreference(val *string) {
	if err := j.validateSetPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preference",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference)SetWeeksOfMonth(val *[]*float64) {
	if err := j.validateSetWeeksOfMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeksOfMonth",
		val,
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) PutDaysOfWeek(value interface{}) {
	if err := o.validatePutDaysOfWeekParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDaysOfWeek",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) PutMonths(value interface{}) {
	if err := o.validatePutMonthsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMonths",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) ResetDaysOfWeek() {
	_jsii_.InvokeVoid(
		o,
		"resetDaysOfWeek",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) ResetHoursOfDay() {
	_jsii_.InvokeVoid(
		o,
		"resetHoursOfDay",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) ResetLeadTimeInWeeks() {
	_jsii_.InvokeVoid(
		o,
		"resetLeadTimeInWeeks",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) ResetMonths() {
	_jsii_.InvokeVoid(
		o,
		"resetMonths",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) ResetWeeksOfMonth() {
	_jsii_.InvokeVoid(
		o,
		"resetWeeksOfMonth",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

