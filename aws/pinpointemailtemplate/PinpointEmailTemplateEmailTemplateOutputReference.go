// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pinpointemailtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/pinpointemailtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PinpointEmailTemplateEmailTemplateOutputReference interface {
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
	DefaultSubstitutions() *string
	SetDefaultSubstitutions(val *string)
	DefaultSubstitutionsInput() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	Header() PinpointEmailTemplateEmailTemplateHeaderList
	HeaderInput() interface{}
	HtmlPart() *string
	SetHtmlPart(val *string)
	HtmlPartInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RecommenderId() *string
	SetRecommenderId(val *string)
	RecommenderIdInput() *string
	Subject() *string
	SetSubject(val *string)
	SubjectInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TextPart() *string
	SetTextPart(val *string)
	TextPartInput() *string
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
	PutHeader(value interface{})
	ResetDefaultSubstitutions()
	ResetDescription()
	ResetHeader()
	ResetHtmlPart()
	ResetRecommenderId()
	ResetSubject()
	ResetTextPart()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PinpointEmailTemplateEmailTemplateOutputReference
type jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) DefaultSubstitutions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSubstitutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) DefaultSubstitutionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSubstitutionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) Header() PinpointEmailTemplateEmailTemplateHeaderList {
	var returns PinpointEmailTemplateEmailTemplateHeaderList
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) HeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) HtmlPart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"htmlPart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) HtmlPartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"htmlPartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) RecommenderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommenderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) RecommenderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommenderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) SubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) TextPart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textPart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) TextPartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textPartInput",
		&returns,
	)
	return returns
}


func NewPinpointEmailTemplateEmailTemplateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PinpointEmailTemplateEmailTemplateOutputReference {
	_init_.Initialize()

	if err := validateNewPinpointEmailTemplateEmailTemplateOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.pinpointEmailTemplate.PinpointEmailTemplateEmailTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPinpointEmailTemplateEmailTemplateOutputReference_Override(p PinpointEmailTemplateEmailTemplateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.pinpointEmailTemplate.PinpointEmailTemplateEmailTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference)SetDefaultSubstitutions(val *string) {
	if err := j.validateSetDefaultSubstitutionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSubstitutions",
		val,
	)
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference)SetHtmlPart(val *string) {
	if err := j.validateSetHtmlPartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"htmlPart",
		val,
	)
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference)SetRecommenderId(val *string) {
	if err := j.validateSetRecommenderIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recommenderId",
		val,
	)
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference)SetSubject(val *string) {
	if err := j.validateSetSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subject",
		val,
	)
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference)SetTextPart(val *string) {
	if err := j.validateSetTextPartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textPart",
		val,
	)
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) PutHeader(value interface{}) {
	if err := p.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHeader",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) ResetDefaultSubstitutions() {
	_jsii_.InvokeVoid(
		p,
		"resetDefaultSubstitutions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		p,
		"resetDescription",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		p,
		"resetHeader",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) ResetHtmlPart() {
	_jsii_.InvokeVoid(
		p,
		"resetHtmlPart",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) ResetRecommenderId() {
	_jsii_.InvokeVoid(
		p,
		"resetRecommenderId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) ResetSubject() {
	_jsii_.InvokeVoid(
		p,
		"resetSubject",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) ResetTextPart() {
	_jsii_.InvokeVoid(
		p,
		"resetTextPart",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointEmailTemplateEmailTemplateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

