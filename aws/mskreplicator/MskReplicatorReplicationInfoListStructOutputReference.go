// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskreplicator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/mskreplicator/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MskReplicatorReplicationInfoListStructOutputReference interface {
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
	ConsumerGroupReplication() MskReplicatorReplicationInfoListConsumerGroupReplicationList
	ConsumerGroupReplicationInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *MskReplicatorReplicationInfoListStruct
	SetInternalValue(val *MskReplicatorReplicationInfoListStruct)
	SourceKafkaClusterAlias() *string
	SourceKafkaClusterArn() *string
	SetSourceKafkaClusterArn(val *string)
	SourceKafkaClusterArnInput() *string
	TargetCompressionType() *string
	SetTargetCompressionType(val *string)
	TargetCompressionTypeInput() *string
	TargetKafkaClusterAlias() *string
	TargetKafkaClusterArn() *string
	SetTargetKafkaClusterArn(val *string)
	TargetKafkaClusterArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TopicReplication() MskReplicatorReplicationInfoListTopicReplicationList
	TopicReplicationInput() interface{}
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
	PutConsumerGroupReplication(value interface{})
	PutTopicReplication(value interface{})
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MskReplicatorReplicationInfoListStructOutputReference
type jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) ConsumerGroupReplication() MskReplicatorReplicationInfoListConsumerGroupReplicationList {
	var returns MskReplicatorReplicationInfoListConsumerGroupReplicationList
	_jsii_.Get(
		j,
		"consumerGroupReplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) ConsumerGroupReplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"consumerGroupReplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) InternalValue() *MskReplicatorReplicationInfoListStruct {
	var returns *MskReplicatorReplicationInfoListStruct
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) SourceKafkaClusterAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceKafkaClusterAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) SourceKafkaClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceKafkaClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) SourceKafkaClusterArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceKafkaClusterArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) TargetCompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCompressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) TargetCompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCompressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) TargetKafkaClusterAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetKafkaClusterAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) TargetKafkaClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetKafkaClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) TargetKafkaClusterArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetKafkaClusterArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) TopicReplication() MskReplicatorReplicationInfoListTopicReplicationList {
	var returns MskReplicatorReplicationInfoListTopicReplicationList
	_jsii_.Get(
		j,
		"topicReplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) TopicReplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"topicReplicationInput",
		&returns,
	)
	return returns
}


func NewMskReplicatorReplicationInfoListStructOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MskReplicatorReplicationInfoListStructOutputReference {
	_init_.Initialize()

	if err := validateNewMskReplicatorReplicationInfoListStructOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.mskReplicator.MskReplicatorReplicationInfoListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMskReplicatorReplicationInfoListStructOutputReference_Override(m MskReplicatorReplicationInfoListStructOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.mskReplicator.MskReplicatorReplicationInfoListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference)SetInternalValue(val *MskReplicatorReplicationInfoListStruct) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference)SetSourceKafkaClusterArn(val *string) {
	if err := j.validateSetSourceKafkaClusterArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceKafkaClusterArn",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference)SetTargetCompressionType(val *string) {
	if err := j.validateSetTargetCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetCompressionType",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference)SetTargetKafkaClusterArn(val *string) {
	if err := j.validateSetTargetKafkaClusterArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetKafkaClusterArn",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) PutConsumerGroupReplication(value interface{}) {
	if err := m.validatePutConsumerGroupReplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putConsumerGroupReplication",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) PutTopicReplication(value interface{}) {
	if err := m.validatePutTopicReplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTopicReplication",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

