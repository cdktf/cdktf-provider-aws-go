// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessanalyzeranalyzer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/accessanalyzeranalyzer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference interface {
	cdktf.ComplexObject
	AnalysisRule() AccessanalyzerAnalyzerConfigurationUnusedAccessAnalysisRuleOutputReference
	AnalysisRuleInput() *AccessanalyzerAnalyzerConfigurationUnusedAccessAnalysisRule
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
	InternalValue() *AccessanalyzerAnalyzerConfigurationUnusedAccess
	SetInternalValue(val *AccessanalyzerAnalyzerConfigurationUnusedAccess)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UnusedAccessAge() *float64
	SetUnusedAccessAge(val *float64)
	UnusedAccessAgeInput() *float64
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
	PutAnalysisRule(value *AccessanalyzerAnalyzerConfigurationUnusedAccessAnalysisRule)
	ResetAnalysisRule()
	ResetUnusedAccessAge()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference
type jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) AnalysisRule() AccessanalyzerAnalyzerConfigurationUnusedAccessAnalysisRuleOutputReference {
	var returns AccessanalyzerAnalyzerConfigurationUnusedAccessAnalysisRuleOutputReference
	_jsii_.Get(
		j,
		"analysisRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) AnalysisRuleInput() *AccessanalyzerAnalyzerConfigurationUnusedAccessAnalysisRule {
	var returns *AccessanalyzerAnalyzerConfigurationUnusedAccessAnalysisRule
	_jsii_.Get(
		j,
		"analysisRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) InternalValue() *AccessanalyzerAnalyzerConfigurationUnusedAccess {
	var returns *AccessanalyzerAnalyzerConfigurationUnusedAccess
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) UnusedAccessAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unusedAccessAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) UnusedAccessAgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unusedAccessAgeInput",
		&returns,
	)
	return returns
}


func NewAccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference {
	_init_.Initialize()

	if err := validateNewAccessanalyzerAnalyzerConfigurationUnusedAccessOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.accessanalyzerAnalyzer.AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference_Override(a AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.accessanalyzerAnalyzer.AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference)SetInternalValue(val *AccessanalyzerAnalyzerConfigurationUnusedAccess) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference)SetUnusedAccessAge(val *float64) {
	if err := j.validateSetUnusedAccessAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unusedAccessAge",
		val,
	)
}

func (a *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) PutAnalysisRule(value *AccessanalyzerAnalyzerConfigurationUnusedAccessAnalysisRule) {
	if err := a.validatePutAnalysisRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAnalysisRule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) ResetAnalysisRule() {
	_jsii_.InvokeVoid(
		a,
		"resetAnalysisRule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) ResetUnusedAccessAge() {
	_jsii_.InvokeVoid(
		a,
		"resetUnusedAccessAge",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccessanalyzerAnalyzerConfigurationUnusedAccessOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

