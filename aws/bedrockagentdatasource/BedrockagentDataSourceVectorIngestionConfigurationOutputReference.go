// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/bedrockagentdatasource/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BedrockagentDataSourceVectorIngestionConfigurationOutputReference interface {
	cdktf.ComplexObject
	ChunkingConfiguration() BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationList
	ChunkingConfigurationInput() interface{}
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
	CustomTransformationConfiguration() BedrockagentDataSourceVectorIngestionConfigurationCustomTransformationConfigurationList
	CustomTransformationConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ParsingConfiguration() BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationList
	ParsingConfigurationInput() interface{}
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
	PutChunkingConfiguration(value interface{})
	PutCustomTransformationConfiguration(value interface{})
	PutParsingConfiguration(value interface{})
	ResetChunkingConfiguration()
	ResetCustomTransformationConfiguration()
	ResetParsingConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentDataSourceVectorIngestionConfigurationOutputReference
type jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) ChunkingConfiguration() BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationList {
	var returns BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationList
	_jsii_.Get(
		j,
		"chunkingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) ChunkingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chunkingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) CustomTransformationConfiguration() BedrockagentDataSourceVectorIngestionConfigurationCustomTransformationConfigurationList {
	var returns BedrockagentDataSourceVectorIngestionConfigurationCustomTransformationConfigurationList
	_jsii_.Get(
		j,
		"customTransformationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) CustomTransformationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customTransformationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) ParsingConfiguration() BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationList {
	var returns BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationList
	_jsii_.Get(
		j,
		"parsingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) ParsingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parsingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentDataSourceVectorIngestionConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentDataSourceVectorIngestionConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentDataSourceVectorIngestionConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentDataSource.BedrockagentDataSourceVectorIngestionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentDataSourceVectorIngestionConfigurationOutputReference_Override(b BedrockagentDataSourceVectorIngestionConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentDataSource.BedrockagentDataSourceVectorIngestionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) PutChunkingConfiguration(value interface{}) {
	if err := b.validatePutChunkingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putChunkingConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) PutCustomTransformationConfiguration(value interface{}) {
	if err := b.validatePutCustomTransformationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCustomTransformationConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) PutParsingConfiguration(value interface{}) {
	if err := b.validatePutParsingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putParsingConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) ResetChunkingConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetChunkingConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) ResetCustomTransformationConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetCustomTransformationConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) ResetParsingConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetParsingConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

