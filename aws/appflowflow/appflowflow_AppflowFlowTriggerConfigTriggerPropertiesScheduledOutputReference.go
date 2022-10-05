package appflowflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/appflowflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference interface {
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
	DataPullMode() *string
	SetDataPullMode(val *string)
	DataPullModeInput() *string
	FirstExecutionFrom() *string
	SetFirstExecutionFrom(val *string)
	FirstExecutionFromInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *AppflowFlowTriggerConfigTriggerPropertiesScheduled
	SetInternalValue(val *AppflowFlowTriggerConfigTriggerPropertiesScheduled)
	ScheduleEndTime() *string
	SetScheduleEndTime(val *string)
	ScheduleEndTimeInput() *string
	ScheduleExpression() *string
	SetScheduleExpression(val *string)
	ScheduleExpressionInput() *string
	ScheduleOffset() *float64
	SetScheduleOffset(val *float64)
	ScheduleOffsetInput() *float64
	ScheduleStartTime() *string
	SetScheduleStartTime(val *string)
	ScheduleStartTimeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDataPullMode()
	ResetFirstExecutionFrom()
	ResetScheduleEndTime()
	ResetScheduleOffset()
	ResetScheduleStartTime()
	ResetTimezone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference
type jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) DataPullMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPullMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) DataPullModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPullModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) FirstExecutionFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstExecutionFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) FirstExecutionFromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstExecutionFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) InternalValue() *AppflowFlowTriggerConfigTriggerPropertiesScheduled {
	var returns *AppflowFlowTriggerConfigTriggerPropertiesScheduled
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) ScheduleEndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) ScheduleEndTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleEndTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) ScheduleExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) ScheduleExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) ScheduleOffset() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduleOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) ScheduleOffsetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduleOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) ScheduleStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) ScheduleStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}


func NewAppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference {
	_init_.Initialize()

	if err := validateNewAppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference_Override(a AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference)SetDataPullMode(val *string) {
	if err := j.validateSetDataPullModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataPullMode",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference)SetFirstExecutionFrom(val *string) {
	if err := j.validateSetFirstExecutionFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstExecutionFrom",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference)SetInternalValue(val *AppflowFlowTriggerConfigTriggerPropertiesScheduled) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference)SetScheduleEndTime(val *string) {
	if err := j.validateSetScheduleEndTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleEndTime",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference)SetScheduleExpression(val *string) {
	if err := j.validateSetScheduleExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleExpression",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference)SetScheduleOffset(val *float64) {
	if err := j.validateSetScheduleOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleOffset",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference)SetScheduleStartTime(val *string) {
	if err := j.validateSetScheduleStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleStartTime",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) ResetDataPullMode() {
	_jsii_.InvokeVoid(
		a,
		"resetDataPullMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) ResetFirstExecutionFrom() {
	_jsii_.InvokeVoid(
		a,
		"resetFirstExecutionFrom",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) ResetScheduleEndTime() {
	_jsii_.InvokeVoid(
		a,
		"resetScheduleEndTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) ResetScheduleOffset() {
	_jsii_.InvokeVoid(
		a,
		"resetScheduleOffset",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) ResetScheduleStartTime() {
	_jsii_.InvokeVoid(
		a,
		"resetScheduleStartTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) ResetTimezone() {
	_jsii_.InvokeVoid(
		a,
		"resetTimezone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesScheduledOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

