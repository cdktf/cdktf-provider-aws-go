package iottopicrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/iottopicrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotTopicRuleErrorActionOutputReference interface {
	cdktf.ComplexObject
	CloudwatchAlarm() IotTopicRuleErrorActionCloudwatchAlarmOutputReference
	CloudwatchAlarmInput() *IotTopicRuleErrorActionCloudwatchAlarm
	CloudwatchLogs() IotTopicRuleErrorActionCloudwatchLogsOutputReference
	CloudwatchLogsInput() *IotTopicRuleErrorActionCloudwatchLogs
	CloudwatchMetric() IotTopicRuleErrorActionCloudwatchMetricOutputReference
	CloudwatchMetricInput() *IotTopicRuleErrorActionCloudwatchMetric
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
	Dynamodb() IotTopicRuleErrorActionDynamodbOutputReference
	DynamodbInput() *IotTopicRuleErrorActionDynamodb
	Dynamodbv2() IotTopicRuleErrorActionDynamodbv2OutputReference
	Dynamodbv2Input() *IotTopicRuleErrorActionDynamodbv2
	Elasticsearch() IotTopicRuleErrorActionElasticsearchOutputReference
	ElasticsearchInput() *IotTopicRuleErrorActionElasticsearch
	Firehose() IotTopicRuleErrorActionFirehoseOutputReference
	FirehoseInput() *IotTopicRuleErrorActionFirehose
	// Experimental.
	Fqn() *string
	Http() IotTopicRuleErrorActionHttpOutputReference
	HttpInput() *IotTopicRuleErrorActionHttp
	InternalValue() *IotTopicRuleErrorAction
	SetInternalValue(val *IotTopicRuleErrorAction)
	IotAnalytics() IotTopicRuleErrorActionIotAnalyticsOutputReference
	IotAnalyticsInput() *IotTopicRuleErrorActionIotAnalytics
	IotEvents() IotTopicRuleErrorActionIotEventsOutputReference
	IotEventsInput() *IotTopicRuleErrorActionIotEvents
	Kafka() IotTopicRuleErrorActionKafkaOutputReference
	KafkaInput() *IotTopicRuleErrorActionKafka
	Kinesis() IotTopicRuleErrorActionKinesisOutputReference
	KinesisInput() *IotTopicRuleErrorActionKinesis
	Lambda() IotTopicRuleErrorActionLambdaOutputReference
	LambdaInput() *IotTopicRuleErrorActionLambda
	Republish() IotTopicRuleErrorActionRepublishOutputReference
	RepublishInput() *IotTopicRuleErrorActionRepublish
	S3() IotTopicRuleErrorActionS3OutputReference
	S3Input() *IotTopicRuleErrorActionS3
	Sns() IotTopicRuleErrorActionSnsOutputReference
	SnsInput() *IotTopicRuleErrorActionSns
	Sqs() IotTopicRuleErrorActionSqsOutputReference
	SqsInput() *IotTopicRuleErrorActionSqs
	StepFunctions() IotTopicRuleErrorActionStepFunctionsOutputReference
	StepFunctionsInput() *IotTopicRuleErrorActionStepFunctions
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timestream() IotTopicRuleErrorActionTimestreamOutputReference
	TimestreamInput() *IotTopicRuleErrorActionTimestream
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
	PutCloudwatchAlarm(value *IotTopicRuleErrorActionCloudwatchAlarm)
	PutCloudwatchLogs(value *IotTopicRuleErrorActionCloudwatchLogs)
	PutCloudwatchMetric(value *IotTopicRuleErrorActionCloudwatchMetric)
	PutDynamodb(value *IotTopicRuleErrorActionDynamodb)
	PutDynamodbv2(value *IotTopicRuleErrorActionDynamodbv2)
	PutElasticsearch(value *IotTopicRuleErrorActionElasticsearch)
	PutFirehose(value *IotTopicRuleErrorActionFirehose)
	PutHttp(value *IotTopicRuleErrorActionHttp)
	PutIotAnalytics(value *IotTopicRuleErrorActionIotAnalytics)
	PutIotEvents(value *IotTopicRuleErrorActionIotEvents)
	PutKafka(value *IotTopicRuleErrorActionKafka)
	PutKinesis(value *IotTopicRuleErrorActionKinesis)
	PutLambda(value *IotTopicRuleErrorActionLambda)
	PutRepublish(value *IotTopicRuleErrorActionRepublish)
	PutS3(value *IotTopicRuleErrorActionS3)
	PutSns(value *IotTopicRuleErrorActionSns)
	PutSqs(value *IotTopicRuleErrorActionSqs)
	PutStepFunctions(value *IotTopicRuleErrorActionStepFunctions)
	PutTimestream(value *IotTopicRuleErrorActionTimestream)
	ResetCloudwatchAlarm()
	ResetCloudwatchLogs()
	ResetCloudwatchMetric()
	ResetDynamodb()
	ResetDynamodbv2()
	ResetElasticsearch()
	ResetFirehose()
	ResetHttp()
	ResetIotAnalytics()
	ResetIotEvents()
	ResetKafka()
	ResetKinesis()
	ResetLambda()
	ResetRepublish()
	ResetS3()
	ResetSns()
	ResetSqs()
	ResetStepFunctions()
	ResetTimestream()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotTopicRuleErrorActionOutputReference
type jsiiProxy_IotTopicRuleErrorActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) CloudwatchAlarm() IotTopicRuleErrorActionCloudwatchAlarmOutputReference {
	var returns IotTopicRuleErrorActionCloudwatchAlarmOutputReference
	_jsii_.Get(
		j,
		"cloudwatchAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) CloudwatchAlarmInput() *IotTopicRuleErrorActionCloudwatchAlarm {
	var returns *IotTopicRuleErrorActionCloudwatchAlarm
	_jsii_.Get(
		j,
		"cloudwatchAlarmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) CloudwatchLogs() IotTopicRuleErrorActionCloudwatchLogsOutputReference {
	var returns IotTopicRuleErrorActionCloudwatchLogsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) CloudwatchLogsInput() *IotTopicRuleErrorActionCloudwatchLogs {
	var returns *IotTopicRuleErrorActionCloudwatchLogs
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) CloudwatchMetric() IotTopicRuleErrorActionCloudwatchMetricOutputReference {
	var returns IotTopicRuleErrorActionCloudwatchMetricOutputReference
	_jsii_.Get(
		j,
		"cloudwatchMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) CloudwatchMetricInput() *IotTopicRuleErrorActionCloudwatchMetric {
	var returns *IotTopicRuleErrorActionCloudwatchMetric
	_jsii_.Get(
		j,
		"cloudwatchMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Dynamodb() IotTopicRuleErrorActionDynamodbOutputReference {
	var returns IotTopicRuleErrorActionDynamodbOutputReference
	_jsii_.Get(
		j,
		"dynamodb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) DynamodbInput() *IotTopicRuleErrorActionDynamodb {
	var returns *IotTopicRuleErrorActionDynamodb
	_jsii_.Get(
		j,
		"dynamodbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Dynamodbv2() IotTopicRuleErrorActionDynamodbv2OutputReference {
	var returns IotTopicRuleErrorActionDynamodbv2OutputReference
	_jsii_.Get(
		j,
		"dynamodbv2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Dynamodbv2Input() *IotTopicRuleErrorActionDynamodbv2 {
	var returns *IotTopicRuleErrorActionDynamodbv2
	_jsii_.Get(
		j,
		"dynamodbv2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Elasticsearch() IotTopicRuleErrorActionElasticsearchOutputReference {
	var returns IotTopicRuleErrorActionElasticsearchOutputReference
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) ElasticsearchInput() *IotTopicRuleErrorActionElasticsearch {
	var returns *IotTopicRuleErrorActionElasticsearch
	_jsii_.Get(
		j,
		"elasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Firehose() IotTopicRuleErrorActionFirehoseOutputReference {
	var returns IotTopicRuleErrorActionFirehoseOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) FirehoseInput() *IotTopicRuleErrorActionFirehose {
	var returns *IotTopicRuleErrorActionFirehose
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Http() IotTopicRuleErrorActionHttpOutputReference {
	var returns IotTopicRuleErrorActionHttpOutputReference
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) HttpInput() *IotTopicRuleErrorActionHttp {
	var returns *IotTopicRuleErrorActionHttp
	_jsii_.Get(
		j,
		"httpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) InternalValue() *IotTopicRuleErrorAction {
	var returns *IotTopicRuleErrorAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) IotAnalytics() IotTopicRuleErrorActionIotAnalyticsOutputReference {
	var returns IotTopicRuleErrorActionIotAnalyticsOutputReference
	_jsii_.Get(
		j,
		"iotAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) IotAnalyticsInput() *IotTopicRuleErrorActionIotAnalytics {
	var returns *IotTopicRuleErrorActionIotAnalytics
	_jsii_.Get(
		j,
		"iotAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) IotEvents() IotTopicRuleErrorActionIotEventsOutputReference {
	var returns IotTopicRuleErrorActionIotEventsOutputReference
	_jsii_.Get(
		j,
		"iotEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) IotEventsInput() *IotTopicRuleErrorActionIotEvents {
	var returns *IotTopicRuleErrorActionIotEvents
	_jsii_.Get(
		j,
		"iotEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Kafka() IotTopicRuleErrorActionKafkaOutputReference {
	var returns IotTopicRuleErrorActionKafkaOutputReference
	_jsii_.Get(
		j,
		"kafka",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) KafkaInput() *IotTopicRuleErrorActionKafka {
	var returns *IotTopicRuleErrorActionKafka
	_jsii_.Get(
		j,
		"kafkaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Kinesis() IotTopicRuleErrorActionKinesisOutputReference {
	var returns IotTopicRuleErrorActionKinesisOutputReference
	_jsii_.Get(
		j,
		"kinesis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) KinesisInput() *IotTopicRuleErrorActionKinesis {
	var returns *IotTopicRuleErrorActionKinesis
	_jsii_.Get(
		j,
		"kinesisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Lambda() IotTopicRuleErrorActionLambdaOutputReference {
	var returns IotTopicRuleErrorActionLambdaOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) LambdaInput() *IotTopicRuleErrorActionLambda {
	var returns *IotTopicRuleErrorActionLambda
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Republish() IotTopicRuleErrorActionRepublishOutputReference {
	var returns IotTopicRuleErrorActionRepublishOutputReference
	_jsii_.Get(
		j,
		"republish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) RepublishInput() *IotTopicRuleErrorActionRepublish {
	var returns *IotTopicRuleErrorActionRepublish
	_jsii_.Get(
		j,
		"republishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) S3() IotTopicRuleErrorActionS3OutputReference {
	var returns IotTopicRuleErrorActionS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) S3Input() *IotTopicRuleErrorActionS3 {
	var returns *IotTopicRuleErrorActionS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Sns() IotTopicRuleErrorActionSnsOutputReference {
	var returns IotTopicRuleErrorActionSnsOutputReference
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) SnsInput() *IotTopicRuleErrorActionSns {
	var returns *IotTopicRuleErrorActionSns
	_jsii_.Get(
		j,
		"snsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Sqs() IotTopicRuleErrorActionSqsOutputReference {
	var returns IotTopicRuleErrorActionSqsOutputReference
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) SqsInput() *IotTopicRuleErrorActionSqs {
	var returns *IotTopicRuleErrorActionSqs
	_jsii_.Get(
		j,
		"sqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) StepFunctions() IotTopicRuleErrorActionStepFunctionsOutputReference {
	var returns IotTopicRuleErrorActionStepFunctionsOutputReference
	_jsii_.Get(
		j,
		"stepFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) StepFunctionsInput() *IotTopicRuleErrorActionStepFunctions {
	var returns *IotTopicRuleErrorActionStepFunctions
	_jsii_.Get(
		j,
		"stepFunctionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) Timestream() IotTopicRuleErrorActionTimestreamOutputReference {
	var returns IotTopicRuleErrorActionTimestreamOutputReference
	_jsii_.Get(
		j,
		"timestream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference) TimestreamInput() *IotTopicRuleErrorActionTimestream {
	var returns *IotTopicRuleErrorActionTimestream
	_jsii_.Get(
		j,
		"timestreamInput",
		&returns,
	)
	return returns
}


func NewIotTopicRuleErrorActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IotTopicRuleErrorActionOutputReference {
	_init_.Initialize()

	if err := validateNewIotTopicRuleErrorActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotTopicRuleErrorActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.iotTopicRule.IotTopicRuleErrorActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotTopicRuleErrorActionOutputReference_Override(i IotTopicRuleErrorActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.iotTopicRule.IotTopicRuleErrorActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference)SetInternalValue(val *IotTopicRuleErrorAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutCloudwatchAlarm(value *IotTopicRuleErrorActionCloudwatchAlarm) {
	if err := i.validatePutCloudwatchAlarmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCloudwatchAlarm",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutCloudwatchLogs(value *IotTopicRuleErrorActionCloudwatchLogs) {
	if err := i.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutCloudwatchMetric(value *IotTopicRuleErrorActionCloudwatchMetric) {
	if err := i.validatePutCloudwatchMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCloudwatchMetric",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutDynamodb(value *IotTopicRuleErrorActionDynamodb) {
	if err := i.validatePutDynamodbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDynamodb",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutDynamodbv2(value *IotTopicRuleErrorActionDynamodbv2) {
	if err := i.validatePutDynamodbv2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDynamodbv2",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutElasticsearch(value *IotTopicRuleErrorActionElasticsearch) {
	if err := i.validatePutElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putElasticsearch",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutFirehose(value *IotTopicRuleErrorActionFirehose) {
	if err := i.validatePutFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFirehose",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutHttp(value *IotTopicRuleErrorActionHttp) {
	if err := i.validatePutHttpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putHttp",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutIotAnalytics(value *IotTopicRuleErrorActionIotAnalytics) {
	if err := i.validatePutIotAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotAnalytics",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutIotEvents(value *IotTopicRuleErrorActionIotEvents) {
	if err := i.validatePutIotEventsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIotEvents",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutKafka(value *IotTopicRuleErrorActionKafka) {
	if err := i.validatePutKafkaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putKafka",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutKinesis(value *IotTopicRuleErrorActionKinesis) {
	if err := i.validatePutKinesisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putKinesis",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutLambda(value *IotTopicRuleErrorActionLambda) {
	if err := i.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLambda",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutRepublish(value *IotTopicRuleErrorActionRepublish) {
	if err := i.validatePutRepublishParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putRepublish",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutS3(value *IotTopicRuleErrorActionS3) {
	if err := i.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putS3",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutSns(value *IotTopicRuleErrorActionSns) {
	if err := i.validatePutSnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSns",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutSqs(value *IotTopicRuleErrorActionSqs) {
	if err := i.validatePutSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSqs",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutStepFunctions(value *IotTopicRuleErrorActionStepFunctions) {
	if err := i.validatePutStepFunctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putStepFunctions",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) PutTimestream(value *IotTopicRuleErrorActionTimestream) {
	if err := i.validatePutTimestreamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimestream",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetCloudwatchAlarm() {
	_jsii_.InvokeVoid(
		i,
		"resetCloudwatchAlarm",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		i,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetCloudwatchMetric() {
	_jsii_.InvokeVoid(
		i,
		"resetCloudwatchMetric",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetDynamodb() {
	_jsii_.InvokeVoid(
		i,
		"resetDynamodb",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetDynamodbv2() {
	_jsii_.InvokeVoid(
		i,
		"resetDynamodbv2",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetElasticsearch() {
	_jsii_.InvokeVoid(
		i,
		"resetElasticsearch",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetFirehose() {
	_jsii_.InvokeVoid(
		i,
		"resetFirehose",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetHttp() {
	_jsii_.InvokeVoid(
		i,
		"resetHttp",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetIotAnalytics() {
	_jsii_.InvokeVoid(
		i,
		"resetIotAnalytics",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetIotEvents() {
	_jsii_.InvokeVoid(
		i,
		"resetIotEvents",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetKafka() {
	_jsii_.InvokeVoid(
		i,
		"resetKafka",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetKinesis() {
	_jsii_.InvokeVoid(
		i,
		"resetKinesis",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetLambda() {
	_jsii_.InvokeVoid(
		i,
		"resetLambda",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetRepublish() {
	_jsii_.InvokeVoid(
		i,
		"resetRepublish",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		i,
		"resetS3",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetSns() {
	_jsii_.InvokeVoid(
		i,
		"resetSns",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetSqs() {
	_jsii_.InvokeVoid(
		i,
		"resetSqs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetStepFunctions() {
	_jsii_.InvokeVoid(
		i,
		"resetStepFunctions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ResetTimestream() {
	_jsii_.InvokeVoid(
		i,
		"resetTimestream",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

