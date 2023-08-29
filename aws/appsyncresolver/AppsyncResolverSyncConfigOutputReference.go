// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncresolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/appsyncresolver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppsyncResolverSyncConfigOutputReference interface {
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
	ConflictDetection() *string
	SetConflictDetection(val *string)
	ConflictDetectionInput() *string
	ConflictHandler() *string
	SetConflictHandler(val *string)
	ConflictHandlerInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *AppsyncResolverSyncConfig
	SetInternalValue(val *AppsyncResolverSyncConfig)
	LambdaConflictHandlerConfig() AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference
	LambdaConflictHandlerConfigInput() *AppsyncResolverSyncConfigLambdaConflictHandlerConfig
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
	PutLambdaConflictHandlerConfig(value *AppsyncResolverSyncConfigLambdaConflictHandlerConfig)
	ResetConflictDetection()
	ResetConflictHandler()
	ResetLambdaConflictHandlerConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppsyncResolverSyncConfigOutputReference
type jsiiProxy_AppsyncResolverSyncConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) ConflictDetection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conflictDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) ConflictDetectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conflictDetectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) ConflictHandler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conflictHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) ConflictHandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conflictHandlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) InternalValue() *AppsyncResolverSyncConfig {
	var returns *AppsyncResolverSyncConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) LambdaConflictHandlerConfig() AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference {
	var returns AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference
	_jsii_.Get(
		j,
		"lambdaConflictHandlerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) LambdaConflictHandlerConfigInput() *AppsyncResolverSyncConfigLambdaConflictHandlerConfig {
	var returns *AppsyncResolverSyncConfigLambdaConflictHandlerConfig
	_jsii_.Get(
		j,
		"lambdaConflictHandlerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppsyncResolverSyncConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppsyncResolverSyncConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAppsyncResolverSyncConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsyncResolverSyncConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appsyncResolver.AppsyncResolverSyncConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppsyncResolverSyncConfigOutputReference_Override(a AppsyncResolverSyncConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appsyncResolver.AppsyncResolverSyncConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference)SetConflictDetection(val *string) {
	if err := j.validateSetConflictDetectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conflictDetection",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference)SetConflictHandler(val *string) {
	if err := j.validateSetConflictHandlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conflictHandler",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference)SetInternalValue(val *AppsyncResolverSyncConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) PutLambdaConflictHandlerConfig(value *AppsyncResolverSyncConfigLambdaConflictHandlerConfig) {
	if err := a.validatePutLambdaConflictHandlerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaConflictHandlerConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) ResetConflictDetection() {
	_jsii_.InvokeVoid(
		a,
		"resetConflictDetection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) ResetConflictHandler() {
	_jsii_.InvokeVoid(
		a,
		"resetConflictHandler",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) ResetLambdaConflictHandlerConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaConflictHandlerConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

