// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsbatchjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/dataawsbatchjobdefinition/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference interface {
	cdktf.ComplexObject
	Command() *[]*string
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
	Environment() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerEnvironmentList
	EphemeralStorage() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerEphemeralStorageList
	ExecutionRoleArn() *string
	FargatePlatformConfiguration() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerFargatePlatformConfigurationList
	// Experimental.
	Fqn() *string
	Image() *string
	InstanceType() *string
	InternalValue() *DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainer
	SetInternalValue(val *DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainer)
	JobRoleArn() *string
	LinuxParameters() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersList
	LogConfiguration() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLogConfigurationList
	MountPoints() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerMountPointsList
	NetworkConfiguration() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerNetworkConfigurationList
	Privileged() cdktf.IResolvable
	ReadonlyRootFilesystem() cdktf.IResolvable
	ResourceRequirements() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerResourceRequirementsList
	RuntimePlatform() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerRuntimePlatformList
	Secrets() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerSecretsList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Ulimits() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerUlimitsList
	User() *string
	Volumes() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerVolumesList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference
type jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Environment() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerEnvironmentList {
	var returns DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerEnvironmentList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) EphemeralStorage() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerEphemeralStorageList {
	var returns DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerEphemeralStorageList
	_jsii_.Get(
		j,
		"ephemeralStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) FargatePlatformConfiguration() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerFargatePlatformConfigurationList {
	var returns DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerFargatePlatformConfigurationList
	_jsii_.Get(
		j,
		"fargatePlatformConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) InternalValue() *DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainer {
	var returns *DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainer
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) JobRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) LinuxParameters() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersList {
	var returns DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersList
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) LogConfiguration() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLogConfigurationList {
	var returns DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLogConfigurationList
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) MountPoints() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerMountPointsList {
	var returns DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerMountPointsList
	_jsii_.Get(
		j,
		"mountPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) NetworkConfiguration() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerNetworkConfigurationList {
	var returns DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerNetworkConfigurationList
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Privileged() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ReadonlyRootFilesystem() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"readonlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ResourceRequirements() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerResourceRequirementsList {
	var returns DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerResourceRequirementsList
	_jsii_.Get(
		j,
		"resourceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) RuntimePlatform() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerRuntimePlatformList {
	var returns DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerRuntimePlatformList
	_jsii_.Get(
		j,
		"runtimePlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Secrets() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerSecretsList {
	var returns DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerSecretsList
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Ulimits() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerUlimitsList {
	var returns DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerUlimitsList
	_jsii_.Get(
		j,
		"ulimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Volumes() DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerVolumesList {
	var returns DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerVolumesList
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


func NewDataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsBatchJobDefinition.DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference_Override(d DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsBatchJobDefinition.DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetInternalValue(val *DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainer) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

