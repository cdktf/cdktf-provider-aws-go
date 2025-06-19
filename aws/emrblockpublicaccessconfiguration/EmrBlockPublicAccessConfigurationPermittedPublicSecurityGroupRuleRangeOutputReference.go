// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrblockpublicaccessconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/emrblockpublicaccessconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxRange() *float64
	SetMaxRange(val *float64)
	MaxRangeInput() *float64
	MinRange() *float64
	SetMinRange(val *float64)
	MinRangeInput() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference
type jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) MaxRange() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) MaxRangeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) MinRange() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) MinRangeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference {
	_init_.Initialize()

	if err := validateNewEmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.emrBlockPublicAccessConfiguration.EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference_Override(e EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.emrBlockPublicAccessConfiguration.EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference)SetMaxRange(val *float64) {
	if err := j.validateSetMaxRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRange",
		val,
	)
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference)SetMinRange(val *float64) {
	if err := j.validateSetMinRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minRange",
		val,
	)
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EmrBlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

