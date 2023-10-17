// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyspacestable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/keyspacestable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KeyspacesTableSchemaDefinitionOutputReference interface {
	cdktf.ComplexObject
	ClusteringKey() KeyspacesTableSchemaDefinitionClusteringKeyList
	ClusteringKeyInput() interface{}
	Column() KeyspacesTableSchemaDefinitionColumnList
	ColumnInput() interface{}
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
	InternalValue() *KeyspacesTableSchemaDefinition
	SetInternalValue(val *KeyspacesTableSchemaDefinition)
	PartitionKey() KeyspacesTableSchemaDefinitionPartitionKeyList
	PartitionKeyInput() interface{}
	StaticColumn() KeyspacesTableSchemaDefinitionStaticColumnList
	StaticColumnInput() interface{}
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
	PutClusteringKey(value interface{})
	PutColumn(value interface{})
	PutPartitionKey(value interface{})
	PutStaticColumn(value interface{})
	ResetClusteringKey()
	ResetStaticColumn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KeyspacesTableSchemaDefinitionOutputReference
type jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) ClusteringKey() KeyspacesTableSchemaDefinitionClusteringKeyList {
	var returns KeyspacesTableSchemaDefinitionClusteringKeyList
	_jsii_.Get(
		j,
		"clusteringKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) ClusteringKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusteringKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) Column() KeyspacesTableSchemaDefinitionColumnList {
	var returns KeyspacesTableSchemaDefinitionColumnList
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) ColumnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) InternalValue() *KeyspacesTableSchemaDefinition {
	var returns *KeyspacesTableSchemaDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) PartitionKey() KeyspacesTableSchemaDefinitionPartitionKeyList {
	var returns KeyspacesTableSchemaDefinitionPartitionKeyList
	_jsii_.Get(
		j,
		"partitionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) PartitionKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) StaticColumn() KeyspacesTableSchemaDefinitionStaticColumnList {
	var returns KeyspacesTableSchemaDefinitionStaticColumnList
	_jsii_.Get(
		j,
		"staticColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) StaticColumnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKeyspacesTableSchemaDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KeyspacesTableSchemaDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewKeyspacesTableSchemaDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.keyspacesTable.KeyspacesTableSchemaDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKeyspacesTableSchemaDefinitionOutputReference_Override(k KeyspacesTableSchemaDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.keyspacesTable.KeyspacesTableSchemaDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference)SetInternalValue(val *KeyspacesTableSchemaDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) PutClusteringKey(value interface{}) {
	if err := k.validatePutClusteringKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putClusteringKey",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) PutColumn(value interface{}) {
	if err := k.validatePutColumnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putColumn",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) PutPartitionKey(value interface{}) {
	if err := k.validatePutPartitionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putPartitionKey",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) PutStaticColumn(value interface{}) {
	if err := k.validatePutStaticColumnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putStaticColumn",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) ResetClusteringKey() {
	_jsii_.InvokeVoid(
		k,
		"resetClusteringKey",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) ResetStaticColumn() {
	_jsii_.InvokeVoid(
		k,
		"resetStaticColumn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

