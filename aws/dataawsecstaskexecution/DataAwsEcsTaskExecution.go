// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsecstaskexecution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/dataawsecstaskexecution/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/ecs_task_execution aws_ecs_task_execution}.
type DataAwsEcsTaskExecution interface {
	cdktf.TerraformDataSource
	CapacityProviderStrategy() DataAwsEcsTaskExecutionCapacityProviderStrategyList
	CapacityProviderStrategyInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientToken() *string
	SetClientToken(val *string)
	ClientTokenInput() *string
	Cluster() *string
	SetCluster(val *string)
	ClusterInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DesiredCount() *float64
	SetDesiredCount(val *float64)
	DesiredCountInput() *float64
	EnableEcsManagedTags() interface{}
	SetEnableEcsManagedTags(val interface{})
	EnableEcsManagedTagsInput() interface{}
	EnableExecuteCommand() interface{}
	SetEnableExecuteCommand(val interface{})
	EnableExecuteCommandInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Group() *string
	SetGroup(val *string)
	GroupInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	LaunchType() *string
	SetLaunchType(val *string)
	LaunchTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	NetworkConfiguration() DataAwsEcsTaskExecutionNetworkConfigurationOutputReference
	NetworkConfigurationInput() *DataAwsEcsTaskExecutionNetworkConfiguration
	// The tree node.
	Node() constructs.Node
	Overrides() DataAwsEcsTaskExecutionOverridesOutputReference
	OverridesInput() *DataAwsEcsTaskExecutionOverrides
	PlacementConstraints() DataAwsEcsTaskExecutionPlacementConstraintsList
	PlacementConstraintsInput() interface{}
	PlacementStrategy() DataAwsEcsTaskExecutionPlacementStrategyList
	PlacementStrategyInput() interface{}
	PlatformVersion() *string
	SetPlatformVersion(val *string)
	PlatformVersionInput() *string
	PropagateTags() *string
	SetPropagateTags(val *string)
	PropagateTagsInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ReferenceId() *string
	SetReferenceId(val *string)
	ReferenceIdInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	StartedBy() *string
	SetStartedBy(val *string)
	StartedByInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TaskArns() *[]*string
	TaskDefinition() *string
	SetTaskDefinition(val *string)
	TaskDefinitionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCapacityProviderStrategy(value interface{})
	PutNetworkConfiguration(value *DataAwsEcsTaskExecutionNetworkConfiguration)
	PutOverrides(value *DataAwsEcsTaskExecutionOverrides)
	PutPlacementConstraints(value interface{})
	PutPlacementStrategy(value interface{})
	ResetCapacityProviderStrategy()
	ResetClientToken()
	ResetDesiredCount()
	ResetEnableEcsManagedTags()
	ResetEnableExecuteCommand()
	ResetGroup()
	ResetId()
	ResetLaunchType()
	ResetNetworkConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOverrides()
	ResetPlacementConstraints()
	ResetPlacementStrategy()
	ResetPlatformVersion()
	ResetPropagateTags()
	ResetReferenceId()
	ResetRegion()
	ResetStartedBy()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataAwsEcsTaskExecution
type jsiiProxy_DataAwsEcsTaskExecution struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) CapacityProviderStrategy() DataAwsEcsTaskExecutionCapacityProviderStrategyList {
	var returns DataAwsEcsTaskExecutionCapacityProviderStrategyList
	_jsii_.Get(
		j,
		"capacityProviderStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) CapacityProviderStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityProviderStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) ClientToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) ClientTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) Cluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) ClusterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) DesiredCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) EnableEcsManagedTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) EnableEcsManagedTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) EnableExecuteCommand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) EnableExecuteCommandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) Group() *string {
	var returns *string
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) GroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) LaunchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) LaunchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) NetworkConfiguration() DataAwsEcsTaskExecutionNetworkConfigurationOutputReference {
	var returns DataAwsEcsTaskExecutionNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) NetworkConfigurationInput() *DataAwsEcsTaskExecutionNetworkConfiguration {
	var returns *DataAwsEcsTaskExecutionNetworkConfiguration
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) Overrides() DataAwsEcsTaskExecutionOverridesOutputReference {
	var returns DataAwsEcsTaskExecutionOverridesOutputReference
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) OverridesInput() *DataAwsEcsTaskExecutionOverrides {
	var returns *DataAwsEcsTaskExecutionOverrides
	_jsii_.Get(
		j,
		"overridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) PlacementConstraints() DataAwsEcsTaskExecutionPlacementConstraintsList {
	var returns DataAwsEcsTaskExecutionPlacementConstraintsList
	_jsii_.Get(
		j,
		"placementConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) PlacementConstraintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) PlacementStrategy() DataAwsEcsTaskExecutionPlacementStrategyList {
	var returns DataAwsEcsTaskExecutionPlacementStrategyList
	_jsii_.Get(
		j,
		"placementStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) PlacementStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) PlatformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) PlatformVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) PropagateTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) PropagateTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) ReferenceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) ReferenceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) StartedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) StartedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) TaskArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"taskArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) TaskDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) TaskDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcsTaskExecution) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/ecs_task_execution aws_ecs_task_execution} Data Source.
