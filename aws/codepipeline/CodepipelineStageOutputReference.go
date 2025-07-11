// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/codepipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodepipelineStageOutputReference interface {
	cdktf.ComplexObject
	Action() CodepipelineStageActionList
	ActionInput() interface{}
	BeforeEntry() CodepipelineStageBeforeEntryOutputReference
	BeforeEntryInput() *CodepipelineStageBeforeEntry
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	OnFailure() CodepipelineStageOnFailureOutputReference
	OnFailureInput() *CodepipelineStageOnFailure
	OnSuccess() CodepipelineStageOnSuccessOutputReference
	OnSuccessInput() *CodepipelineStageOnSuccess
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
	PutAction(value interface{})
	PutBeforeEntry(value *CodepipelineStageBeforeEntry)
	PutOnFailure(value *CodepipelineStageOnFailure)
	PutOnSuccess(value *CodepipelineStageOnSuccess)
	ResetBeforeEntry()
	ResetOnFailure()
	ResetOnSuccess()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CodepipelineStageOutputReference
type jsiiProxy_CodepipelineStageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodepipelineStageOutputReference) Action() CodepipelineStageActionList {
	var returns CodepipelineStageActionList
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageOutputReference) ActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageOutputReference) BeforeEntry() CodepipelineStageBeforeEntryOutputReference {
	var returns CodepipelineStageBeforeEntryOutputReference
	_jsii_.Get(
		j,
		"beforeEntry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageOutputReference) BeforeEntryInput() *CodepipelineStageBeforeEntry {
	var returns *CodepipelineStageBeforeEntry
	_jsii_.Get(
		j,
		"beforeEntryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageOutputReference) OnFailure() CodepipelineStageOnFailureOutputReference {
	var returns CodepipelineStageOnFailureOutputReference
	_jsii_.Get(
		j,
		"onFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageOutputReference) OnFailureInput() *CodepipelineStageOnFailure {
	var returns *CodepipelineStageOnFailure
	_jsii_.Get(
		j,
		"onFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageOutputReference) OnSuccess() CodepipelineStageOnSuccessOutputReference {
	var returns CodepipelineStageOnSuccessOutputReference
	_jsii_.Get(
		j,
		"onSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageOutputReference) OnSuccessInput() *CodepipelineStageOnSuccess {
	var returns *CodepipelineStageOnSuccess
	_jsii_.Get(
		j,
		"onSuccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCodepipelineStageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CodepipelineStageOutputReference {
	_init_.Initialize()

	if err := validateNewCodepipelineStageOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodepipelineStageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.codepipeline.CodepipelineStageOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCodepipelineStageOutputReference_Override(c CodepipelineStageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.codepipeline.CodepipelineStageOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CodepipelineStageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CodepipelineStageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CodepipelineStageOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodepipelineStageOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CodepipelineStageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodepipelineStageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CodepipelineStageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineStageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineStageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineStageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineStageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineStageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineStageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineStageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineStageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineStageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineStageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineStageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineStageOutputReference) PutAction(value interface{}) {
	if err := c.validatePutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodepipelineStageOutputReference) PutBeforeEntry(value *CodepipelineStageBeforeEntry) {
	if err := c.validatePutBeforeEntryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBeforeEntry",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodepipelineStageOutputReference) PutOnFailure(value *CodepipelineStageOnFailure) {
	if err := c.validatePutOnFailureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOnFailure",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodepipelineStageOutputReference) PutOnSuccess(value *CodepipelineStageOnSuccess) {
	if err := c.validatePutOnSuccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOnSuccess",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodepipelineStageOutputReference) ResetBeforeEntry() {
	_jsii_.InvokeVoid(
		c,
		"resetBeforeEntry",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelineStageOutputReference) ResetOnFailure() {
	_jsii_.InvokeVoid(
		c,
		"resetOnFailure",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelineStageOutputReference) ResetOnSuccess() {
	_jsii_.InvokeVoid(
		c,
		"resetOnSuccess",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelineStageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineStageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

