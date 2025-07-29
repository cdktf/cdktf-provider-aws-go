// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ecstaskdefinition/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/ecs_task_definition aws_ecs_task_definition}.
type EcsTaskDefinition interface {
	cdktf.TerraformResource
	Arn() *string
	ArnWithoutRevision() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerDefinitions() *string
	SetContainerDefinitions(val *string)
	ContainerDefinitionsInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Cpu() *string
	SetCpu(val *string)
	CpuInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableFaultInjection() interface{}
	SetEnableFaultInjection(val interface{})
	EnableFaultInjectionInput() interface{}
	EphemeralStorage() EcsTaskDefinitionEphemeralStorageOutputReference
	EphemeralStorageInput() *EcsTaskDefinitionEphemeralStorage
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	ExecutionRoleArnInput() *string
	Family() *string
	SetFamily(val *string)
	FamilyInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IpcMode() *string
	SetIpcMode(val *string)
	IpcModeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Memory() *string
	SetMemory(val *string)
	MemoryInput() *string
	NetworkMode() *string
	SetNetworkMode(val *string)
	NetworkModeInput() *string
	// The tree node.
	Node() constructs.Node
	PidMode() *string
	SetPidMode(val *string)
	PidModeInput() *string
	PlacementConstraints() EcsTaskDefinitionPlacementConstraintsList
	PlacementConstraintsInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProxyConfiguration() EcsTaskDefinitionProxyConfigurationOutputReference
	ProxyConfigurationInput() *EcsTaskDefinitionProxyConfiguration
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RequiresCompatibilities() *[]*string
	SetRequiresCompatibilities(val *[]*string)
	RequiresCompatibilitiesInput() *[]*string
	Revision() *float64
	RuntimePlatform() EcsTaskDefinitionRuntimePlatformOutputReference
	RuntimePlatformInput() *EcsTaskDefinitionRuntimePlatform
	SkipDestroy() interface{}
	SetSkipDestroy(val interface{})
	SkipDestroyInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TaskRoleArn() *string
	SetTaskRoleArn(val *string)
	TaskRoleArnInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TrackLatest() interface{}
	SetTrackLatest(val interface{})
	TrackLatestInput() interface{}
	Volume() EcsTaskDefinitionVolumeList
	VolumeInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutEphemeralStorage(value *EcsTaskDefinitionEphemeralStorage)
	PutPlacementConstraints(value interface{})
	PutProxyConfiguration(value *EcsTaskDefinitionProxyConfiguration)
	PutRuntimePlatform(value *EcsTaskDefinitionRuntimePlatform)
	PutVolume(value interface{})
	ResetCpu()
	ResetEnableFaultInjection()
	ResetEphemeralStorage()
	ResetExecutionRoleArn()
	ResetId()
	ResetIpcMode()
	ResetMemory()
	ResetNetworkMode()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPidMode()
	ResetPlacementConstraints()
	ResetProxyConfiguration()
	ResetRegion()
	ResetRequiresCompatibilities()
	ResetRuntimePlatform()
	ResetSkipDestroy()
	ResetTags()
	ResetTagsAll()
	ResetTaskRoleArn()
	ResetTrackLatest()
	ResetVolume()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for EcsTaskDefinition
type jsiiProxy_EcsTaskDefinition struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EcsTaskDefinition) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) ArnWithoutRevision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnWithoutRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) ContainerDefinitions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) ContainerDefinitionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerDefinitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) Cpu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) CpuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) EnableFaultInjection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFaultInjection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) EnableFaultInjectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFaultInjectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) EphemeralStorage() EcsTaskDefinitionEphemeralStorageOutputReference {
	var returns EcsTaskDefinitionEphemeralStorageOutputReference
	_jsii_.Get(
		j,
		"ephemeralStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) EphemeralStorageInput() *EcsTaskDefinitionEphemeralStorage {
	var returns *EcsTaskDefinitionEphemeralStorage
	_jsii_.Get(
		j,
		"ephemeralStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) FamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"familyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) IpcMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipcMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) IpcModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipcModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) Memory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) MemoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) NetworkMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) NetworkModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) PidMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pidMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) PidModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pidModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) PlacementConstraints() EcsTaskDefinitionPlacementConstraintsList {
	var returns EcsTaskDefinitionPlacementConstraintsList
	_jsii_.Get(
		j,
		"placementConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) PlacementConstraintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) ProxyConfiguration() EcsTaskDefinitionProxyConfigurationOutputReference {
	var returns EcsTaskDefinitionProxyConfigurationOutputReference
	_jsii_.Get(
		j,
		"proxyConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) ProxyConfigurationInput() *EcsTaskDefinitionProxyConfiguration {
	var returns *EcsTaskDefinitionProxyConfiguration
	_jsii_.Get(
		j,
		"proxyConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) RequiresCompatibilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiresCompatibilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) RequiresCompatibilitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiresCompatibilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) Revision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"revision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) RuntimePlatform() EcsTaskDefinitionRuntimePlatformOutputReference {
	var returns EcsTaskDefinitionRuntimePlatformOutputReference
	_jsii_.Get(
		j,
		"runtimePlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) RuntimePlatformInput() *EcsTaskDefinitionRuntimePlatform {
	var returns *EcsTaskDefinitionRuntimePlatform
	_jsii_.Get(
		j,
		"runtimePlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) SkipDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) SkipDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) TaskRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) TaskRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) TrackLatest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trackLatest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) TrackLatestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trackLatestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) Volume() EcsTaskDefinitionVolumeList {
	var returns EcsTaskDefinitionVolumeList
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinition) VolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/ecs_task_definition aws_ecs_task_definition} Resource.
func NewEcsTaskDefinition(scope constructs.Construct, id *string, config *EcsTaskDefinitionConfig) EcsTaskDefinition {
	_init_.Initialize()

	if err := validateNewEcsTaskDefinitionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsTaskDefinition{}

	_jsii_.Create(
		"@cdktf/provider-aws.ecsTaskDefinition.EcsTaskDefinition",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/ecs_task_definition aws_ecs_task_definition} Resource.
func NewEcsTaskDefinition_Override(e EcsTaskDefinition, scope constructs.Construct, id *string, config *EcsTaskDefinitionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ecsTaskDefinition.EcsTaskDefinition",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetContainerDefinitions(val *string) {
	if err := j.validateSetContainerDefinitionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerDefinitions",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetCpu(val *string) {
	if err := j.validateSetCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpu",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetEnableFaultInjection(val interface{}) {
	if err := j.validateSetEnableFaultInjectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableFaultInjection",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetFamily(val *string) {
	if err := j.validateSetFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"family",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetIpcMode(val *string) {
	if err := j.validateSetIpcModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipcMode",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetMemory(val *string) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetNetworkMode(val *string) {
	if err := j.validateSetNetworkModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkMode",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetPidMode(val *string) {
	if err := j.validateSetPidModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pidMode",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetRequiresCompatibilities(val *[]*string) {
	if err := j.validateSetRequiresCompatibilitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiresCompatibilities",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetSkipDestroy(val interface{}) {
	if err := j.validateSetSkipDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipDestroy",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetTaskRoleArn(val *string) {
	if err := j.validateSetTaskRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskRoleArn",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinition)SetTrackLatest(val interface{}) {
	if err := j.validateSetTrackLatestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trackLatest",
		val,
	)
}

// Generates CDKTF code for importing a EcsTaskDefinition resource upon running "cdktf plan <stack-name>".
func EcsTaskDefinition_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEcsTaskDefinition_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ecsTaskDefinition.EcsTaskDefinition",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func EcsTaskDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcsTaskDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ecsTaskDefinition.EcsTaskDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EcsTaskDefinition_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcsTaskDefinition_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ecsTaskDefinition.EcsTaskDefinition",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EcsTaskDefinition_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcsTaskDefinition_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ecsTaskDefinition.EcsTaskDefinition",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EcsTaskDefinition_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.ecsTaskDefinition.EcsTaskDefinition",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EcsTaskDefinition) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EcsTaskDefinition) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EcsTaskDefinition) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinition) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinition) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinition) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinition) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinition) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinition) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinition) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinition) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinition) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinition) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EcsTaskDefinition) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinition) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EcsTaskDefinition) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EcsTaskDefinition) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EcsTaskDefinition) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EcsTaskDefinition) PutEphemeralStorage(value *EcsTaskDefinitionEphemeralStorage) {
	if err := e.validatePutEphemeralStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEphemeralStorage",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinition) PutPlacementConstraints(value interface{}) {
	if err := e.validatePutPlacementConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPlacementConstraints",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinition) PutProxyConfiguration(value *EcsTaskDefinitionProxyConfiguration) {
	if err := e.validatePutProxyConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putProxyConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinition) PutRuntimePlatform(value *EcsTaskDefinitionRuntimePlatform) {
	if err := e.validatePutRuntimePlatformParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRuntimePlatform",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinition) PutVolume(value interface{}) {
	if err := e.validatePutVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVolume",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetCpu() {
	_jsii_.InvokeVoid(
		e,
		"resetCpu",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetEnableFaultInjection() {
	_jsii_.InvokeVoid(
		e,
		"resetEnableFaultInjection",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetEphemeralStorage() {
	_jsii_.InvokeVoid(
		e,
		"resetEphemeralStorage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetExecutionRoleArn() {
	_jsii_.InvokeVoid(
		e,
		"resetExecutionRoleArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetIpcMode() {
	_jsii_.InvokeVoid(
		e,
		"resetIpcMode",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetMemory() {
	_jsii_.InvokeVoid(
		e,
		"resetMemory",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetNetworkMode() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkMode",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetPidMode() {
	_jsii_.InvokeVoid(
		e,
		"resetPidMode",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetPlacementConstraints() {
	_jsii_.InvokeVoid(
		e,
		"resetPlacementConstraints",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetProxyConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetProxyConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetRegion() {
	_jsii_.InvokeVoid(
		e,
		"resetRegion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetRequiresCompatibilities() {
	_jsii_.InvokeVoid(
		e,
		"resetRequiresCompatibilities",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetRuntimePlatform() {
	_jsii_.InvokeVoid(
		e,
		"resetRuntimePlatform",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetSkipDestroy() {
	_jsii_.InvokeVoid(
		e,
		"resetSkipDestroy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetTagsAll() {
	_jsii_.InvokeVoid(
		e,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetTaskRoleArn() {
	_jsii_.InvokeVoid(
		e,
		"resetTaskRoleArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetTrackLatest() {
	_jsii_.InvokeVoid(
		e,
		"resetTrackLatest",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) ResetVolume() {
	_jsii_.InvokeVoid(
		e,
		"resetVolume",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinition) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinition) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinition) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinition) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinition) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

