// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdaeventsourcemapping

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/lambdaeventsourcemapping/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/lambda_event_source_mapping aws_lambda_event_source_mapping}.
type LambdaEventSourceMapping interface {
	cdktf.TerraformResource
	AmazonManagedKafkaEventSourceConfig() LambdaEventSourceMappingAmazonManagedKafkaEventSourceConfigOutputReference
	AmazonManagedKafkaEventSourceConfigInput() *LambdaEventSourceMappingAmazonManagedKafkaEventSourceConfig
	BatchSize() *float64
	SetBatchSize(val *float64)
	BatchSizeInput() *float64
	BisectBatchOnFunctionError() interface{}
	SetBisectBatchOnFunctionError(val interface{})
	BisectBatchOnFunctionErrorInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
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
	DestinationConfig() LambdaEventSourceMappingDestinationConfigOutputReference
	DestinationConfigInput() *LambdaEventSourceMappingDestinationConfig
	DocumentDbEventSourceConfig() LambdaEventSourceMappingDocumentDbEventSourceConfigOutputReference
	DocumentDbEventSourceConfigInput() *LambdaEventSourceMappingDocumentDbEventSourceConfig
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EventSourceArn() *string
	SetEventSourceArn(val *string)
	EventSourceArnInput() *string
	FilterCriteria() LambdaEventSourceMappingFilterCriteriaOutputReference
	FilterCriteriaInput() *LambdaEventSourceMappingFilterCriteria
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FunctionArn() *string
	FunctionName() *string
	SetFunctionName(val *string)
	FunctionNameInput() *string
	FunctionResponseTypes() *[]*string
	SetFunctionResponseTypes(val *[]*string)
	FunctionResponseTypesInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	LastModified() *string
	LastProcessingResult() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaximumBatchingWindowInSeconds() *float64
	SetMaximumBatchingWindowInSeconds(val *float64)
	MaximumBatchingWindowInSecondsInput() *float64
	MaximumRecordAgeInSeconds() *float64
	SetMaximumRecordAgeInSeconds(val *float64)
	MaximumRecordAgeInSecondsInput() *float64
	MaximumRetryAttempts() *float64
	SetMaximumRetryAttempts(val *float64)
	MaximumRetryAttemptsInput() *float64
	// The tree node.
	Node() constructs.Node
	ParallelizationFactor() *float64
	SetParallelizationFactor(val *float64)
	ParallelizationFactorInput() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Queues() *[]*string
	SetQueues(val *[]*string)
	QueuesInput() *[]*string
	// Experimental.
	RawOverrides() interface{}
	ScalingConfig() LambdaEventSourceMappingScalingConfigOutputReference
	ScalingConfigInput() *LambdaEventSourceMappingScalingConfig
	SelfManagedEventSource() LambdaEventSourceMappingSelfManagedEventSourceOutputReference
	SelfManagedEventSourceInput() *LambdaEventSourceMappingSelfManagedEventSource
	SelfManagedKafkaEventSourceConfig() LambdaEventSourceMappingSelfManagedKafkaEventSourceConfigOutputReference
	SelfManagedKafkaEventSourceConfigInput() *LambdaEventSourceMappingSelfManagedKafkaEventSourceConfig
	SourceAccessConfiguration() LambdaEventSourceMappingSourceAccessConfigurationList
	SourceAccessConfigurationInput() interface{}
	StartingPosition() *string
	SetStartingPosition(val *string)
	StartingPositionInput() *string
	StartingPositionTimestamp() *string
	SetStartingPositionTimestamp(val *string)
	StartingPositionTimestampInput() *string
	State() *string
	StateTransitionReason() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Topics() *[]*string
	SetTopics(val *[]*string)
	TopicsInput() *[]*string
	TumblingWindowInSeconds() *float64
	SetTumblingWindowInSeconds(val *float64)
	TumblingWindowInSecondsInput() *float64
	Uuid() *string
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAmazonManagedKafkaEventSourceConfig(value *LambdaEventSourceMappingAmazonManagedKafkaEventSourceConfig)
	PutDestinationConfig(value *LambdaEventSourceMappingDestinationConfig)
	PutDocumentDbEventSourceConfig(value *LambdaEventSourceMappingDocumentDbEventSourceConfig)
	PutFilterCriteria(value *LambdaEventSourceMappingFilterCriteria)
	PutScalingConfig(value *LambdaEventSourceMappingScalingConfig)
	PutSelfManagedEventSource(value *LambdaEventSourceMappingSelfManagedEventSource)
	PutSelfManagedKafkaEventSourceConfig(value *LambdaEventSourceMappingSelfManagedKafkaEventSourceConfig)
	PutSourceAccessConfiguration(value interface{})
	ResetAmazonManagedKafkaEventSourceConfig()
	ResetBatchSize()
	ResetBisectBatchOnFunctionError()
	ResetDestinationConfig()
	ResetDocumentDbEventSourceConfig()
	ResetEnabled()
	ResetEventSourceArn()
	ResetFilterCriteria()
	ResetFunctionResponseTypes()
	ResetId()
	ResetMaximumBatchingWindowInSeconds()
	ResetMaximumRecordAgeInSeconds()
	ResetMaximumRetryAttempts()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParallelizationFactor()
	ResetQueues()
	ResetScalingConfig()
	ResetSelfManagedEventSource()
	ResetSelfManagedKafkaEventSourceConfig()
	ResetSourceAccessConfiguration()
	ResetStartingPosition()
	ResetStartingPositionTimestamp()
	ResetTopics()
	ResetTumblingWindowInSeconds()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for LambdaEventSourceMapping
type jsiiProxy_LambdaEventSourceMapping struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LambdaEventSourceMapping) AmazonManagedKafkaEventSourceConfig() LambdaEventSourceMappingAmazonManagedKafkaEventSourceConfigOutputReference {
	var returns LambdaEventSourceMappingAmazonManagedKafkaEventSourceConfigOutputReference
	_jsii_.Get(
		j,
		"amazonManagedKafkaEventSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) AmazonManagedKafkaEventSourceConfigInput() *LambdaEventSourceMappingAmazonManagedKafkaEventSourceConfig {
	var returns *LambdaEventSourceMappingAmazonManagedKafkaEventSourceConfig
	_jsii_.Get(
		j,
		"amazonManagedKafkaEventSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) BatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) BatchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) BisectBatchOnFunctionError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bisectBatchOnFunctionError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) BisectBatchOnFunctionErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bisectBatchOnFunctionErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) DestinationConfig() LambdaEventSourceMappingDestinationConfigOutputReference {
	var returns LambdaEventSourceMappingDestinationConfigOutputReference
	_jsii_.Get(
		j,
		"destinationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) DestinationConfigInput() *LambdaEventSourceMappingDestinationConfig {
	var returns *LambdaEventSourceMappingDestinationConfig
	_jsii_.Get(
		j,
		"destinationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) DocumentDbEventSourceConfig() LambdaEventSourceMappingDocumentDbEventSourceConfigOutputReference {
	var returns LambdaEventSourceMappingDocumentDbEventSourceConfigOutputReference
	_jsii_.Get(
		j,
		"documentDbEventSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) DocumentDbEventSourceConfigInput() *LambdaEventSourceMappingDocumentDbEventSourceConfig {
	var returns *LambdaEventSourceMappingDocumentDbEventSourceConfig
	_jsii_.Get(
		j,
		"documentDbEventSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) EventSourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) EventSourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) FilterCriteria() LambdaEventSourceMappingFilterCriteriaOutputReference {
	var returns LambdaEventSourceMappingFilterCriteriaOutputReference
	_jsii_.Get(
		j,
		"filterCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) FilterCriteriaInput() *LambdaEventSourceMappingFilterCriteria {
	var returns *LambdaEventSourceMappingFilterCriteria
	_jsii_.Get(
		j,
		"filterCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) FunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) FunctionResponseTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"functionResponseTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) FunctionResponseTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"functionResponseTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) LastModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) LastProcessingResult() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastProcessingResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) MaximumBatchingWindowInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) MaximumBatchingWindowInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) MaximumRecordAgeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRecordAgeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) MaximumRecordAgeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRecordAgeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) MaximumRetryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) MaximumRetryAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) ParallelizationFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelizationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) ParallelizationFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelizationFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Queues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) QueuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) ScalingConfig() LambdaEventSourceMappingScalingConfigOutputReference {
	var returns LambdaEventSourceMappingScalingConfigOutputReference
	_jsii_.Get(
		j,
		"scalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) ScalingConfigInput() *LambdaEventSourceMappingScalingConfig {
	var returns *LambdaEventSourceMappingScalingConfig
	_jsii_.Get(
		j,
		"scalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) SelfManagedEventSource() LambdaEventSourceMappingSelfManagedEventSourceOutputReference {
	var returns LambdaEventSourceMappingSelfManagedEventSourceOutputReference
	_jsii_.Get(
		j,
		"selfManagedEventSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) SelfManagedEventSourceInput() *LambdaEventSourceMappingSelfManagedEventSource {
	var returns *LambdaEventSourceMappingSelfManagedEventSource
	_jsii_.Get(
		j,
		"selfManagedEventSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) SelfManagedKafkaEventSourceConfig() LambdaEventSourceMappingSelfManagedKafkaEventSourceConfigOutputReference {
	var returns LambdaEventSourceMappingSelfManagedKafkaEventSourceConfigOutputReference
	_jsii_.Get(
		j,
		"selfManagedKafkaEventSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) SelfManagedKafkaEventSourceConfigInput() *LambdaEventSourceMappingSelfManagedKafkaEventSourceConfig {
	var returns *LambdaEventSourceMappingSelfManagedKafkaEventSourceConfig
	_jsii_.Get(
		j,
		"selfManagedKafkaEventSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) SourceAccessConfiguration() LambdaEventSourceMappingSourceAccessConfigurationList {
	var returns LambdaEventSourceMappingSourceAccessConfigurationList
	_jsii_.Get(
		j,
		"sourceAccessConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) SourceAccessConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceAccessConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) StartingPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) StartingPositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) StartingPositionTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) StartingPositionTimestampInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionTimestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) StateTransitionReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateTransitionReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Topics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) TopicsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) TumblingWindowInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tumblingWindowInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) TumblingWindowInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tumblingWindowInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaEventSourceMapping) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/lambda_event_source_mapping aws_lambda_event_source_mapping} Resource.
func NewLambdaEventSourceMapping(scope constructs.Construct, id *string, config *LambdaEventSourceMappingConfig) LambdaEventSourceMapping {
	_init_.Initialize()

	if err := validateNewLambdaEventSourceMappingParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaEventSourceMapping{}

	_jsii_.Create(
		"@cdktf/provider-aws.lambdaEventSourceMapping.LambdaEventSourceMapping",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/lambda_event_source_mapping aws_lambda_event_source_mapping} Resource.
func NewLambdaEventSourceMapping_Override(l LambdaEventSourceMapping, scope constructs.Construct, id *string, config *LambdaEventSourceMappingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lambdaEventSourceMapping.LambdaEventSourceMapping",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetBatchSize(val *float64) {
	if err := j.validateSetBatchSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSize",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetBisectBatchOnFunctionError(val interface{}) {
	if err := j.validateSetBisectBatchOnFunctionErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bisectBatchOnFunctionError",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetEventSourceArn(val *string) {
	if err := j.validateSetEventSourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventSourceArn",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetFunctionName(val *string) {
	if err := j.validateSetFunctionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetFunctionResponseTypes(val *[]*string) {
	if err := j.validateSetFunctionResponseTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionResponseTypes",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetMaximumBatchingWindowInSeconds(val *float64) {
	if err := j.validateSetMaximumBatchingWindowInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBatchingWindowInSeconds",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetMaximumRecordAgeInSeconds(val *float64) {
	if err := j.validateSetMaximumRecordAgeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumRecordAgeInSeconds",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetMaximumRetryAttempts(val *float64) {
	if err := j.validateSetMaximumRetryAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumRetryAttempts",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetParallelizationFactor(val *float64) {
	if err := j.validateSetParallelizationFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parallelizationFactor",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetQueues(val *[]*string) {
	if err := j.validateSetQueuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queues",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetStartingPosition(val *string) {
	if err := j.validateSetStartingPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingPosition",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetStartingPositionTimestamp(val *string) {
	if err := j.validateSetStartingPositionTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingPositionTimestamp",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetTopics(val *[]*string) {
	if err := j.validateSetTopicsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topics",
		val,
	)
}

func (j *jsiiProxy_LambdaEventSourceMapping)SetTumblingWindowInSeconds(val *float64) {
	if err := j.validateSetTumblingWindowInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tumblingWindowInSeconds",
		val,
	)
}

// Generates CDKTF code for importing a LambdaEventSourceMapping resource upon running "cdktf plan <stack-name>".
func LambdaEventSourceMapping_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLambdaEventSourceMapping_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lambdaEventSourceMapping.LambdaEventSourceMapping",
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
func LambdaEventSourceMapping_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaEventSourceMapping_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lambdaEventSourceMapping.LambdaEventSourceMapping",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LambdaEventSourceMapping_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaEventSourceMapping_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lambdaEventSourceMapping.LambdaEventSourceMapping",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LambdaEventSourceMapping_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaEventSourceMapping_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lambdaEventSourceMapping.LambdaEventSourceMapping",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LambdaEventSourceMapping_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.lambdaEventSourceMapping.LambdaEventSourceMapping",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LambdaEventSourceMapping) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaEventSourceMapping) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaEventSourceMapping) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaEventSourceMapping) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaEventSourceMapping) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaEventSourceMapping) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaEventSourceMapping) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaEventSourceMapping) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaEventSourceMapping) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaEventSourceMapping) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaEventSourceMapping) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) PutAmazonManagedKafkaEventSourceConfig(value *LambdaEventSourceMappingAmazonManagedKafkaEventSourceConfig) {
	if err := l.validatePutAmazonManagedKafkaEventSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAmazonManagedKafkaEventSourceConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) PutDestinationConfig(value *LambdaEventSourceMappingDestinationConfig) {
	if err := l.validatePutDestinationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDestinationConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) PutDocumentDbEventSourceConfig(value *LambdaEventSourceMappingDocumentDbEventSourceConfig) {
	if err := l.validatePutDocumentDbEventSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDocumentDbEventSourceConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) PutFilterCriteria(value *LambdaEventSourceMappingFilterCriteria) {
	if err := l.validatePutFilterCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFilterCriteria",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) PutScalingConfig(value *LambdaEventSourceMappingScalingConfig) {
	if err := l.validatePutScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putScalingConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) PutSelfManagedEventSource(value *LambdaEventSourceMappingSelfManagedEventSource) {
	if err := l.validatePutSelfManagedEventSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSelfManagedEventSource",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) PutSelfManagedKafkaEventSourceConfig(value *LambdaEventSourceMappingSelfManagedKafkaEventSourceConfig) {
	if err := l.validatePutSelfManagedKafkaEventSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSelfManagedKafkaEventSourceConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) PutSourceAccessConfiguration(value interface{}) {
	if err := l.validatePutSourceAccessConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSourceAccessConfiguration",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetAmazonManagedKafkaEventSourceConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetAmazonManagedKafkaEventSourceConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetBatchSize() {
	_jsii_.InvokeVoid(
		l,
		"resetBatchSize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetBisectBatchOnFunctionError() {
	_jsii_.InvokeVoid(
		l,
		"resetBisectBatchOnFunctionError",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetDestinationConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetDestinationConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetDocumentDbEventSourceConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetDocumentDbEventSourceConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetEventSourceArn() {
	_jsii_.InvokeVoid(
		l,
		"resetEventSourceArn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetFilterCriteria() {
	_jsii_.InvokeVoid(
		l,
		"resetFilterCriteria",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetFunctionResponseTypes() {
	_jsii_.InvokeVoid(
		l,
		"resetFunctionResponseTypes",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetMaximumBatchingWindowInSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetMaximumBatchingWindowInSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetMaximumRecordAgeInSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetMaximumRecordAgeInSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetMaximumRetryAttempts() {
	_jsii_.InvokeVoid(
		l,
		"resetMaximumRetryAttempts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetParallelizationFactor() {
	_jsii_.InvokeVoid(
		l,
		"resetParallelizationFactor",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetQueues() {
	_jsii_.InvokeVoid(
		l,
		"resetQueues",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetScalingConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetScalingConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetSelfManagedEventSource() {
	_jsii_.InvokeVoid(
		l,
		"resetSelfManagedEventSource",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetSelfManagedKafkaEventSourceConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetSelfManagedKafkaEventSourceConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetSourceAccessConfiguration() {
	_jsii_.InvokeVoid(
		l,
		"resetSourceAccessConfiguration",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetStartingPosition() {
	_jsii_.InvokeVoid(
		l,
		"resetStartingPosition",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetStartingPositionTimestamp() {
	_jsii_.InvokeVoid(
		l,
		"resetStartingPositionTimestamp",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetTopics() {
	_jsii_.InvokeVoid(
		l,
		"resetTopics",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) ResetTumblingWindowInSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetTumblingWindowInSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaEventSourceMapping) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaEventSourceMapping) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaEventSourceMapping) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaEventSourceMapping) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

