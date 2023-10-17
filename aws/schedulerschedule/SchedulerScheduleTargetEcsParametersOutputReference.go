// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schedulerschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/schedulerschedule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SchedulerScheduleTargetEcsParametersOutputReference interface {
	cdktf.ComplexObject
	CapacityProviderStrategy() SchedulerScheduleTargetEcsParametersCapacityProviderStrategyList
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
	InternalValue() *SchedulerScheduleTargetEcsParameters
	SetInternalValue(val *SchedulerScheduleTargetEcsParameters)
	LaunchType() *string
	SetLaunchType(val *string)
	LaunchTypeInput() *string
	NetworkConfiguration() SchedulerScheduleTargetEcsParametersNetworkConfigurationOutputReference
	NetworkConfigurationInput() *SchedulerScheduleTargetEcsParametersNetworkConfiguration
	PlacementConstraints() SchedulerScheduleTargetEcsParametersPlacementConstraintsList
	PlacementConstraintsInput() interface{}
	PlacementStrategy() SchedulerScheduleTargetEcsParametersPlacementStrategyList
	PlacementStrategyInput() interface{}
	PlatformVersion() *string
	SetPlatformVersion(val *string)
	PlatformVersionInput() *string
	PropagateTags() *string
	SetPropagateTags(val *string)
	PropagateTagsInput() *string
	ReferenceId() *string
	SetReferenceId(val *string)
	ReferenceIdInput() *string
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
	PutNetworkConfiguration(value *SchedulerScheduleTargetEcsParametersNetworkConfiguration)
	PutPlacementConstraints(value interface{})
	PutPlacementStrategy(value interface{})
	ResetCapacityProviderStrategy()
	ResetEnableEcsManagedTags()
	ResetEnableExecuteCommand()
	ResetGroup()
	ResetLaunchType()
	ResetNetworkConfiguration()
	ResetPlacementConstraints()
	ResetPlacementStrategy()
	ResetPlatformVersion()
	ResetPropagateTags()
	ResetReferenceId()
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

// The jsii proxy struct for SchedulerScheduleTargetEcsParametersOutputReference
type jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) CapacityProviderStrategy() SchedulerScheduleTargetEcsParametersCapacityProviderStrategyList {
	var returns SchedulerScheduleTargetEcsParametersCapacityProviderStrategyList
	_jsii_.Get(
		j,
		"capacityProviderStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) CapacityProviderStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityProviderStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) EnableEcsManagedTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) EnableEcsManagedTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) EnableExecuteCommand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) EnableExecuteCommandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) Group() *string {
	var returns *string
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) GroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) InternalValue() *SchedulerScheduleTargetEcsParameters {
	var returns *SchedulerScheduleTargetEcsParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) LaunchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) LaunchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) NetworkConfiguration() SchedulerScheduleTargetEcsParametersNetworkConfigurationOutputReference {
	var returns SchedulerScheduleTargetEcsParametersNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) NetworkConfigurationInput() *SchedulerScheduleTargetEcsParametersNetworkConfiguration {
	var returns *SchedulerScheduleTargetEcsParametersNetworkConfiguration
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) PlacementConstraints() SchedulerScheduleTargetEcsParametersPlacementConstraintsList {
	var returns SchedulerScheduleTargetEcsParametersPlacementConstraintsList
	_jsii_.Get(
		j,
		"placementConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) PlacementConstraintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) PlacementStrategy() SchedulerScheduleTargetEcsParametersPlacementStrategyList {
	var returns SchedulerScheduleTargetEcsParametersPlacementStrategyList
	_jsii_.Get(
		j,
		"placementStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) PlacementStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) PlatformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) PlatformVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) PropagateTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) PropagateTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) ReferenceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) ReferenceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) TaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) TaskCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) TaskDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) TaskDefinitionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSchedulerScheduleTargetEcsParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SchedulerScheduleTargetEcsParametersOutputReference {
	_init_.Initialize()

	if err := validateNewSchedulerScheduleTargetEcsParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.schedulerSchedule.SchedulerScheduleTargetEcsParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSchedulerScheduleTargetEcsParametersOutputReference_Override(s SchedulerScheduleTargetEcsParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.schedulerSchedule.SchedulerScheduleTargetEcsParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference)SetEnableEcsManagedTags(val interface{}) {
	if err := j.validateSetEnableEcsManagedTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableEcsManagedTags",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference)SetEnableExecuteCommand(val interface{}) {
	if err := j.validateSetEnableExecuteCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableExecuteCommand",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference)SetGroup(val *string) {
	if err := j.validateSetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"group",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference)SetInternalValue(val *SchedulerScheduleTargetEcsParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference)SetLaunchType(val *string) {
	if err := j.validateSetLaunchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchType",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference)SetPlatformVersion(val *string) {
	if err := j.validateSetPlatformVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformVersion",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference)SetPropagateTags(val *string) {
	if err := j.validateSetPropagateTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference)SetReferenceId(val *string) {
	if err := j.validateSetReferenceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referenceId",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference)SetTaskCount(val *float64) {
	if err := j.validateSetTaskCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskCount",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference)SetTaskDefinitionArn(val *string) {
	if err := j.validateSetTaskDefinitionArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskDefinitionArn",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) PutCapacityProviderStrategy(value interface{}) {
	if err := s.validatePutCapacityProviderStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCapacityProviderStrategy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) PutNetworkConfiguration(value *SchedulerScheduleTargetEcsParametersNetworkConfiguration) {
	if err := s.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) PutPlacementConstraints(value interface{}) {
	if err := s.validatePutPlacementConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPlacementConstraints",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) PutPlacementStrategy(value interface{}) {
	if err := s.validatePutPlacementStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPlacementStrategy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) ResetCapacityProviderStrategy() {
	_jsii_.InvokeVoid(
		s,
		"resetCapacityProviderStrategy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) ResetEnableEcsManagedTags() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableEcsManagedTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) ResetEnableExecuteCommand() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableExecuteCommand",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		s,
		"resetGroup",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) ResetLaunchType() {
	_jsii_.InvokeVoid(
		s,
		"resetLaunchType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) ResetPlacementConstraints() {
	_jsii_.InvokeVoid(
		s,
		"resetPlacementConstraints",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) ResetPlacementStrategy() {
	_jsii_.InvokeVoid(
		s,
		"resetPlacementStrategy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) ResetPlatformVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetPlatformVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) ResetPropagateTags() {
	_jsii_.InvokeVoid(
		s,
		"resetPropagateTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) ResetReferenceId() {
	_jsii_.InvokeVoid(
		s,
		"resetReferenceId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) ResetTaskCount() {
	_jsii_.InvokeVoid(
		s,
		"resetTaskCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchedulerScheduleTargetEcsParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

