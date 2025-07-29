// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iottopicrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/iottopicrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/iot_topic_rule aws_iot_topic_rule}.
type IotTopicRule interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudwatchAlarm() IotTopicRuleCloudwatchAlarmList
	CloudwatchAlarmInput() interface{}
	CloudwatchLogs() IotTopicRuleCloudwatchLogsList
	CloudwatchLogsInput() interface{}
	CloudwatchMetric() IotTopicRuleCloudwatchMetricList
	CloudwatchMetricInput() interface{}
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Dynamodb() IotTopicRuleDynamodbList
	DynamodbInput() interface{}
	Dynamodbv2() IotTopicRuleDynamodbv2List
	Dynamodbv2Input() interface{}
	Elasticsearch() IotTopicRuleElasticsearchList
	ElasticsearchInput() interface{}
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	ErrorAction() IotTopicRuleErrorActionOutputReference
	ErrorActionInput() *IotTopicRuleErrorAction
	Firehose() IotTopicRuleFirehoseList
	FirehoseInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Http() IotTopicRuleHttpList
	HttpInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	IotAnalytics() IotTopicRuleIotAnalyticsList
	IotAnalyticsInput() interface{}
	IotEvents() IotTopicRuleIotEventsList
	IotEventsInput() interface{}
	Kafka() IotTopicRuleKafkaList
	KafkaInput() interface{}
	Kinesis() IotTopicRuleKinesisList
	KinesisInput() interface{}
	Lambda() IotTopicRuleLambdaList
	LambdaInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	Republish() IotTopicRuleRepublishList
	RepublishInput() interface{}
	S3() IotTopicRuleS3List
	S3Input() interface{}
	Sns() IotTopicRuleSnsList
	SnsInput() interface{}
	Sql() *string
	SetSql(val *string)
	SqlInput() *string
	SqlVersion() *string
	SetSqlVersion(val *string)
	SqlVersionInput() *string
	Sqs() IotTopicRuleSqsList
	SqsInput() interface{}
	StepFunctions() IotTopicRuleStepFunctionsList
	StepFunctionsInput() interface{}
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
	Timestream() IotTopicRuleTimestreamList
	TimestreamInput() interface{}
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
	PutCloudwatchAlarm(value interface{})
	PutCloudwatchLogs(value interface{})
	PutCloudwatchMetric(value interface{})
	PutDynamodb(value interface{})
	PutDynamodbv2(value interface{})
	PutElasticsearch(value interface{})
	PutErrorAction(value *IotTopicRuleErrorAction)
	PutFirehose(value interface{})
	PutHttp(value interface{})
	PutIotAnalytics(value interface{})
	PutIotEvents(value interface{})
	PutKafka(value interface{})
	PutKinesis(value interface{})
	PutLambda(value interface{})
	PutRepublish(value interface{})
	PutS3(value interface{})
	PutSns(value interface{})
	PutSqs(value interface{})
	PutStepFunctions(value interface{})
	PutTimestream(value interface{})
	ResetCloudwatchAlarm()
	ResetCloudwatchLogs()
	ResetCloudwatchMetric()
	ResetDescription()
	ResetDynamodb()
	ResetDynamodbv2()
	ResetElasticsearch()
	ResetErrorAction()
	ResetFirehose()
	ResetHttp()
	ResetId()
	ResetIotAnalytics()
	ResetIotEvents()
	ResetKafka()
	ResetKinesis()
	ResetLambda()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetRepublish()
	ResetS3()
	ResetSns()
	ResetSqs()
	ResetStepFunctions()
	ResetTags()
	ResetTagsAll()
	ResetTimestream()
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

