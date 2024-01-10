// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskreplicator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/mskreplicator/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MskReplicatorReplicationInfoListTopicReplicationOutputReference interface {
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
	CopyAccessControlListsForTopics() interface{}
	SetCopyAccessControlListsForTopics(val interface{})
	CopyAccessControlListsForTopicsInput() interface{}
	CopyTopicConfigurations() interface{}
	SetCopyTopicConfigurations(val interface{})
	CopyTopicConfigurationsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DetectAndCopyNewTopics() interface{}
	SetDetectAndCopyNewTopics(val interface{})
	DetectAndCopyNewTopicsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TopicsToExclude() *[]*string
	SetTopicsToExclude(val *[]*string)
	TopicsToExcludeInput() *[]*string
	TopicsToReplicate() *[]*string
	SetTopicsToReplicate(val *[]*string)
	TopicsToReplicateInput() *[]*string
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
	ResetCopyAccessControlListsForTopics()
	ResetCopyTopicConfigurations()
	ResetDetectAndCopyNewTopics()
	ResetTopicsToExclude()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MskReplicatorReplicationInfoListTopicReplicationOutputReference
type jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) CopyAccessControlListsForTopics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyAccessControlListsForTopics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) CopyAccessControlListsForTopicsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyAccessControlListsForTopicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) CopyTopicConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTopicConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) CopyTopicConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTopicConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) DetectAndCopyNewTopics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectAndCopyNewTopics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) DetectAndCopyNewTopicsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectAndCopyNewTopicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) TopicsToExclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topicsToExclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) TopicsToExcludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topicsToExcludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) TopicsToReplicate() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topicsToReplicate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) TopicsToReplicateInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topicsToReplicateInput",
		&returns,
	)
	return returns
}


func NewMskReplicatorReplicationInfoListTopicReplicationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MskReplicatorReplicationInfoListTopicReplicationOutputReference {
	_init_.Initialize()

	if err := validateNewMskReplicatorReplicationInfoListTopicReplicationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.mskReplicator.MskReplicatorReplicationInfoListTopicReplicationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMskReplicatorReplicationInfoListTopicReplicationOutputReference_Override(m MskReplicatorReplicationInfoListTopicReplicationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.mskReplicator.MskReplicatorReplicationInfoListTopicReplicationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference)SetCopyAccessControlListsForTopics(val interface{}) {
	if err := j.validateSetCopyAccessControlListsForTopicsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyAccessControlListsForTopics",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference)SetCopyTopicConfigurations(val interface{}) {
	if err := j.validateSetCopyTopicConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTopicConfigurations",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference)SetDetectAndCopyNewTopics(val interface{}) {
	if err := j.validateSetDetectAndCopyNewTopicsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detectAndCopyNewTopics",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference)SetTopicsToExclude(val *[]*string) {
	if err := j.validateSetTopicsToExcludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topicsToExclude",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference)SetTopicsToReplicate(val *[]*string) {
	if err := j.validateSetTopicsToReplicateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topicsToReplicate",
		val,
	)
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) ResetCopyAccessControlListsForTopics() {
	_jsii_.InvokeVoid(
		m,
		"resetCopyAccessControlListsForTopics",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) ResetCopyTopicConfigurations() {
	_jsii_.InvokeVoid(
		m,
		"resetCopyTopicConfigurations",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) ResetDetectAndCopyNewTopics() {
	_jsii_.InvokeVoid(
		m,
		"resetDetectAndCopyNewTopics",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) ResetTopicsToExclude() {
	_jsii_.InvokeVoid(
		m,
		"resetTopicsToExclude",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MskReplicatorReplicationInfoListTopicReplicationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

