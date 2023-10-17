// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gameliftfleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/gameliftfleet/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GameliftFleetRuntimeConfigurationOutputReference interface {
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
	GameSessionActivationTimeoutSeconds() *float64
	SetGameSessionActivationTimeoutSeconds(val *float64)
	GameSessionActivationTimeoutSecondsInput() *float64
	InternalValue() *GameliftFleetRuntimeConfiguration
	SetInternalValue(val *GameliftFleetRuntimeConfiguration)
	MaxConcurrentGameSessionActivations() *float64
	SetMaxConcurrentGameSessionActivations(val *float64)
	MaxConcurrentGameSessionActivationsInput() *float64
	ServerProcess() GameliftFleetRuntimeConfigurationServerProcessList
	ServerProcessInput() interface{}
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
	PutServerProcess(value interface{})
	ResetGameSessionActivationTimeoutSeconds()
	ResetMaxConcurrentGameSessionActivations()
	ResetServerProcess()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GameliftFleetRuntimeConfigurationOutputReference
type jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) GameSessionActivationTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gameSessionActivationTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) GameSessionActivationTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gameSessionActivationTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) InternalValue() *GameliftFleetRuntimeConfiguration {
	var returns *GameliftFleetRuntimeConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) MaxConcurrentGameSessionActivations() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentGameSessionActivations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) MaxConcurrentGameSessionActivationsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentGameSessionActivationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) ServerProcess() GameliftFleetRuntimeConfigurationServerProcessList {
	var returns GameliftFleetRuntimeConfigurationServerProcessList
	_jsii_.Get(
		j,
		"serverProcess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) ServerProcessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverProcessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGameliftFleetRuntimeConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GameliftFleetRuntimeConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewGameliftFleetRuntimeConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.gameliftFleet.GameliftFleetRuntimeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGameliftFleetRuntimeConfigurationOutputReference_Override(g GameliftFleetRuntimeConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.gameliftFleet.GameliftFleetRuntimeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference)SetGameSessionActivationTimeoutSeconds(val *float64) {
	if err := j.validateSetGameSessionActivationTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gameSessionActivationTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference)SetInternalValue(val *GameliftFleetRuntimeConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference)SetMaxConcurrentGameSessionActivations(val *float64) {
	if err := j.validateSetMaxConcurrentGameSessionActivationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentGameSessionActivations",
		val,
	)
}

func (j *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) PutServerProcess(value interface{}) {
	if err := g.validatePutServerProcessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServerProcess",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) ResetGameSessionActivationTimeoutSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetGameSessionActivationTimeoutSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) ResetMaxConcurrentGameSessionActivations() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxConcurrentGameSessionActivations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) ResetServerProcess() {
	_jsii_.InvokeVoid(
		g,
		"resetServerProcess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleetRuntimeConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

