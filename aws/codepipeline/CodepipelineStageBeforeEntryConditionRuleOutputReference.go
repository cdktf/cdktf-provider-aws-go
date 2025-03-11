// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/codepipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodepipelineStageBeforeEntryConditionRuleOutputReference interface {
	cdktf.ComplexObject
	Commands() *[]*string
	SetCommands(val *[]*string)
	CommandsInput() *[]*string
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
	Configuration() *map[string]*string
	SetConfiguration(val *map[string]*string)
	ConfigurationInput() *map[string]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InputArtifacts() *[]*string
	SetInputArtifacts(val *[]*string)
	InputArtifactsInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	RuleTypeId() CodepipelineStageBeforeEntryConditionRuleRuleTypeIdOutputReference
	RuleTypeIdInput() *CodepipelineStageBeforeEntryConditionRuleRuleTypeId
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutInMinutes() *float64
	SetTimeoutInMinutes(val *float64)
	TimeoutInMinutesInput() *float64
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
	PutRuleTypeId(value *CodepipelineStageBeforeEntryConditionRuleRuleTypeId)
	ResetCommands()
	ResetConfiguration()
	ResetInputArtifacts()
	ResetRegion()
	ResetRoleArn()
	ResetTimeoutInMinutes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CodepipelineStageBeforeEntryConditionRuleOutputReference
type jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) Commands() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commands",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) CommandsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) Configuration() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) ConfigurationInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) InputArtifacts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) InputArtifactsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputArtifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) RuleTypeId() CodepipelineStageBeforeEntryConditionRuleRuleTypeIdOutputReference {
	var returns CodepipelineStageBeforeEntryConditionRuleRuleTypeIdOutputReference
	_jsii_.Get(
		j,
		"ruleTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) RuleTypeIdInput() *CodepipelineStageBeforeEntryConditionRuleRuleTypeId {
	var returns *CodepipelineStageBeforeEntryConditionRuleRuleTypeId
	_jsii_.Get(
		j,
		"ruleTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) TimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) TimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMinutesInput",
		&returns,
	)
	return returns
}


func NewCodepipelineStageBeforeEntryConditionRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CodepipelineStageBeforeEntryConditionRuleOutputReference {
	_init_.Initialize()

	if err := validateNewCodepipelineStageBeforeEntryConditionRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.codepipeline.CodepipelineStageBeforeEntryConditionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCodepipelineStageBeforeEntryConditionRuleOutputReference_Override(c CodepipelineStageBeforeEntryConditionRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.codepipeline.CodepipelineStageBeforeEntryConditionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference)SetCommands(val *[]*string) {
	if err := j.validateSetCommandsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commands",
		val,
	)
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference)SetConfiguration(val *map[string]*string) {
	if err := j.validateSetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference)SetInputArtifacts(val *[]*string) {
	if err := j.validateSetInputArtifactsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputArtifacts",
		val,
	)
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference)SetTimeoutInMinutes(val *float64) {
	if err := j.validateSetTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInMinutes",
		val,
	)
}

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) PutRuleTypeId(value *CodepipelineStageBeforeEntryConditionRuleRuleTypeId) {
	if err := c.validatePutRuleTypeIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRuleTypeId",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) ResetCommands() {
	_jsii_.InvokeVoid(
		c,
		"resetCommands",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) ResetConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) ResetInputArtifacts() {
	_jsii_.InvokeVoid(
		c,
		"resetInputArtifacts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		c,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) ResetTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeoutInMinutes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineStageBeforeEntryConditionRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

