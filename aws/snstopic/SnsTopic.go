// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package snstopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/snstopic/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/sns_topic aws_sns_topic}.
type SnsTopic interface {
	cdktf.TerraformResource
	ApplicationFailureFeedbackRoleArn() *string
	SetApplicationFailureFeedbackRoleArn(val *string)
	ApplicationFailureFeedbackRoleArnInput() *string
	ApplicationSuccessFeedbackRoleArn() *string
	SetApplicationSuccessFeedbackRoleArn(val *string)
	ApplicationSuccessFeedbackRoleArnInput() *string
	ApplicationSuccessFeedbackSampleRate() *float64
	SetApplicationSuccessFeedbackSampleRate(val *float64)
	ApplicationSuccessFeedbackSampleRateInput() *float64
	ArchivePolicy() *string
	SetArchivePolicy(val *string)
	ArchivePolicyInput() *string
	Arn() *string
	BeginningArchiveTime() *string
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
	DeliveryPolicy() *string
	SetDeliveryPolicy(val *string)
	DeliveryPolicyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	FifoTopic() interface{}
	SetFifoTopic(val interface{})
	FifoTopicInput() interface{}
	FirehoseFailureFeedbackRoleArn() *string
	SetFirehoseFailureFeedbackRoleArn(val *string)
	FirehoseFailureFeedbackRoleArnInput() *string
	FirehoseSuccessFeedbackRoleArn() *string
	SetFirehoseSuccessFeedbackRoleArn(val *string)
	FirehoseSuccessFeedbackRoleArnInput() *string
	FirehoseSuccessFeedbackSampleRate() *float64
	SetFirehoseSuccessFeedbackSampleRate(val *float64)
	FirehoseSuccessFeedbackSampleRateInput() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpFailureFeedbackRoleArn() *string
	SetHttpFailureFeedbackRoleArn(val *string)
	HttpFailureFeedbackRoleArnInput() *string
	HttpSuccessFeedbackRoleArn() *string
	SetHttpSuccessFeedbackRoleArn(val *string)
	HttpSuccessFeedbackRoleArnInput() *string
	HttpSuccessFeedbackSampleRate() *float64
	SetHttpSuccessFeedbackSampleRate(val *float64)
	HttpSuccessFeedbackSampleRateInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	KmsMasterKeyId() *string
	SetKmsMasterKeyId(val *string)
	KmsMasterKeyIdInput() *string
	LambdaFailureFeedbackRoleArn() *string
	SetLambdaFailureFeedbackRoleArn(val *string)
	LambdaFailureFeedbackRoleArnInput() *string
	LambdaSuccessFeedbackRoleArn() *string
	SetLambdaSuccessFeedbackRoleArn(val *string)
	LambdaSuccessFeedbackRoleArnInput() *string
	LambdaSuccessFeedbackSampleRate() *float64
	SetLambdaSuccessFeedbackSampleRate(val *float64)
	LambdaSuccessFeedbackSampleRateInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	// The tree node.
	Node() constructs.Node
	Owner() *string
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
	SignatureVersion() *float64
	SetSignatureVersion(val *float64)
	SignatureVersionInput() *float64
	SqsFailureFeedbackRoleArn() *string
	SetSqsFailureFeedbackRoleArn(val *string)
	SqsFailureFeedbackRoleArnInput() *string
	SqsSuccessFeedbackRoleArn() *string
	SetSqsSuccessFeedbackRoleArn(val *string)
	SqsSuccessFeedbackRoleArnInput() *string
	SqsSuccessFeedbackSampleRate() *float64
	SetSqsSuccessFeedbackSampleRate(val *float64)
	SqsSuccessFeedbackSampleRateInput() *float64
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
	TracingConfig() *string
	SetTracingConfig(val *string)
	TracingConfigInput() *string
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
	ResetApplicationFailureFeedbackRoleArn()
	ResetApplicationSuccessFeedbackRoleArn()
	ResetApplicationSuccessFeedbackSampleRate()
	ResetArchivePolicy()
	ResetContentBasedDeduplication()
	ResetDeliveryPolicy()
	ResetDisplayName()
	ResetFifoTopic()
	ResetFirehoseFailureFeedbackRoleArn()
	ResetFirehoseSuccessFeedbackRoleArn()
	ResetFirehoseSuccessFeedbackSampleRate()
	ResetHttpFailureFeedbackRoleArn()
	ResetHttpSuccessFeedbackRoleArn()
	ResetHttpSuccessFeedbackSampleRate()
	ResetId()
	ResetKmsMasterKeyId()
	ResetLambdaFailureFeedbackRoleArn()
	ResetLambdaSuccessFeedbackRoleArn()
	ResetLambdaSuccessFeedbackSampleRate()
	ResetName()
	ResetNamePrefix()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicy()
	ResetSignatureVersion()
	ResetSqsFailureFeedbackRoleArn()
	ResetSqsSuccessFeedbackRoleArn()
	ResetSqsSuccessFeedbackSampleRate()
	ResetTags()
	ResetTagsAll()
	ResetTracingConfig()
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

// The jsii proxy struct for SnsTopic
type jsiiProxy_SnsTopic struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SnsTopic) ApplicationFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ApplicationFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ApplicationSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ApplicationSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ApplicationSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"applicationSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ApplicationSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"applicationSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ArchivePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archivePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ArchivePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archivePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) BeginningArchiveTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beginningArchiveTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ContentBasedDeduplication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentBasedDeduplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ContentBasedDeduplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentBasedDeduplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) DeliveryPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) DeliveryPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FifoTopic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fifoTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FifoTopicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fifoTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FirehoseFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FirehoseFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FirehoseSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FirehoseSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FirehoseSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firehoseSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FirehoseSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firehoseSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) HttpFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) HttpFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) HttpSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) HttpSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) HttpSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) HttpSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) KmsMasterKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsMasterKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) KmsMasterKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsMasterKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) LambdaFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) LambdaFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) LambdaSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) LambdaSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) LambdaSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lambdaSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) LambdaSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lambdaSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) SignatureVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signatureVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) SignatureVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signatureVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) SqsFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqsFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) SqsFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqsFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) SqsSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqsSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) SqsSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqsSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) SqsSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sqsSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) SqsSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sqsSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TracingConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tracingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopic) TracingConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tracingConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/sns_topic aws_sns_topic} Resource.
func NewSnsTopic(scope constructs.Construct, id *string, config *SnsTopicConfig) SnsTopic {
	_init_.Initialize()

	if err := validateNewSnsTopicParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnsTopic{}

	_jsii_.Create(
		"@cdktf/provider-aws.snsTopic.SnsTopic",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/sns_topic aws_sns_topic} Resource.
func NewSnsTopic_Override(s SnsTopic, scope constructs.Construct, id *string, config *SnsTopicConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.snsTopic.SnsTopic",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SnsTopic)SetApplicationFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetApplicationFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetApplicationSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetApplicationSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetApplicationSuccessFeedbackSampleRate(val *float64) {
	if err := j.validateSetApplicationSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetArchivePolicy(val *string) {
	if err := j.validateSetArchivePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archivePolicy",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetContentBasedDeduplication(val interface{}) {
	if err := j.validateSetContentBasedDeduplicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentBasedDeduplication",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetDeliveryPolicy(val *string) {
	if err := j.validateSetDeliveryPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deliveryPolicy",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetFifoTopic(val interface{}) {
	if err := j.validateSetFifoTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fifoTopic",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetFirehoseFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetFirehoseFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firehoseFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetFirehoseSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetFirehoseSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firehoseSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetFirehoseSuccessFeedbackSampleRate(val *float64) {
	if err := j.validateSetFirehoseSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firehoseSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetHttpFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetHttpFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetHttpSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetHttpSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetHttpSuccessFeedbackSampleRate(val *float64) {
	if err := j.validateSetHttpSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetKmsMasterKeyId(val *string) {
	if err := j.validateSetKmsMasterKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsMasterKeyId",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetLambdaFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetLambdaFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetLambdaSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetLambdaSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetLambdaSuccessFeedbackSampleRate(val *float64) {
	if err := j.validateSetLambdaSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetPolicy(val *string) {
	if err := j.validateSetPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetSignatureVersion(val *float64) {
	if err := j.validateSetSignatureVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signatureVersion",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetSqsFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetSqsFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqsFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetSqsSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetSqsSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqsSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetSqsSuccessFeedbackSampleRate(val *float64) {
	if err := j.validateSetSqsSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqsSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_SnsTopic)SetTracingConfig(val *string) {
	if err := j.validateSetTracingConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tracingConfig",
		val,
	)
}

// Generates CDKTF code for importing a SnsTopic resource upon running "cdktf plan <stack-name>".
func SnsTopic_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSnsTopic_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.snsTopic.SnsTopic",
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
func SnsTopic_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSnsTopic_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.snsTopic.SnsTopic",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SnsTopic_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSnsTopic_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.snsTopic.SnsTopic",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SnsTopic_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSnsTopic_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.snsTopic.SnsTopic",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SnsTopic_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.snsTopic.SnsTopic",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SnsTopic) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SnsTopic) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SnsTopic) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SnsTopic) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SnsTopic) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SnsTopic) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SnsTopic) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SnsTopic) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SnsTopic) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SnsTopic) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SnsTopic) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SnsTopic) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopic) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SnsTopic) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SnsTopic) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SnsTopic) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SnsTopic) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SnsTopic) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SnsTopic) ResetApplicationFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetApplicationFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetApplicationSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetApplicationSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetApplicationSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		s,
		"resetApplicationSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetArchivePolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetArchivePolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetContentBasedDeduplication() {
	_jsii_.InvokeVoid(
		s,
		"resetContentBasedDeduplication",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetDeliveryPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetDeliveryPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetDisplayName() {
	_jsii_.InvokeVoid(
		s,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetFifoTopic() {
	_jsii_.InvokeVoid(
		s,
		"resetFifoTopic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetFirehoseFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetFirehoseFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetFirehoseSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetFirehoseSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetFirehoseSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		s,
		"resetFirehoseSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetHttpFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetHttpSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetHttpSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetKmsMasterKeyId() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsMasterKeyId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetLambdaFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetLambdaFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetLambdaSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetLambdaSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetLambdaSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		s,
		"resetLambdaSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetSignatureVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetSignatureVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetSqsFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetSqsFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetSqsSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetSqsSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetSqsSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		s,
		"resetSqsSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) ResetTracingConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetTracingConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopic) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopic) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopic) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopic) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopic) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopic) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

