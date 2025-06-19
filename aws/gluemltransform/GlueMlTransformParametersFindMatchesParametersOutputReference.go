// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluemltransform

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/gluemltransform/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlueMlTransformParametersFindMatchesParametersOutputReference interface {
	cdktf.ComplexObject
	AccuracyCostTradeOff() *float64
	SetAccuracyCostTradeOff(val *float64)
	AccuracyCostTradeOffInput() *float64
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
	EnforceProvidedLabels() interface{}
	SetEnforceProvidedLabels(val interface{})
	EnforceProvidedLabelsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GlueMlTransformParametersFindMatchesParameters
	SetInternalValue(val *GlueMlTransformParametersFindMatchesParameters)
	PrecisionRecallTradeOff() *float64
	SetPrecisionRecallTradeOff(val *float64)
	PrecisionRecallTradeOffInput() *float64
	PrimaryKeyColumnName() *string
	SetPrimaryKeyColumnName(val *string)
	PrimaryKeyColumnNameInput() *string
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
	ResetAccuracyCostTradeOff()
	ResetEnforceProvidedLabels()
	ResetPrecisionRecallTradeOff()
	ResetPrimaryKeyColumnName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueMlTransformParametersFindMatchesParametersOutputReference
type jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) AccuracyCostTradeOff() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accuracyCostTradeOff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) AccuracyCostTradeOffInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accuracyCostTradeOffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) EnforceProvidedLabels() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceProvidedLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) EnforceProvidedLabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceProvidedLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) InternalValue() *GlueMlTransformParametersFindMatchesParameters {
	var returns *GlueMlTransformParametersFindMatchesParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) PrecisionRecallTradeOff() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precisionRecallTradeOff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) PrecisionRecallTradeOffInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precisionRecallTradeOffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) PrimaryKeyColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKeyColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) PrimaryKeyColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKeyColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueMlTransformParametersFindMatchesParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GlueMlTransformParametersFindMatchesParametersOutputReference {
	_init_.Initialize()

	if err := validateNewGlueMlTransformParametersFindMatchesParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.glueMlTransform.GlueMlTransformParametersFindMatchesParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueMlTransformParametersFindMatchesParametersOutputReference_Override(g GlueMlTransformParametersFindMatchesParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.glueMlTransform.GlueMlTransformParametersFindMatchesParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference)SetAccuracyCostTradeOff(val *float64) {
	if err := j.validateSetAccuracyCostTradeOffParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accuracyCostTradeOff",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference)SetEnforceProvidedLabels(val interface{}) {
	if err := j.validateSetEnforceProvidedLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceProvidedLabels",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference)SetInternalValue(val *GlueMlTransformParametersFindMatchesParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference)SetPrecisionRecallTradeOff(val *float64) {
	if err := j.validateSetPrecisionRecallTradeOffParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"precisionRecallTradeOff",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference)SetPrimaryKeyColumnName(val *string) {
	if err := j.validateSetPrimaryKeyColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryKeyColumnName",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) ResetAccuracyCostTradeOff() {
	_jsii_.InvokeVoid(
		g,
		"resetAccuracyCostTradeOff",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) ResetEnforceProvidedLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetEnforceProvidedLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) ResetPrecisionRecallTradeOff() {
	_jsii_.InvokeVoid(
		g,
		"resetPrecisionRecallTradeOff",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) ResetPrimaryKeyColumnName() {
	_jsii_.InvokeVoid(
		g,
		"resetPrimaryKeyColumnName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueMlTransformParametersFindMatchesParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

