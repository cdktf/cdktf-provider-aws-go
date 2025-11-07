// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rekognitionstreamprocessor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/rekognitionstreamprocessor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference interface {
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
	Height() *float64
	SetHeight(val *float64)
	HeightInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Left() *float64
	SetLeft(val *float64)
	LeftInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Top() *float64
	SetTop(val *float64)
	TopInput() *float64
	Width() *float64
	SetWidth(val *float64)
	WidthInput() *float64
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetHeight()
	ResetLeft()
	ResetTop()
	ResetWidth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference
type jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) HeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) Left() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"left",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) LeftInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) Top() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"top",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) TopInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) WidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widthInput",
		&returns,
	)
	return returns
}


func NewRekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference {
	_init_.Initialize()

	if err := validateNewRekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.rekognitionStreamProcessor.RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewRekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference_Override(r RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.rekognitionStreamProcessor.RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference)SetHeight(val *float64) {
	if err := j.validateSetHeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"height",
		val,
	)
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference)SetLeft(val *float64) {
	if err := j.validateSetLeftParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"left",
		val,
	)
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference)SetTop(val *float64) {
	if err := j.validateSetTopParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"top",
		val,
	)
}

func (j *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference)SetWidth(val *float64) {
	if err := j.validateSetWidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"width",
		val,
	)
}

func (r *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) ResetHeight() {
	_jsii_.InvokeVoid(
		r,
		"resetHeight",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) ResetLeft() {
	_jsii_.InvokeVoid(
		r,
		"resetLeft",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) ResetTop() {
	_jsii_.InvokeVoid(
		r,
		"resetTop",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) ResetWidth() {
	_jsii_.InvokeVoid(
		r,
		"resetWidth",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RekognitionStreamProcessorRegionsOfInterestBoundingBoxOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

