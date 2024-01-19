// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/lexv2modelsintent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference interface {
	cdktf.ComplexObject
	Button() Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardButtonList
	ButtonInput() interface{}
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
	ImageUrl() *string
	SetImageUrl(val *string)
	ImageUrlInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Subtitle() *string
	SetSubtitle(val *string)
	SubtitleInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
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
	PutButton(value interface{})
	ResetButton()
	ResetImageUrl()
	ResetSubtitle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference
type jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) Button() Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardButtonList {
	var returns Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardButtonList
	_jsii_.Get(
		j,
		"button",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) ButtonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"buttonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) ImageUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) ImageUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) Subtitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) SubtitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtitleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}


func NewLexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference {
	_init_.Initialize()

	if err := validateNewLexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference_Override(l Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference)SetImageUrl(val *string) {
	if err := j.validateSetImageUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageUrl",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference)SetSubtitle(val *string) {
	if err := j.validateSetSubtitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subtitle",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) PutButton(value interface{}) {
	if err := l.validatePutButtonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putButton",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) ResetButton() {
	_jsii_.InvokeVoid(
		l,
		"resetButton",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) ResetImageUrl() {
	_jsii_.InvokeVoid(
		l,
		"resetImageUrl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) ResetSubtitle() {
	_jsii_.InvokeVoid(
		l,
		"resetSubtitle",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_Lexv2ModelsIntentClosingSettingConditionalDefaultBranchResponseMessageGroupVariationImageResponseCardOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