func NewDataAwsEcsTaskExecution(scope constructs.Construct, id *string, config *DataAwsEcsTaskExecutionConfig) DataAwsEcsTaskExecution {
	_init_.Initialize()

	if err := validateNewDataAwsEcsTaskExecutionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsEcsTaskExecution{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsEcsTaskExecution.DataAwsEcsTaskExecution",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/ecs_task_execution aws_ecs_task_execution} Data Source.
func NewDataAwsEcsTaskExecution_Override(d DataAwsEcsTaskExecution, scope constructs.Construct, id *string, config *DataAwsEcsTaskExecutionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsEcsTaskExecution.DataAwsEcsTaskExecution",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecution)SetClientToken(val *string) {
	if err := j.validateSetClientTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientToken",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecution)SetCluster(val *string) {
	if err := j.validateSetClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cluster",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecution)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecution)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecution)SetDesiredCount(val *float64) {
	if err := j.validateSetDesiredCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCount",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecution)SetEnableEcsManagedTags(val interface{}) {
	if err := j.validateSetEnableEcsManagedTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableEcsManagedTags",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecution)SetEnableExecuteCommand(val interface{}) {
	if err := j.validateSetEnableExecuteCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableExecuteCommand",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecution)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecution)SetGroup(val *string) {
	if err := j.validateSetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"group",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecution)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecution)SetLaunchType(val *string) {
	if err := j.validateSetLaunchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchType",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecution)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecution)SetPlatformVersion(val *string) {
	if err := j.validateSetPlatformVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformVersion",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecution)SetPropagateTags(val *string) {
	if err := j.validateSetPropagateTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecution)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecution)SetReferenceId(val *string) {
	if err := j.validateSetReferenceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referenceId",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecution)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecution)SetStartedBy(val *string) {
	if err := j.validateSetStartedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startedBy",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecution)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcsTaskExecution)SetTaskDefinition(val *string) {
	if err := j.validateSetTaskDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskDefinition",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsEcsTaskExecution resource upon running "cdktf plan <stack-name>".
func DataAwsEcsTaskExecution_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsEcsTaskExecution_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsEcsTaskExecution.DataAwsEcsTaskExecution",
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
func DataAwsEcsTaskExecution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsEcsTaskExecution_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsEcsTaskExecution.DataAwsEcsTaskExecution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsEcsTaskExecution_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsEcsTaskExecution_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsEcsTaskExecution.DataAwsEcsTaskExecution",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsEcsTaskExecution_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsEcsTaskExecution_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsEcsTaskExecution.DataAwsEcsTaskExecution",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsEcsTaskExecution_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsEcsTaskExecution.DataAwsEcsTaskExecution",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsEcsTaskExecution) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsEcsTaskExecution) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsEcsTaskExecution) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsEcsTaskExecution) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsEcsTaskExecution) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsEcsTaskExecution) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsEcsTaskExecution) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsEcsTaskExecution) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsEcsTaskExecution) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) PutCapacityProviderStrategy(value interface{}) {
	if err := d.validatePutCapacityProviderStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCapacityProviderStrategy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) PutNetworkConfiguration(value *DataAwsEcsTaskExecutionNetworkConfiguration) {
	if err := d.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) PutOverrides(value *DataAwsEcsTaskExecutionOverrides) {
	if err := d.validatePutOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOverrides",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) PutPlacementConstraints(value interface{}) {
	if err := d.validatePutPlacementConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPlacementConstraints",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) PutPlacementStrategy(value interface{}) {
	if err := d.validatePutPlacementStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPlacementStrategy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ResetCapacityProviderStrategy() {
	_jsii_.InvokeVoid(
		d,
		"resetCapacityProviderStrategy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ResetClientToken() {
	_jsii_.InvokeVoid(
		d,
		"resetClientToken",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ResetDesiredCount() {
	_jsii_.InvokeVoid(
		d,
		"resetDesiredCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ResetEnableEcsManagedTags() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableEcsManagedTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ResetEnableExecuteCommand() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableExecuteCommand",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ResetGroup() {
	_jsii_.InvokeVoid(
		d,
		"resetGroup",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ResetLaunchType() {
	_jsii_.InvokeVoid(
		d,
		"resetLaunchType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ResetOverrides() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrides",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ResetPlacementConstraints() {
	_jsii_.InvokeVoid(
		d,
		"resetPlacementConstraints",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ResetPlacementStrategy() {
	_jsii_.InvokeVoid(
		d,
		"resetPlacementStrategy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ResetPlatformVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetPlatformVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ResetPropagateTags() {
	_jsii_.InvokeVoid(
		d,
		"resetPropagateTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ResetReferenceId() {
	_jsii_.InvokeVoid(
		d,
		"resetReferenceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ResetStartedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetStartedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcsTaskExecution) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

