// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ecsservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcsServiceDeploymentConfigurationLifecycleHookOutputReference interface {
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
	HookDetails() *string
	SetHookDetails(val *string)
	HookDetailsInput() *string
	HookTargetArn() *string
	SetHookTargetArn(val *string)
	HookTargetArnInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LifecycleStages() *[]*string
	SetLifecycleStages(val *[]*string)
	LifecycleStagesInput() *[]*string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
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
	ResetHookDetails()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsServiceDeploymentConfigurationLifecycleHookOutputReference
type jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) HookDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hookDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) HookDetailsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hookDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) HookTargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hookTargetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) HookTargetArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hookTargetArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) LifecycleStages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleStages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) LifecycleStagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleStagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEcsServiceDeploymentConfigurationLifecycleHookOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EcsServiceDeploymentConfigurationLifecycleHookOutputReference {
	_init_.Initialize()

	if err := validateNewEcsServiceDeploymentConfigurationLifecycleHookOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ecsService.EcsServiceDeploymentConfigurationLifecycleHookOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEcsServiceDeploymentConfigurationLifecycleHookOutputReference_Override(e EcsServiceDeploymentConfigurationLifecycleHookOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ecsService.EcsServiceDeploymentConfigurationLifecycleHookOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference)SetHookDetails(val *string) {
	if err := j.validateSetHookDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hookDetails",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference)SetHookTargetArn(val *string) {
	if err := j.validateSetHookTargetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hookTargetArn",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference)SetLifecycleStages(val *[]*string) {
	if err := j.validateSetLifecycleStagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleStages",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) ResetHookDetails() {
	_jsii_.InvokeVoid(
		e,
		"resetHookDetails",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationLifecycleHookOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

