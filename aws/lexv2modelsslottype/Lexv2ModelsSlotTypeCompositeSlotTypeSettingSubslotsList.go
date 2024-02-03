// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsslottype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/lexv2modelsslottype/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList
type jsiiProxy_Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewLexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList {
	_init_.Initialize()

	if err := validateNewLexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList{}

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsSlotType.Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewLexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList_Override(l Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsSlotType.Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		l,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := l.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		l,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList) Get(index *float64) Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsOutputReference {
	if err := l.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsOutputReference

	_jsii_.Invoke(
		l,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_Lexv2ModelsSlotTypeCompositeSlotTypeSettingSubslotsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

