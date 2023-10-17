// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/kinesisanalyticsapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KinesisAnalyticsApplicationOutputsOutputReference interface {
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
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KinesisFirehose() KinesisAnalyticsApplicationOutputsKinesisFirehoseOutputReference
	KinesisFirehoseInput() *KinesisAnalyticsApplicationOutputsKinesisFirehose
	KinesisStream() KinesisAnalyticsApplicationOutputsKinesisStreamOutputReference
	KinesisStreamInput() *KinesisAnalyticsApplicationOutputsKinesisStream
	Lambda() KinesisAnalyticsApplicationOutputsLambdaOutputReference
	LambdaInput() *KinesisAnalyticsApplicationOutputsLambda
	Name() *string
	SetName(val *string)
	NameInput() *string
	Schema() KinesisAnalyticsApplicationOutputsSchemaOutputReference
	SchemaInput() *KinesisAnalyticsApplicationOutputsSchema
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
	PutKinesisFirehose(value *KinesisAnalyticsApplicationOutputsKinesisFirehose)
	PutKinesisStream(value *KinesisAnalyticsApplicationOutputsKinesisStream)
	PutLambda(value *KinesisAnalyticsApplicationOutputsLambda)
	PutSchema(value *KinesisAnalyticsApplicationOutputsSchema)
	ResetKinesisFirehose()
	ResetKinesisStream()
	ResetLambda()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisAnalyticsApplicationOutputsOutputReference
type jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) KinesisFirehose() KinesisAnalyticsApplicationOutputsKinesisFirehoseOutputReference {
	var returns KinesisAnalyticsApplicationOutputsKinesisFirehoseOutputReference
	_jsii_.Get(
		j,
		"kinesisFirehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) KinesisFirehoseInput() *KinesisAnalyticsApplicationOutputsKinesisFirehose {
	var returns *KinesisAnalyticsApplicationOutputsKinesisFirehose
	_jsii_.Get(
		j,
		"kinesisFirehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) KinesisStream() KinesisAnalyticsApplicationOutputsKinesisStreamOutputReference {
	var returns KinesisAnalyticsApplicationOutputsKinesisStreamOutputReference
	_jsii_.Get(
		j,
		"kinesisStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) KinesisStreamInput() *KinesisAnalyticsApplicationOutputsKinesisStream {
	var returns *KinesisAnalyticsApplicationOutputsKinesisStream
	_jsii_.Get(
		j,
		"kinesisStreamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) Lambda() KinesisAnalyticsApplicationOutputsLambdaOutputReference {
	var returns KinesisAnalyticsApplicationOutputsLambdaOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) LambdaInput() *KinesisAnalyticsApplicationOutputsLambda {
	var returns *KinesisAnalyticsApplicationOutputsLambda
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) Schema() KinesisAnalyticsApplicationOutputsSchemaOutputReference {
	var returns KinesisAnalyticsApplicationOutputsSchemaOutputReference
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) SchemaInput() *KinesisAnalyticsApplicationOutputsSchema {
	var returns *KinesisAnalyticsApplicationOutputsSchema
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKinesisAnalyticsApplicationOutputsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) KinesisAnalyticsApplicationOutputsOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisAnalyticsApplicationOutputsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisAnalyticsApplication.KinesisAnalyticsApplicationOutputsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewKinesisAnalyticsApplicationOutputsOutputReference_Override(k KinesisAnalyticsApplicationOutputsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisAnalyticsApplication.KinesisAnalyticsApplicationOutputsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		k,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) PutKinesisFirehose(value *KinesisAnalyticsApplicationOutputsKinesisFirehose) {
	if err := k.validatePutKinesisFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putKinesisFirehose",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) PutKinesisStream(value *KinesisAnalyticsApplicationOutputsKinesisStream) {
	if err := k.validatePutKinesisStreamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putKinesisStream",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) PutLambda(value *KinesisAnalyticsApplicationOutputsLambda) {
	if err := k.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putLambda",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) PutSchema(value *KinesisAnalyticsApplicationOutputsSchema) {
	if err := k.validatePutSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSchema",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) ResetKinesisFirehose() {
	_jsii_.InvokeVoid(
		k,
		"resetKinesisFirehose",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) ResetKinesisStream() {
	_jsii_.InvokeVoid(
		k,
		"resetKinesisStream",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) ResetLambda() {
	_jsii_.InvokeVoid(
		k,
		"resetLambda",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

