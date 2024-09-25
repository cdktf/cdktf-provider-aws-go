// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/bedrockagentdatasource/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference interface {
	cdktf.ComplexObject
	ChunkingStrategy() *string
	SetChunkingStrategy(val *string)
	ChunkingStrategyInput() *string
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
	FixedSizeChunkingConfiguration() BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationFixedSizeChunkingConfigurationList
	FixedSizeChunkingConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	HierarchicalChunkingConfiguration() BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationHierarchicalChunkingConfigurationList
	HierarchicalChunkingConfigurationInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SemanticChunkingConfiguration() BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationSemanticChunkingConfigurationList
	SemanticChunkingConfigurationInput() interface{}
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
	PutFixedSizeChunkingConfiguration(value interface{})
	PutHierarchicalChunkingConfiguration(value interface{})
	PutSemanticChunkingConfiguration(value interface{})
	ResetFixedSizeChunkingConfiguration()
	ResetHierarchicalChunkingConfiguration()
	ResetSemanticChunkingConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference
type jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) ChunkingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chunkingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) ChunkingStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chunkingStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) FixedSizeChunkingConfiguration() BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationFixedSizeChunkingConfigurationList {
	var returns BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationFixedSizeChunkingConfigurationList
	_jsii_.Get(
		j,
		"fixedSizeChunkingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) FixedSizeChunkingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fixedSizeChunkingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) HierarchicalChunkingConfiguration() BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationHierarchicalChunkingConfigurationList {
	var returns BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationHierarchicalChunkingConfigurationList
	_jsii_.Get(
		j,
		"hierarchicalChunkingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) HierarchicalChunkingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hierarchicalChunkingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) SemanticChunkingConfiguration() BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationSemanticChunkingConfigurationList {
	var returns BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationSemanticChunkingConfigurationList
	_jsii_.Get(
		j,
		"semanticChunkingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) SemanticChunkingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"semanticChunkingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentDataSource.BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference_Override(b BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentDataSource.BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference)SetChunkingStrategy(val *string) {
	if err := j.validateSetChunkingStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chunkingStrategy",
		val,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) PutFixedSizeChunkingConfiguration(value interface{}) {
	if err := b.validatePutFixedSizeChunkingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putFixedSizeChunkingConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) PutHierarchicalChunkingConfiguration(value interface{}) {
	if err := b.validatePutHierarchicalChunkingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putHierarchicalChunkingConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) PutSemanticChunkingConfiguration(value interface{}) {
	if err := b.validatePutSemanticChunkingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSemanticChunkingConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) ResetFixedSizeChunkingConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetFixedSizeChunkingConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) ResetHierarchicalChunkingConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetHierarchicalChunkingConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) ResetSemanticChunkingConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetSemanticChunkingConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

