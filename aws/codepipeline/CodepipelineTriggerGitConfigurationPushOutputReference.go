// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/codepipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodepipelineTriggerGitConfigurationPushOutputReference interface {
	cdktf.ComplexObject
	Branches() CodepipelineTriggerGitConfigurationPushBranchesOutputReference
	BranchesInput() *CodepipelineTriggerGitConfigurationPushBranches
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
	FilePaths() CodepipelineTriggerGitConfigurationPushFilePathsOutputReference
	FilePathsInput() *CodepipelineTriggerGitConfigurationPushFilePaths
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Tags() CodepipelineTriggerGitConfigurationPushTagsOutputReference
	TagsInput() *CodepipelineTriggerGitConfigurationPushTags
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
	PutBranches(value *CodepipelineTriggerGitConfigurationPushBranches)
	PutFilePaths(value *CodepipelineTriggerGitConfigurationPushFilePaths)
	PutTags(value *CodepipelineTriggerGitConfigurationPushTags)
	ResetBranches()
	ResetFilePaths()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CodepipelineTriggerGitConfigurationPushOutputReference
type jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) Branches() CodepipelineTriggerGitConfigurationPushBranchesOutputReference {
	var returns CodepipelineTriggerGitConfigurationPushBranchesOutputReference
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) BranchesInput() *CodepipelineTriggerGitConfigurationPushBranches {
	var returns *CodepipelineTriggerGitConfigurationPushBranches
	_jsii_.Get(
		j,
		"branchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) FilePaths() CodepipelineTriggerGitConfigurationPushFilePathsOutputReference {
	var returns CodepipelineTriggerGitConfigurationPushFilePathsOutputReference
	_jsii_.Get(
		j,
		"filePaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) FilePathsInput() *CodepipelineTriggerGitConfigurationPushFilePaths {
	var returns *CodepipelineTriggerGitConfigurationPushFilePaths
	_jsii_.Get(
		j,
		"filePathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) Tags() CodepipelineTriggerGitConfigurationPushTagsOutputReference {
	var returns CodepipelineTriggerGitConfigurationPushTagsOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) TagsInput() *CodepipelineTriggerGitConfigurationPushTags {
	var returns *CodepipelineTriggerGitConfigurationPushTags
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCodepipelineTriggerGitConfigurationPushOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CodepipelineTriggerGitConfigurationPushOutputReference {
	_init_.Initialize()

	if err := validateNewCodepipelineTriggerGitConfigurationPushOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.codepipeline.CodepipelineTriggerGitConfigurationPushOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCodepipelineTriggerGitConfigurationPushOutputReference_Override(c CodepipelineTriggerGitConfigurationPushOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.codepipeline.CodepipelineTriggerGitConfigurationPushOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) PutBranches(value *CodepipelineTriggerGitConfigurationPushBranches) {
	if err := c.validatePutBranchesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBranches",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) PutFilePaths(value *CodepipelineTriggerGitConfigurationPushFilePaths) {
	if err := c.validatePutFilePathsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFilePaths",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) PutTags(value *CodepipelineTriggerGitConfigurationPushTags) {
	if err := c.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTags",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) ResetBranches() {
	_jsii_.InvokeVoid(
		c,
		"resetBranches",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) ResetFilePaths() {
	_jsii_.InvokeVoid(
		c,
		"resetFilePaths",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CodepipelineTriggerGitConfigurationPushOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

