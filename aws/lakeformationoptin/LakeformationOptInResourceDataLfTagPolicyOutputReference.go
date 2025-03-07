// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakeformationoptin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/lakeformationoptin/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LakeformationOptInResourceDataLfTagPolicyOutputReference interface {
	cdktf.ComplexObject
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
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
	Expression() *[]*string
	SetExpression(val *[]*string)
	ExpressionInput() *[]*string
	ExpressionName() *string
	SetExpressionName(val *string)
	ExpressionNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ResourceType() *string
	SetResourceType(val *string)
	ResourceTypeInput() *string
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
	ResetCatalogId()
	ResetExpression()
	ResetExpressionName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LakeformationOptInResourceDataLfTagPolicyOutputReference
type jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) Expression() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) ExpressionInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"expressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) ExpressionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) ExpressionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLakeformationOptInResourceDataLfTagPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LakeformationOptInResourceDataLfTagPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewLakeformationOptInResourceDataLfTagPolicyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.lakeformationOptIn.LakeformationOptInResourceDataLfTagPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLakeformationOptInResourceDataLfTagPolicyOutputReference_Override(l LakeformationOptInResourceDataLfTagPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lakeformationOptIn.LakeformationOptInResourceDataLfTagPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference)SetCatalogId(val *string) {
	if err := j.validateSetCatalogIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference)SetExpression(val *[]*string) {
	if err := j.validateSetExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expression",
		val,
	)
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference)SetExpressionName(val *string) {
	if err := j.validateSetExpressionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expressionName",
		val,
	)
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference)SetResourceType(val *string) {
	if err := j.validateSetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) ResetCatalogId() {
	_jsii_.InvokeVoid(
		l,
		"resetCatalogId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) ResetExpression() {
	_jsii_.InvokeVoid(
		l,
		"resetExpression",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) ResetExpressionName() {
	_jsii_.InvokeVoid(
		l,
		"resetExpressionName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationOptInResourceDataLfTagPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

