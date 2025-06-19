// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrcontainersjobtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/emrcontainersjobtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmrcontainersJobTemplateJobTemplateDataOutputReference interface {
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
	ConfigurationOverrides() EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference
	ConfigurationOverridesInput() *EmrcontainersJobTemplateJobTemplateDataConfigurationOverrides
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	ExecutionRoleArnInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *EmrcontainersJobTemplateJobTemplateData
	SetInternalValue(val *EmrcontainersJobTemplateJobTemplateData)
	JobDriver() EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference
	JobDriverInput() *EmrcontainersJobTemplateJobTemplateDataJobDriver
	JobTags() *map[string]*string
	SetJobTags(val *map[string]*string)
	JobTagsInput() *map[string]*string
	ReleaseLabel() *string
	SetReleaseLabel(val *string)
	ReleaseLabelInput() *string
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutConfigurationOverrides(value *EmrcontainersJobTemplateJobTemplateDataConfigurationOverrides)
	PutJobDriver(value *EmrcontainersJobTemplateJobTemplateDataJobDriver)
	ResetConfigurationOverrides()
	ResetJobTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EmrcontainersJobTemplateJobTemplateDataOutputReference
type jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) ConfigurationOverrides() EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference {
	var returns EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference
	_jsii_.Get(
		j,
		"configurationOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) ConfigurationOverridesInput() *EmrcontainersJobTemplateJobTemplateDataConfigurationOverrides {
	var returns *EmrcontainersJobTemplateJobTemplateDataConfigurationOverrides
	_jsii_.Get(
		j,
		"configurationOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) InternalValue() *EmrcontainersJobTemplateJobTemplateData {
	var returns *EmrcontainersJobTemplateJobTemplateData
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) JobDriver() EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference {
	var returns EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference
	_jsii_.Get(
		j,
		"jobDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) JobDriverInput() *EmrcontainersJobTemplateJobTemplateDataJobDriver {
	var returns *EmrcontainersJobTemplateJobTemplateDataJobDriver
	_jsii_.Get(
		j,
		"jobDriverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) JobTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"jobTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) JobTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"jobTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) ReleaseLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) ReleaseLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEmrcontainersJobTemplateJobTemplateDataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EmrcontainersJobTemplateJobTemplateDataOutputReference {
	_init_.Initialize()

	if err := validateNewEmrcontainersJobTemplateJobTemplateDataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.emrcontainersJobTemplate.EmrcontainersJobTemplateJobTemplateDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEmrcontainersJobTemplateJobTemplateDataOutputReference_Override(e EmrcontainersJobTemplateJobTemplateDataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.emrcontainersJobTemplate.EmrcontainersJobTemplateJobTemplateDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference)SetInternalValue(val *EmrcontainersJobTemplateJobTemplateData) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference)SetJobTags(val *map[string]*string) {
	if err := j.validateSetJobTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobTags",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference)SetReleaseLabel(val *string) {
	if err := j.validateSetReleaseLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseLabel",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) PutConfigurationOverrides(value *EmrcontainersJobTemplateJobTemplateDataConfigurationOverrides) {
	if err := e.validatePutConfigurationOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putConfigurationOverrides",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) PutJobDriver(value *EmrcontainersJobTemplateJobTemplateDataJobDriver) {
	if err := e.validatePutJobDriverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putJobDriver",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) ResetConfigurationOverrides() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigurationOverrides",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) ResetJobTags() {
	_jsii_.InvokeVoid(
		e,
		"resetJobTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

