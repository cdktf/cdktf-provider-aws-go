// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/opensearchdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpensearchDomainAimlOptionsOutputReference interface {
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
	InternalValue() *OpensearchDomainAimlOptions
	SetInternalValue(val *OpensearchDomainAimlOptions)
	NaturalLanguageQueryGenerationOptions() OpensearchDomainAimlOptionsNaturalLanguageQueryGenerationOptionsOutputReference
	NaturalLanguageQueryGenerationOptionsInput() *OpensearchDomainAimlOptionsNaturalLanguageQueryGenerationOptions
	S3VectorsEngine() OpensearchDomainAimlOptionsS3VectorsEngineOutputReference
	S3VectorsEngineInput() *OpensearchDomainAimlOptionsS3VectorsEngine
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
	PutNaturalLanguageQueryGenerationOptions(value *OpensearchDomainAimlOptionsNaturalLanguageQueryGenerationOptions)
	PutS3VectorsEngine(value *OpensearchDomainAimlOptionsS3VectorsEngine)
	ResetNaturalLanguageQueryGenerationOptions()
	ResetS3VectorsEngine()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OpensearchDomainAimlOptionsOutputReference
type jsiiProxy_OpensearchDomainAimlOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) InternalValue() *OpensearchDomainAimlOptions {
	var returns *OpensearchDomainAimlOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) NaturalLanguageQueryGenerationOptions() OpensearchDomainAimlOptionsNaturalLanguageQueryGenerationOptionsOutputReference {
	var returns OpensearchDomainAimlOptionsNaturalLanguageQueryGenerationOptionsOutputReference
	_jsii_.Get(
		j,
		"naturalLanguageQueryGenerationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) NaturalLanguageQueryGenerationOptionsInput() *OpensearchDomainAimlOptionsNaturalLanguageQueryGenerationOptions {
	var returns *OpensearchDomainAimlOptionsNaturalLanguageQueryGenerationOptions
	_jsii_.Get(
		j,
		"naturalLanguageQueryGenerationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) S3VectorsEngine() OpensearchDomainAimlOptionsS3VectorsEngineOutputReference {
	var returns OpensearchDomainAimlOptionsS3VectorsEngineOutputReference
	_jsii_.Get(
		j,
		"s3VectorsEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) S3VectorsEngineInput() *OpensearchDomainAimlOptionsS3VectorsEngine {
	var returns *OpensearchDomainAimlOptionsS3VectorsEngine
	_jsii_.Get(
		j,
		"s3VectorsEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOpensearchDomainAimlOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OpensearchDomainAimlOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewOpensearchDomainAimlOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpensearchDomainAimlOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.opensearchDomain.OpensearchDomainAimlOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOpensearchDomainAimlOptionsOutputReference_Override(o OpensearchDomainAimlOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.opensearchDomain.OpensearchDomainAimlOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OpensearchDomainAimlOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainAimlOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainAimlOptionsOutputReference)SetInternalValue(val *OpensearchDomainAimlOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainAimlOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainAimlOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) PutNaturalLanguageQueryGenerationOptions(value *OpensearchDomainAimlOptionsNaturalLanguageQueryGenerationOptions) {
	if err := o.validatePutNaturalLanguageQueryGenerationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNaturalLanguageQueryGenerationOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) PutS3VectorsEngine(value *OpensearchDomainAimlOptionsS3VectorsEngine) {
	if err := o.validatePutS3VectorsEngineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putS3VectorsEngine",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) ResetNaturalLanguageQueryGenerationOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetNaturalLanguageQueryGenerationOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) ResetS3VectorsEngine() {
	_jsii_.InvokeVoid(
		o,
		"resetS3VectorsEngine",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainAimlOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

