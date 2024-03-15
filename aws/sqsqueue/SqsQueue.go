// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqsqueue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/sqsqueue/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/sqs_queue aws_sqs_queue}.
type SqsQueue interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContentBasedDeduplication() interface{}
	SetContentBasedDeduplication(val interface{})
	ContentBasedDeduplicationInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DeduplicationScope() *string
	SetDeduplicationScope(val *string)
	DeduplicationScopeInput() *string
	DelaySeconds() *float64
	SetDelaySeconds(val *float64)
	DelaySecondsInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FifoQueue() interface{}
	SetFifoQueue(val interface{})
	FifoQueueInput() interface{}
	FifoThroughputLimit() *string
	SetFifoThroughputLimit(val *string)
	FifoThroughputLimitInput() *string
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
	KmsDataKeyReusePeriodSeconds() *float64
	SetKmsDataKeyReusePeriodSeconds(val *float64)
	KmsDataKeyReusePeriodSecondsInput() *float64
	KmsMasterKeyId() *string
	SetKmsMasterKeyId(val *string)
	KmsMasterKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxMessageSize() *float64
	SetMaxMessageSize(val *float64)
	MaxMessageSizeInput() *float64
	MessageRetentionSeconds() *float64
	SetMessageRetentionSeconds(val *float64)
	MessageRetentionSecondsInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	// The tree node.
	Node() constructs.Node
	Policy() *string
	SetPolicy(val *string)
	PolicyInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReceiveWaitTimeSeconds() *float64
	SetReceiveWaitTimeSeconds(val *float64)
	ReceiveWaitTimeSecondsInput() *float64
	RedriveAllowPolicy() *string
	SetRedriveAllowPolicy(val *string)
	RedriveAllowPolicyInput() *string
	RedrivePolicy() *string
	SetRedrivePolicy(val *string)
	RedrivePolicyInput() *string
	SqsManagedSseEnabled() interface{}
	SetSqsManagedSseEnabled(val interface{})
	SqsManagedSseEnabledInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Url() *string
	VisibilityTimeoutSeconds() *float64
	SetVisibilityTimeoutSeconds(val *float64)
	VisibilityTimeoutSecondsInput() *float64
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
	ResetContentBasedDeduplication()
	ResetDeduplicationScope()
	ResetDelaySeconds()
	ResetFifoQueue()
	ResetFifoThroughputLimit()
	ResetId()
	ResetKmsDataKeyReusePeriodSeconds()
	ResetKmsMasterKeyId()
	ResetMaxMessageSize()
	ResetMessageRetentionSeconds()
	ResetName()
	ResetNamePrefix()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicy()
	ResetReceiveWaitTimeSeconds()
	ResetRedriveAllowPolicy()
	ResetRedrivePolicy()
	ResetSqsManagedSseEnabled()
	ResetTags()
	ResetTagsAll()
	ResetVisibilityTimeoutSeconds()
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

