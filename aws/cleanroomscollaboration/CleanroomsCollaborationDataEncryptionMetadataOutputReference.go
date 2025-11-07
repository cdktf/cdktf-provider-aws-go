// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomscollaboration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/cleanroomscollaboration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CleanroomsCollaborationDataEncryptionMetadataOutputReference interface {
	cdktf.ComplexObject
	AllowClearText() interface{}
	SetAllowClearText(val interface{})
	AllowClearTextInput() interface{}
	AllowDuplicates() interface{}
	SetAllowDuplicates(val interface{})
	AllowDuplicatesInput() interface{}
	AllowJoinsOnColumnsWithDifferentNames() interface{}
	SetAllowJoinsOnColumnsWithDifferentNames(val interface{})
	AllowJoinsOnColumnsWithDifferentNamesInput() interface{}
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
	InternalValue() *CleanroomsCollaborationDataEncryptionMetadata
	SetInternalValue(val *CleanroomsCollaborationDataEncryptionMetadata)
	PreserveNulls() interface{}
	SetPreserveNulls(val interface{})
	PreserveNullsInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CleanroomsCollaborationDataEncryptionMetadataOutputReference
type jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) AllowClearText() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClearText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) AllowClearTextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClearTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) AllowDuplicates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowDuplicates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) AllowDuplicatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowDuplicatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) AllowJoinsOnColumnsWithDifferentNames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowJoinsOnColumnsWithDifferentNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) AllowJoinsOnColumnsWithDifferentNamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowJoinsOnColumnsWithDifferentNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) InternalValue() *CleanroomsCollaborationDataEncryptionMetadata {
	var returns *CleanroomsCollaborationDataEncryptionMetadata
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) PreserveNulls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveNulls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) PreserveNullsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveNullsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCleanroomsCollaborationDataEncryptionMetadataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CleanroomsCollaborationDataEncryptionMetadataOutputReference {
	_init_.Initialize()

	if err := validateNewCleanroomsCollaborationDataEncryptionMetadataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cleanroomsCollaboration.CleanroomsCollaborationDataEncryptionMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCleanroomsCollaborationDataEncryptionMetadataOutputReference_Override(c CleanroomsCollaborationDataEncryptionMetadataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cleanroomsCollaboration.CleanroomsCollaborationDataEncryptionMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference)SetAllowClearText(val interface{}) {
	if err := j.validateSetAllowClearTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowClearText",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference)SetAllowDuplicates(val interface{}) {
	if err := j.validateSetAllowDuplicatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowDuplicates",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference)SetAllowJoinsOnColumnsWithDifferentNames(val interface{}) {
	if err := j.validateSetAllowJoinsOnColumnsWithDifferentNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowJoinsOnColumnsWithDifferentNames",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference)SetInternalValue(val *CleanroomsCollaborationDataEncryptionMetadata) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference)SetPreserveNulls(val interface{}) {
	if err := j.validateSetPreserveNullsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveNulls",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsCollaborationDataEncryptionMetadataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

