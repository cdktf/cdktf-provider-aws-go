// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrcontainersjobtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/emrcontainersjobtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference interface {
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
	InternalValue() *EmrcontainersJobTemplateJobTemplateDataJobDriver
	SetInternalValue(val *EmrcontainersJobTemplateJobTemplateDataJobDriver)
	SparkSqlJobDriver() EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference
	SparkSqlJobDriverInput() *EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriver
	SparkSubmitJobDriver() EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSubmitJobDriverOutputReference
	SparkSubmitJobDriverInput() *EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSubmitJobDriver
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
	PutSparkSqlJobDriver(value *EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriver)
	PutSparkSubmitJobDriver(value *EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSubmitJobDriver)
	ResetSparkSqlJobDriver()
	ResetSparkSubmitJobDriver()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference
type jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) InternalValue() *EmrcontainersJobTemplateJobTemplateDataJobDriver {
	var returns *EmrcontainersJobTemplateJobTemplateDataJobDriver
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) SparkSqlJobDriver() EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference {
	var returns EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriverOutputReference
	_jsii_.Get(
		j,
		"sparkSqlJobDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) SparkSqlJobDriverInput() *EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriver {
	var returns *EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriver
	_jsii_.Get(
		j,
		"sparkSqlJobDriverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) SparkSubmitJobDriver() EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSubmitJobDriverOutputReference {
	var returns EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSubmitJobDriverOutputReference
	_jsii_.Get(
		j,
		"sparkSubmitJobDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) SparkSubmitJobDriverInput() *EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSubmitJobDriver {
	var returns *EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSubmitJobDriver
	_jsii_.Get(
		j,
		"sparkSubmitJobDriverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference {
	_init_.Initialize()

	if err := validateNewEmrcontainersJobTemplateJobTemplateDataJobDriverOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.emrcontainersJobTemplate.EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference_Override(e EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.emrcontainersJobTemplate.EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference)SetInternalValue(val *EmrcontainersJobTemplateJobTemplateDataJobDriver) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) PutSparkSqlJobDriver(value *EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriver) {
	if err := e.validatePutSparkSqlJobDriverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSparkSqlJobDriver",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) PutSparkSubmitJobDriver(value *EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSubmitJobDriver) {
	if err := e.validatePutSparkSubmitJobDriverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSparkSubmitJobDriver",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) ResetSparkSqlJobDriver() {
	_jsii_.InvokeVoid(
		e,
		"resetSparkSqlJobDriver",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) ResetSparkSubmitJobDriver() {
	_jsii_.InvokeVoid(
		e,
		"resetSparkSubmitJobDriver",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EmrcontainersJobTemplateJobTemplateDataJobDriverOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