// The jsii proxy struct for IotTopicRule
type jsiiProxy_IotTopicRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IotTopicRule) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) CloudwatchAlarm() IotTopicRuleCloudwatchAlarmList {
	var returns IotTopicRuleCloudwatchAlarmList
	_jsii_.Get(
		j,
		"cloudwatchAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) CloudwatchAlarmInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchAlarmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) CloudwatchLogs() IotTopicRuleCloudwatchLogsList {
	var returns IotTopicRuleCloudwatchLogsList
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) CloudwatchLogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) CloudwatchMetric() IotTopicRuleCloudwatchMetricList {
	var returns IotTopicRuleCloudwatchMetricList
	_jsii_.Get(
		j,
		"cloudwatchMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) CloudwatchMetricInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Dynamodb() IotTopicRuleDynamodbList {
	var returns IotTopicRuleDynamodbList
	_jsii_.Get(
		j,
		"dynamodb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) DynamodbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamodbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Dynamodbv2() IotTopicRuleDynamodbv2List {
	var returns IotTopicRuleDynamodbv2List
	_jsii_.Get(
		j,
		"dynamodbv2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Dynamodbv2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamodbv2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Elasticsearch() IotTopicRuleElasticsearchList {
	var returns IotTopicRuleElasticsearchList
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) ElasticsearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) ErrorAction() IotTopicRuleErrorActionOutputReference {
	var returns IotTopicRuleErrorActionOutputReference
	_jsii_.Get(
		j,
		"errorAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) ErrorActionInput() *IotTopicRuleErrorAction {
	var returns *IotTopicRuleErrorAction
	_jsii_.Get(
		j,
		"errorActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Firehose() IotTopicRuleFirehoseList {
	var returns IotTopicRuleFirehoseList
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) FirehoseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Http() IotTopicRuleHttpList {
	var returns IotTopicRuleHttpList
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) HttpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) IotAnalytics() IotTopicRuleIotAnalyticsList {
	var returns IotTopicRuleIotAnalyticsList
	_jsii_.Get(
		j,
		"iotAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) IotAnalyticsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) IotEvents() IotTopicRuleIotEventsList {
	var returns IotTopicRuleIotEventsList
	_jsii_.Get(
		j,
		"iotEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) IotEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Kafka() IotTopicRuleKafkaList {
	var returns IotTopicRuleKafkaList
	_jsii_.Get(
		j,
		"kafka",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) KafkaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Kinesis() IotTopicRuleKinesisList {
	var returns IotTopicRuleKinesisList
	_jsii_.Get(
		j,
		"kinesis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) KinesisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Lambda() IotTopicRuleLambdaList {
	var returns IotTopicRuleLambdaList
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) LambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Republish() IotTopicRuleRepublishList {
	var returns IotTopicRuleRepublishList
	_jsii_.Get(
		j,
		"republish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) RepublishInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"republishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) S3() IotTopicRuleS3List {
	var returns IotTopicRuleS3List
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) S3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Sns() IotTopicRuleSnsList {
	var returns IotTopicRuleSnsList
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) SnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Sql() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) SqlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) SqlVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) SqlVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Sqs() IotTopicRuleSqsList {
	var returns IotTopicRuleSqsList
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) SqsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) StepFunctions() IotTopicRuleStepFunctionsList {
	var returns IotTopicRuleStepFunctionsList
	_jsii_.Get(
		j,
		"stepFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) StepFunctionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepFunctionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) Timestream() IotTopicRuleTimestreamList {
	var returns IotTopicRuleTimestreamList
	_jsii_.Get(
		j,
		"timestream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRule) TimestreamInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timestreamInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/iot_topic_rule aws_iot_topic_rule} Resource.
func NewIotTopicRule(scope constructs.Construct, id *string, config *IotTopicRuleConfig) IotTopicRule {
	_init_.Initialize()

	if err := validateNewIotTopicRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotTopicRule{}

	_jsii_.Create(
		"@cdktf/provider-aws.iotTopicRule.IotTopicRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/iot_topic_rule aws_iot_topic_rule} Resource.
func NewIotTopicRule_Override(i IotTopicRule, scope constructs.Construct, id *string, config *IotTopicRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.iotTopicRule.IotTopicRule",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotTopicRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule)SetSql(val *string) {
	if err := j.validateSetSqlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sql",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule)SetSqlVersion(val *string) {
	if err := j.validateSetSqlVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlVersion",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_IotTopicRule)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a IotTopicRule resource upon running "cdktf plan <stack-name>".
func IotTopicRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIotTopicRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.iotTopicRule.IotTopicRule",
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
func IotTopicRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotTopicRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.iotTopicRule.IotTopicRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotTopicRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotTopicRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.iotTopicRule.IotTopicRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotTopicRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotTopicRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.iotTopicRule.IotTopicRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotTopicRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.iotTopicRule.IotTopicRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IotTopicRule) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IotTopicRule) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IotTopicRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRule) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IotTopicRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRule) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IotTopicRule) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IotTopicRule) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IotTopicRule) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IotTopicRule) PutCloudwatchAlarm(value interface{}) {
	if err := i.validatePutCloudwatchAlarmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCloudwatchAlarm",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) PutCloudwatchLogs(value interface{}) {
	if err := i.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) PutCloudwatchMetric(value interface{}) {
	if err := i.validatePutCloudwatchMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCloudwatchMetric",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) PutDynamodb(value interface{}) {
	if err := i.validatePutDynamodbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDynamodb",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) PutDynamodbv2(value interface{}) {
	if err := i.validatePutDynamodbv2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDynamodbv2",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) PutElasticsearch(value interface{}) {
	if err := i.validatePutElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putElasticsearch",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) PutErrorAction(value *IotTopicRuleErrorAction) {
	if err := i.validatePutErrorActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putErrorAction",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) PutFirehose(value interface{}) {
	if err := i.validatePutFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFirehose",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) PutHttp(value interface{}) {
	if err := i.validatePutHttpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putHttp",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) PutIotAnalytics(value interface{}) {
	if err := i.validatePutIotAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotAnalytics",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) PutIotEvents(value interface{}) {
	if err := i.validatePutIotEventsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotEvents",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) PutKafka(value interface{}) {
	if err := i.validatePutKafkaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putKafka",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) PutKinesis(value interface{}) {
	if err := i.validatePutKinesisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putKinesis",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) PutLambda(value interface{}) {
	if err := i.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLambda",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) PutRepublish(value interface{}) {
	if err := i.validatePutRepublishParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putRepublish",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) PutS3(value interface{}) {
	if err := i.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putS3",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) PutSns(value interface{}) {
	if err := i.validatePutSnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSns",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) PutSqs(value interface{}) {
	if err := i.validatePutSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSqs",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) PutStepFunctions(value interface{}) {
	if err := i.validatePutStepFunctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putStepFunctions",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) PutTimestream(value interface{}) {
	if err := i.validatePutTimestreamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimestream",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRule) ResetCloudwatchAlarm() {
	_jsii_.InvokeVoid(
		i,
		"resetCloudwatchAlarm",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		i,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetCloudwatchMetric() {
	_jsii_.InvokeVoid(
		i,
		"resetCloudwatchMetric",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetDescription() {
	_jsii_.InvokeVoid(
		i,
		"resetDescription",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetDynamodb() {
	_jsii_.InvokeVoid(
		i,
		"resetDynamodb",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetDynamodbv2() {
	_jsii_.InvokeVoid(
		i,
		"resetDynamodbv2",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetElasticsearch() {
	_jsii_.InvokeVoid(
		i,
		"resetElasticsearch",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetErrorAction() {
	_jsii_.InvokeVoid(
		i,
		"resetErrorAction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetFirehose() {
	_jsii_.InvokeVoid(
		i,
		"resetFirehose",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetHttp() {
	_jsii_.InvokeVoid(
		i,
		"resetHttp",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetIotAnalytics() {
	_jsii_.InvokeVoid(
		i,
		"resetIotAnalytics",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetIotEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetIotEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetKafka() {
	_jsii_.InvokeVoid(
		i,
		"resetKafka",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetKinesis() {
	_jsii_.InvokeVoid(
		i,
		"resetKinesis",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetLambda() {
	_jsii_.InvokeVoid(
		i,
		"resetLambda",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetRegion() {
	_jsii_.InvokeVoid(
		i,
		"resetRegion",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetRepublish() {
	_jsii_.InvokeVoid(
		i,
		"resetRepublish",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetS3() {
	_jsii_.InvokeVoid(
		i,
		"resetS3",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetSns() {
	_jsii_.InvokeVoid(
		i,
		"resetSns",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetSqs() {
	_jsii_.InvokeVoid(
		i,
		"resetSqs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetStepFunctions() {
	_jsii_.InvokeVoid(
		i,
		"resetStepFunctions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetTagsAll() {
	_jsii_.InvokeVoid(
		i,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) ResetTimestream() {
	_jsii_.InvokeVoid(
		i,
		"resetTimestream",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

