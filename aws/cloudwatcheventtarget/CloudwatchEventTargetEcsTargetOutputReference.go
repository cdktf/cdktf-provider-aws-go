// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatcheventtarget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/cloudwatcheventtarget/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudwatchEventTargetEcsTargetOutputReference interface {
	cdktf.ComplexObject
	CapacityProviderStrategy() CloudwatchEventTargetEcsTargetCapacityProviderStrategyList
	CapacityProviderStrategyInput() interface{}
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
	EnableEcsManagedTags() interface{}
	SetEnableEcsManagedTags(val interface{})
	EnableEcsManagedTagsInput() interface{}
	EnableExecuteCommand() interface{}
	SetEnableExecuteCommand(val interface{})
	EnableExecuteCommandInput() interface{}
	// Experimental.
	Fqn() *string
	Group() *string
	SetGroup(val *string)
	GroupInput() *string
	InternalValue() *CloudwatchEventTargetEcsTarget
	SetInternalValue(val *CloudwatchEventTargetEcsTarget)
	LaunchType() *string
	SetLaunchType(val *string)
	LaunchTypeInput() *string
	NetworkConfiguration() CloudwatchEventTargetEcsTargetNetworkConfigurationOutputReference
	NetworkConfigurationInput() *CloudwatchEventTargetEcsTargetNetworkConfiguration
	OrderedPlacementStrategy() CloudwatchEventTargetEcsTargetOrderedPlacementStrategyList
	OrderedPlacementStrategyInput() interface{}
	PlacementConstraint() CloudwatchEventTargetEcsTargetPlacementConstraintList
	PlacementConstraintInput() interface{}
	PlatformVersion() *string
	SetPlatformVersion(val *string)
	PlatformVersionInput() *string
	PropagateTags() *string
	SetPropagateTags(val *string)
	PropagateTagsInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TaskCount() *float64
	SetTaskCount(val *float64)
	TaskCountInput() *float64
	TaskDefinitionArn() *string
	SetTaskDefinitionArn(val *string)
	TaskDefinitionArnInput() *string
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
	PutCapacityProviderStrategy(value interface{})
	PutNetworkConfiguration(value *CloudwatchEventTargetEcsTargetNetworkConfiguration)
	PutOrderedPlacementStrategy(value interface{})
	PutPlacementConstraint(value interface{})
	ResetCapacityProviderStrategy()
	ResetEnableEcsManagedTags()
	ResetEnableExecuteCommand()
	ResetGroup()
	ResetLaunchType()
	ResetNetworkConfiguration()
	ResetOrderedPlacementStrategy()
	ResetPlacementConstraint()
	ResetPlatformVersion()
	ResetPropagateTags()
	ResetTags()
	ResetTaskCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudwatchEventTargetEcsTargetOutputReference
type jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) CapacityProviderStrategy() CloudwatchEventTargetEcsTargetCapacityProviderStrategyList {
	var returns CloudwatchEventTargetEcsTargetCapacityProviderStrategyList
	_jsii_.Get(
		j,
		"capacityProviderStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) CapacityProviderStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityProviderStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) EnableEcsManagedTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) EnableEcsManagedTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) EnableExecuteCommand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) EnableExecuteCommandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) Group() *string {
	var returns *string
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) GroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) InternalValue() *CloudwatchEventTargetEcsTarget {
	var returns *CloudwatchEventTargetEcsTarget
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) LaunchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) LaunchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) NetworkConfiguration() CloudwatchEventTargetEcsTargetNetworkConfigurationOutputReference {
	var returns CloudwatchEventTargetEcsTargetNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) NetworkConfigurationInput() *CloudwatchEventTargetEcsTargetNetworkConfiguration {
	var returns *CloudwatchEventTargetEcsTargetNetworkConfiguration
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) OrderedPlacementStrategy() CloudwatchEventTargetEcsTargetOrderedPlacementStrategyList {
	var returns CloudwatchEventTargetEcsTargetOrderedPlacementStrategyList
	_jsii_.Get(
		j,
		"orderedPlacementStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) OrderedPlacementStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orderedPlacementStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) PlacementConstraint() CloudwatchEventTargetEcsTargetPlacementConstraintList {
	var returns CloudwatchEventTargetEcsTargetPlacementConstraintList
	_jsii_.Get(
		j,
		"placementConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) PlacementConstraintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) PlatformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) PlatformVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) PropagateTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) PropagateTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) TaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) TaskCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) TaskDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) TaskDefinitionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudwatchEventTargetEcsTargetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudwatchEventTargetEcsTargetOutputReference {
	_init_.Initialize()

	if err := validateNewCloudwatchEventTargetEcsTargetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudwatchEventTarget.CloudwatchEventTargetEcsTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudwatchEventTargetEcsTargetOutputReference_Override(c CloudwatchEventTargetEcsTargetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudwatchEventTarget.CloudwatchEventTargetEcsTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference)SetEnableEcsManagedTags(val interface{}) {
	if err := j.validateSetEnableEcsManagedTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableEcsManagedTags",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference)SetEnableExecuteCommand(val interface{}) {
	if err := j.validateSetEnableExecuteCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableExecuteCommand",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference)SetGroup(val *string) {
	if err := j.validateSetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"group",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference)SetInternalValue(val *CloudwatchEventTargetEcsTarget) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference)SetLaunchType(val *string) {
	if err := j.validateSetLaunchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchType",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference)SetPlatformVersion(val *string) {
	if err := j.validateSetPlatformVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformVersion",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference)SetPropagateTags(val *string) {
	if err := j.validateSetPropagateTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference)SetTaskCount(val *float64) {
	if err := j.validateSetTaskCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskCount",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference)SetTaskDefinitionArn(val *string) {
	if err := j.validateSetTaskDefinitionArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskDefinitionArn",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) PutCapacityProviderStrategy(value interface{}) {
	if err := c.validatePutCapacityProviderStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCapacityProviderStrategy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) PutNetworkConfiguration(value *CloudwatchEventTargetEcsTargetNetworkConfiguration) {
	if err := c.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) PutOrderedPlacementStrategy(value interface{}) {
	if err := c.validatePutOrderedPlacementStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOrderedPlacementStrategy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) PutPlacementConstraint(value interface{}) {
	if err := c.validatePutPlacementConstraintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPlacementConstraint",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) ResetCapacityProviderStrategy() {
	_jsii_.InvokeVoid(
		c,
		"resetCapacityProviderStrategy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) ResetEnableEcsManagedTags() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableEcsManagedTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) ResetEnableExecuteCommand() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableExecuteCommand",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) ResetLaunchType() {
	_jsii_.InvokeVoid(
		c,
		"resetLaunchType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) ResetOrderedPlacementStrategy() {
	_jsii_.InvokeVoid(
		c,
		"resetOrderedPlacementStrategy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) ResetPlacementConstraint() {
	_jsii_.InvokeVoid(
		c,
		"resetPlacementConstraint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) ResetPlatformVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetPlatformVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) ResetPropagateTags() {
	_jsii_.InvokeVoid(
		c,
		"resetPropagateTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) ResetTaskCount() {
	_jsii_.InvokeVoid(
		c,
		"resetTaskCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