// The jsii proxy struct for SqsQueue
type jsiiProxy_SqsQueue struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SqsQueue) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) ContentBasedDeduplication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentBasedDeduplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) ContentBasedDeduplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentBasedDeduplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) DeduplicationScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deduplicationScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) DeduplicationScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deduplicationScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) DelaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"delaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) DelaySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"delaySecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) FifoQueue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fifoQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) FifoQueueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fifoQueueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) FifoThroughputLimit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fifoThroughputLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) FifoThroughputLimitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fifoThroughputLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) KmsDataKeyReusePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"kmsDataKeyReusePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) KmsDataKeyReusePeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"kmsDataKeyReusePeriodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) KmsMasterKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsMasterKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) KmsMasterKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsMasterKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) MaxMessageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMessageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) MaxMessageSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMessageSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) MessageRetentionSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messageRetentionSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) MessageRetentionSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messageRetentionSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) ReceiveWaitTimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"receiveWaitTimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) ReceiveWaitTimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"receiveWaitTimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) RedriveAllowPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redriveAllowPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) RedriveAllowPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redriveAllowPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) RedrivePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redrivePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) RedrivePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redrivePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) SqsManagedSseEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqsManagedSseEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) SqsManagedSseEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqsManagedSseEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) VisibilityTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"visibilityTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueue) VisibilityTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"visibilityTimeoutSecondsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/sqs_queue aws_sqs_queue} Resource.
func NewSqsQueue(scope constructs.Construct, id *string, config *SqsQueueConfig) SqsQueue {
	_init_.Initialize()

	if err := validateNewSqsQueueParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqsQueue{}

	_jsii_.Create(
		"@cdktf/provider-aws.sqsQueue.SqsQueue",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/sqs_queue aws_sqs_queue} Resource.
func NewSqsQueue_Override(s SqsQueue, scope constructs.Construct, id *string, config *SqsQueueConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sqsQueue.SqsQueue",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SqsQueue)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetContentBasedDeduplication(val interface{}) {
	if err := j.validateSetContentBasedDeduplicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentBasedDeduplication",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetDeduplicationScope(val *string) {
	if err := j.validateSetDeduplicationScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deduplicationScope",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetDelaySeconds(val *float64) {
	if err := j.validateSetDelaySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delaySeconds",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetFifoQueue(val interface{}) {
	if err := j.validateSetFifoQueueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fifoQueue",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetFifoThroughputLimit(val *string) {
	if err := j.validateSetFifoThroughputLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fifoThroughputLimit",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetKmsDataKeyReusePeriodSeconds(val *float64) {
	if err := j.validateSetKmsDataKeyReusePeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsDataKeyReusePeriodSeconds",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetKmsMasterKeyId(val *string) {
	if err := j.validateSetKmsMasterKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsMasterKeyId",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetMaxMessageSize(val *float64) {
	if err := j.validateSetMaxMessageSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxMessageSize",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetMessageRetentionSeconds(val *float64) {
	if err := j.validateSetMessageRetentionSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageRetentionSeconds",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetPolicy(val *string) {
	if err := j.validateSetPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetReceiveWaitTimeSeconds(val *float64) {
	if err := j.validateSetReceiveWaitTimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"receiveWaitTimeSeconds",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetRedriveAllowPolicy(val *string) {
	if err := j.validateSetRedriveAllowPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redriveAllowPolicy",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetRedrivePolicy(val *string) {
	if err := j.validateSetRedrivePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redrivePolicy",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetSqsManagedSseEnabled(val interface{}) {
	if err := j.validateSetSqsManagedSseEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqsManagedSseEnabled",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_SqsQueue)SetVisibilityTimeoutSeconds(val *float64) {
	if err := j.validateSetVisibilityTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibilityTimeoutSeconds",
		val,
	)
}

// Generates CDKTF code for importing a SqsQueue resource upon running "cdktf plan <stack-name>".
func SqsQueue_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSqsQueue_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sqsQueue.SqsQueue",
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
func SqsQueue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSqsQueue_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sqsQueue.SqsQueue",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SqsQueue_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSqsQueue_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sqsQueue.SqsQueue",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SqsQueue_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSqsQueue_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sqsQueue.SqsQueue",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SqsQueue_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.sqsQueue.SqsQueue",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SqsQueue) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SqsQueue) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SqsQueue) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SqsQueue) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SqsQueue) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SqsQueue) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SqsQueue) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SqsQueue) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SqsQueue) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SqsQueue) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SqsQueue) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SqsQueue) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueue) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SqsQueue) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueue) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SqsQueue) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SqsQueue) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SqsQueue) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SqsQueue) ResetContentBasedDeduplication() {
	_jsii_.InvokeVoid(
		s,
		"resetContentBasedDeduplication",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) ResetDeduplicationScope() {
	_jsii_.InvokeVoid(
		s,
		"resetDeduplicationScope",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) ResetDelaySeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetDelaySeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) ResetFifoQueue() {
	_jsii_.InvokeVoid(
		s,
		"resetFifoQueue",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) ResetFifoThroughputLimit() {
	_jsii_.InvokeVoid(
		s,
		"resetFifoThroughputLimit",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) ResetKmsDataKeyReusePeriodSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsDataKeyReusePeriodSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) ResetKmsMasterKeyId() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsMasterKeyId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) ResetMaxMessageSize() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxMessageSize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) ResetMessageRetentionSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetMessageRetentionSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) ResetPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) ResetReceiveWaitTimeSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetReceiveWaitTimeSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) ResetRedriveAllowPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetRedriveAllowPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) ResetRedrivePolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetRedrivePolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) ResetSqsManagedSseEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetSqsManagedSseEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) ResetVisibilityTimeoutSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetVisibilityTimeoutSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqsQueue) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueue) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueue) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueue) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueue) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueue) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

