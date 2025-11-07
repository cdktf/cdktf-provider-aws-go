// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package odbcloudexadatainfrastructure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/odbcloudexadatainfrastructure/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OdbCloudExadataInfrastructureMaintenanceWindowOutputReference interface {
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
	CustomActionTimeoutInMins() *float64
	SetCustomActionTimeoutInMins(val *float64)
	CustomActionTimeoutInMinsInput() *float64
	DaysOfWeek() OdbCloudExadataInfrastructureMaintenanceWindowDaysOfWeekList
	DaysOfWeekInput() interface{}
	// Experimental.
	Fqn() *string
	HoursOfDay() *[]*float64
	SetHoursOfDay(val *[]*float64)
	HoursOfDayInput() *[]*float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsCustomActionTimeoutEnabled() interface{}
	SetIsCustomActionTimeoutEnabled(val interface{})
	IsCustomActionTimeoutEnabledInput() interface{}
	LeadTimeInWeeks() *float64
	SetLeadTimeInWeeks(val *float64)
	LeadTimeInWeeksInput() *float64
	Months() OdbCloudExadataInfrastructureMaintenanceWindowMonthsList
	MonthsInput() interface{}
	PatchingMode() *string
	SetPatchingMode(val *string)
	PatchingModeInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutDaysOfWeek(value interface{})
	PutMonths(value interface{})
	ResetDaysOfWeek()
	ResetHoursOfDay()
	ResetLeadTimeInWeeks()
	ResetMonths()
	ResetWeeksOfMonth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OdbCloudExadataInfrastructureMaintenanceWindowOutputReference
type jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) CustomActionTimeoutInMins() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customActionTimeoutInMins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) CustomActionTimeoutInMinsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customActionTimeoutInMinsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) DaysOfWeek() OdbCloudExadataInfrastructureMaintenanceWindowDaysOfWeekList {
	var returns OdbCloudExadataInfrastructureMaintenanceWindowDaysOfWeekList
	_jsii_.Get(
		j,
		"daysOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) DaysOfWeekInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"daysOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) HoursOfDay() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"hoursOfDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) HoursOfDayInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"hoursOfDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) IsCustomActionTimeoutEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCustomActionTimeoutEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) IsCustomActionTimeoutEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCustomActionTimeoutEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) LeadTimeInWeeks() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leadTimeInWeeks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) LeadTimeInWeeksInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leadTimeInWeeksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) Months() OdbCloudExadataInfrastructureMaintenanceWindowMonthsList {
	var returns OdbCloudExadataInfrastructureMaintenanceWindowMonthsList
	_jsii_.Get(
		j,
		"months",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) MonthsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monthsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) PatchingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) PatchingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) Preference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) PreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) WeeksOfMonth() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"weeksOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) WeeksOfMonthInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"weeksOfMonthInput",
		&returns,
	)
	return returns
}


func NewOdbCloudExadataInfrastructureMaintenanceWindowOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OdbCloudExadataInfrastructureMaintenanceWindowOutputReference {
	_init_.Initialize()

	if err := validateNewOdbCloudExadataInfrastructureMaintenanceWindowOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.odbCloudExadataInfrastructure.OdbCloudExadataInfrastructureMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOdbCloudExadataInfrastructureMaintenanceWindowOutputReference_Override(o OdbCloudExadataInfrastructureMaintenanceWindowOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.odbCloudExadataInfrastructure.OdbCloudExadataInfrastructureMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference)SetCustomActionTimeoutInMins(val *float64) {
	if err := j.validateSetCustomActionTimeoutInMinsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customActionTimeoutInMins",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference)SetHoursOfDay(val *[]*float64) {
	if err := j.validateSetHoursOfDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hoursOfDay",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference)SetIsCustomActionTimeoutEnabled(val interface{}) {
	if err := j.validateSetIsCustomActionTimeoutEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCustomActionTimeoutEnabled",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference)SetLeadTimeInWeeks(val *float64) {
	if err := j.validateSetLeadTimeInWeeksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"leadTimeInWeeks",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference)SetPatchingMode(val *string) {
	if err := j.validateSetPatchingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patchingMode",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference)SetPreference(val *string) {
	if err := j.validateSetPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preference",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference)SetWeeksOfMonth(val *[]*float64) {
	if err := j.validateSetWeeksOfMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeksOfMonth",
		val,
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) PutDaysOfWeek(value interface{}) {
	if err := o.validatePutDaysOfWeekParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDaysOfWeek",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) PutMonths(value interface{}) {
	if err := o.validatePutMonthsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMonths",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) ResetDaysOfWeek() {
	_jsii_.InvokeVoid(
		o,
		"resetDaysOfWeek",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) ResetHoursOfDay() {
	_jsii_.InvokeVoid(
		o,
		"resetHoursOfDay",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) ResetLeadTimeInWeeks() {
	_jsii_.InvokeVoid(
		o,
		"resetLeadTimeInWeeks",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) ResetMonths() {
	_jsii_.InvokeVoid(
		o,
		"resetMonths",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) ResetWeeksOfMonth() {
	_jsii_.InvokeVoid(
		o,
		"resetWeeksOfMonth",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudExadataInfrastructureMaintenanceWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

