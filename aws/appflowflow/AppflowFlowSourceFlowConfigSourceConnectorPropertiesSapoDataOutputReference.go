// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/appflowflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference interface {
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
	InternalValue() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoData
	SetInternalValue(val *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoData)
	ObjectPath() *string
	SetObjectPath(val *string)
	ObjectPathInput() *string
	PaginationConfig() AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataPaginationConfigOutputReference
	PaginationConfigInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataPaginationConfig
	ParallelismConfig() AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataParallelismConfigOutputReference
	ParallelismConfigInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataParallelismConfig
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
	PutPaginationConfig(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataPaginationConfig)
	PutParallelismConfig(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataParallelismConfig)
	ResetPaginationConfig()
	ResetParallelismConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference
type jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) InternalValue() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoData {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoData
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) ObjectPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) ObjectPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) PaginationConfig() AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataPaginationConfigOutputReference {
	var returns AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataPaginationConfigOutputReference
	_jsii_.Get(
		j,
		"paginationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) PaginationConfigInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataPaginationConfig {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataPaginationConfig
	_jsii_.Get(
		j,
		"paginationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) ParallelismConfig() AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataParallelismConfigOutputReference {
	var returns AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataParallelismConfigOutputReference
	_jsii_.Get(
		j,
		"parallelismConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) ParallelismConfigInput() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataParallelismConfig {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataParallelismConfig
	_jsii_.Get(
		j,
		"parallelismConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference {
	_init_.Initialize()

	if err := validateNewAppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference_Override(a AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference)SetInternalValue(val *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoData) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference)SetObjectPath(val *string) {
	if err := j.validateSetObjectPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectPath",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) PutPaginationConfig(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataPaginationConfig) {
	if err := a.validatePutPaginationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPaginationConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) PutParallelismConfig(value *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataParallelismConfig) {
	if err := a.validatePutParallelismConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParallelismConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) ResetPaginationConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetPaginationConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) ResetParallelismConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetParallelismConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoDataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

