// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrcontainersjobtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/emrcontainersjobtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference interface {
	cdktf.ComplexObject
	ApplicationConfiguration() EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesApplicationConfigurationList
	ApplicationConfigurationInput() interface{}
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
	InternalValue() *EmrcontainersJobTemplateJobTemplateDataConfigurationOverrides
	SetInternalValue(val *EmrcontainersJobTemplateJobTemplateDataConfigurationOverrides)
	MonitoringConfiguration() EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationOutputReference
	MonitoringConfigurationInput() *EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesMonitoringConfiguration
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
	PutApplicationConfiguration(value interface{})
	PutMonitoringConfiguration(value *EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesMonitoringConfiguration)
	ResetApplicationConfiguration()
	ResetMonitoringConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference
type jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) ApplicationConfiguration() EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesApplicationConfigurationList {
	var returns EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesApplicationConfigurationList
	_jsii_.Get(
		j,
		"applicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) ApplicationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) InternalValue() *EmrcontainersJobTemplateJobTemplateDataConfigurationOverrides {
	var returns *EmrcontainersJobTemplateJobTemplateDataConfigurationOverrides
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) MonitoringConfiguration() EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationOutputReference {
	var returns EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationOutputReference
	_jsii_.Get(
		j,
		"monitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) MonitoringConfigurationInput() *EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesMonitoringConfiguration {
	var returns *EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesMonitoringConfiguration
	_jsii_.Get(
		j,
		"monitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference {
	_init_.Initialize()

	if err := validateNewEmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.emrcontainersJobTemplate.EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference_Override(e EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.emrcontainersJobTemplate.EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference)SetInternalValue(val *EmrcontainersJobTemplateJobTemplateDataConfigurationOverrides) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) PutApplicationConfiguration(value interface{}) {
	if err := e.validatePutApplicationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putApplicationConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) PutMonitoringConfiguration(value *EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesMonitoringConfiguration) {
	if err := e.validatePutMonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMonitoringConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) ResetApplicationConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetApplicationConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) ResetMonitoringConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetMonitoringConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataConfigurationOverridesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

