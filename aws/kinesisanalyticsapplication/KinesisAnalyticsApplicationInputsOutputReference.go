// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/kinesisanalyticsapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KinesisAnalyticsApplicationInputsOutputReference interface {
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
	InternalValue() *KinesisAnalyticsApplicationInputs
	SetInternalValue(val *KinesisAnalyticsApplicationInputs)
	KinesisFirehose() KinesisAnalyticsApplicationInputsKinesisFirehoseOutputReference
	KinesisFirehoseInput() *KinesisAnalyticsApplicationInputsKinesisFirehose
	KinesisStream() KinesisAnalyticsApplicationInputsKinesisStreamOutputReference
	KinesisStreamInput() *KinesisAnalyticsApplicationInputsKinesisStream
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	Parallelism() KinesisAnalyticsApplicationInputsParallelismOutputReference
	ParallelismInput() *KinesisAnalyticsApplicationInputsParallelism
	ProcessingConfiguration() KinesisAnalyticsApplicationInputsProcessingConfigurationOutputReference
	ProcessingConfigurationInput() *KinesisAnalyticsApplicationInputsProcessingConfiguration
	Schema() KinesisAnalyticsApplicationInputsSchemaOutputReference
	SchemaInput() *KinesisAnalyticsApplicationInputsSchema
	StartingPositionConfiguration() KinesisAnalyticsApplicationInputsStartingPositionConfigurationList
	StartingPositionConfigurationInput() interface{}
	StreamNames() *[]*string
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
	PutKinesisFirehose(value *KinesisAnalyticsApplicationInputsKinesisFirehose)
	PutKinesisStream(value *KinesisAnalyticsApplicationInputsKinesisStream)
	PutParallelism(value *KinesisAnalyticsApplicationInputsParallelism)
	PutProcessingConfiguration(value *KinesisAnalyticsApplicationInputsProcessingConfiguration)
	PutSchema(value *KinesisAnalyticsApplicationInputsSchema)
	PutStartingPositionConfiguration(value interface{})
	ResetKinesisFirehose()
	ResetKinesisStream()
	ResetParallelism()
	ResetProcessingConfiguration()
	ResetStartingPositionConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisAnalyticsApplicationInputsOutputReference
type jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) InternalValue() *KinesisAnalyticsApplicationInputs {
	var returns *KinesisAnalyticsApplicationInputs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) KinesisFirehose() KinesisAnalyticsApplicationInputsKinesisFirehoseOutputReference {
	var returns KinesisAnalyticsApplicationInputsKinesisFirehoseOutputReference
	_jsii_.Get(
		j,
		"kinesisFirehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) KinesisFirehoseInput() *KinesisAnalyticsApplicationInputsKinesisFirehose {
	var returns *KinesisAnalyticsApplicationInputsKinesisFirehose
	_jsii_.Get(
		j,
		"kinesisFirehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) KinesisStream() KinesisAnalyticsApplicationInputsKinesisStreamOutputReference {
	var returns KinesisAnalyticsApplicationInputsKinesisStreamOutputReference
	_jsii_.Get(
		j,
		"kinesisStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) KinesisStreamInput() *KinesisAnalyticsApplicationInputsKinesisStream {
	var returns *KinesisAnalyticsApplicationInputsKinesisStream
	_jsii_.Get(
		j,
		"kinesisStreamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) Parallelism() KinesisAnalyticsApplicationInputsParallelismOutputReference {
	var returns KinesisAnalyticsApplicationInputsParallelismOutputReference
	_jsii_.Get(
		j,
		"parallelism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) ParallelismInput() *KinesisAnalyticsApplicationInputsParallelism {
	var returns *KinesisAnalyticsApplicationInputsParallelism
	_jsii_.Get(
		j,
		"parallelismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) ProcessingConfiguration() KinesisAnalyticsApplicationInputsProcessingConfigurationOutputReference {
	var returns KinesisAnalyticsApplicationInputsProcessingConfigurationOutputReference
	_jsii_.Get(
		j,
		"processingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) ProcessingConfigurationInput() *KinesisAnalyticsApplicationInputsProcessingConfiguration {
	var returns *KinesisAnalyticsApplicationInputsProcessingConfiguration
	_jsii_.Get(
		j,
		"processingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) Schema() KinesisAnalyticsApplicationInputsSchemaOutputReference {
	var returns KinesisAnalyticsApplicationInputsSchemaOutputReference
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) SchemaInput() *KinesisAnalyticsApplicationInputsSchema {
	var returns *KinesisAnalyticsApplicationInputsSchema
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) StartingPositionConfiguration() KinesisAnalyticsApplicationInputsStartingPositionConfigurationList {
	var returns KinesisAnalyticsApplicationInputsStartingPositionConfigurationList
	_jsii_.Get(
		j,
		"startingPositionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) StartingPositionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startingPositionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) StreamNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"streamNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKinesisAnalyticsApplicationInputsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KinesisAnalyticsApplicationInputsOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisAnalyticsApplicationInputsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisAnalyticsApplication.KinesisAnalyticsApplicationInputsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisAnalyticsApplicationInputsOutputReference_Override(k KinesisAnalyticsApplicationInputsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisAnalyticsApplication.KinesisAnalyticsApplicationInputsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference)SetInternalValue(val *KinesisAnalyticsApplicationInputs) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) PutKinesisFirehose(value *KinesisAnalyticsApplicationInputsKinesisFirehose) {
	if err := k.validatePutKinesisFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putKinesisFirehose",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) PutKinesisStream(value *KinesisAnalyticsApplicationInputsKinesisStream) {
	if err := k.validatePutKinesisStreamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putKinesisStream",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) PutParallelism(value *KinesisAnalyticsApplicationInputsParallelism) {
	if err := k.validatePutParallelismParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putParallelism",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) PutProcessingConfiguration(value *KinesisAnalyticsApplicationInputsProcessingConfiguration) {
	if err := k.validatePutProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putProcessingConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) PutSchema(value *KinesisAnalyticsApplicationInputsSchema) {
	if err := k.validatePutSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSchema",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) PutStartingPositionConfiguration(value interface{}) {
	if err := k.validatePutStartingPositionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putStartingPositionConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) ResetKinesisFirehose() {
	_jsii_.InvokeVoid(
		k,
		"resetKinesisFirehose",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) ResetKinesisStream() {
	_jsii_.InvokeVoid(
		k,
		"resetKinesisStream",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) ResetParallelism() {
	_jsii_.InvokeVoid(
		k,
		"resetParallelism",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) ResetProcessingConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetProcessingConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) ResetStartingPositionConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetStartingPositionConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KinesisAnalyticsApplicationInputsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

