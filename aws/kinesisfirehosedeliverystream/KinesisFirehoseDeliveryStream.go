// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/kinesisfirehosedeliverystream/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/kinesis_firehose_delivery_stream aws_kinesis_firehose_delivery_stream}.
type KinesisFirehoseDeliveryStream interface {
	cdktf.TerraformResource
	Arn() *string
	SetArn(val *string)
	ArnInput() *string
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
	Destination() *string
	SetDestination(val *string)
	DestinationId() *string
	SetDestinationId(val *string)
	DestinationIdInput() *string
	DestinationInput() *string
	ElasticsearchConfiguration() KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference
	ElasticsearchConfigurationInput() *KinesisFirehoseDeliveryStreamElasticsearchConfiguration
	ExtendedS3Configuration() KinesisFirehoseDeliveryStreamExtendedS3ConfigurationOutputReference
	ExtendedS3ConfigurationInput() *KinesisFirehoseDeliveryStreamExtendedS3Configuration
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpEndpointConfiguration() KinesisFirehoseDeliveryStreamHttpEndpointConfigurationOutputReference
	HttpEndpointConfigurationInput() *KinesisFirehoseDeliveryStreamHttpEndpointConfiguration
	IcebergConfiguration() KinesisFirehoseDeliveryStreamIcebergConfigurationOutputReference
	IcebergConfigurationInput() *KinesisFirehoseDeliveryStreamIcebergConfiguration
	Id() *string
	SetId(val *string)
	IdInput() *string
	KinesisSourceConfiguration() KinesisFirehoseDeliveryStreamKinesisSourceConfigurationOutputReference
	KinesisSourceConfigurationInput() *KinesisFirehoseDeliveryStreamKinesisSourceConfiguration
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MskSourceConfiguration() KinesisFirehoseDeliveryStreamMskSourceConfigurationOutputReference
	MskSourceConfigurationInput() *KinesisFirehoseDeliveryStreamMskSourceConfiguration
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OpensearchConfiguration() KinesisFirehoseDeliveryStreamOpensearchConfigurationOutputReference
	OpensearchConfigurationInput() *KinesisFirehoseDeliveryStreamOpensearchConfiguration
	OpensearchserverlessConfiguration() KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference
	OpensearchserverlessConfigurationInput() *KinesisFirehoseDeliveryStreamOpensearchserverlessConfiguration
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
	RedshiftConfiguration() KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference
	RedshiftConfigurationInput() *KinesisFirehoseDeliveryStreamRedshiftConfiguration
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ServerSideEncryption() KinesisFirehoseDeliveryStreamServerSideEncryptionOutputReference
	ServerSideEncryptionInput() *KinesisFirehoseDeliveryStreamServerSideEncryption
	SnowflakeConfiguration() KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference
	SnowflakeConfigurationInput() *KinesisFirehoseDeliveryStreamSnowflakeConfiguration
	SplunkConfiguration() KinesisFirehoseDeliveryStreamSplunkConfigurationOutputReference
	SplunkConfigurationInput() *KinesisFirehoseDeliveryStreamSplunkConfiguration
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
	Timeouts() KinesisFirehoseDeliveryStreamTimeoutsOutputReference
	TimeoutsInput() interface{}
	VersionId() *string
	SetVersionId(val *string)
	VersionIdInput() *string
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
	PutElasticsearchConfiguration(value *KinesisFirehoseDeliveryStreamElasticsearchConfiguration)
	PutExtendedS3Configuration(value *KinesisFirehoseDeliveryStreamExtendedS3Configuration)
	PutHttpEndpointConfiguration(value *KinesisFirehoseDeliveryStreamHttpEndpointConfiguration)
	PutIcebergConfiguration(value *KinesisFirehoseDeliveryStreamIcebergConfiguration)
	PutKinesisSourceConfiguration(value *KinesisFirehoseDeliveryStreamKinesisSourceConfiguration)
	PutMskSourceConfiguration(value *KinesisFirehoseDeliveryStreamMskSourceConfiguration)
	PutOpensearchConfiguration(value *KinesisFirehoseDeliveryStreamOpensearchConfiguration)
	PutOpensearchserverlessConfiguration(value *KinesisFirehoseDeliveryStreamOpensearchserverlessConfiguration)
	PutRedshiftConfiguration(value *KinesisFirehoseDeliveryStreamRedshiftConfiguration)
	PutServerSideEncryption(value *KinesisFirehoseDeliveryStreamServerSideEncryption)
	PutSnowflakeConfiguration(value *KinesisFirehoseDeliveryStreamSnowflakeConfiguration)
	PutSplunkConfiguration(value *KinesisFirehoseDeliveryStreamSplunkConfiguration)
	PutTimeouts(value *KinesisFirehoseDeliveryStreamTimeouts)
	ResetArn()
	ResetDestinationId()
	ResetElasticsearchConfiguration()
	ResetExtendedS3Configuration()
	ResetHttpEndpointConfiguration()
	ResetIcebergConfiguration()
	ResetId()
	ResetKinesisSourceConfiguration()
	ResetMskSourceConfiguration()
	ResetOpensearchConfiguration()
	ResetOpensearchserverlessConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRedshiftConfiguration()
	ResetRegion()
	ResetServerSideEncryption()
	ResetSnowflakeConfiguration()
	ResetSplunkConfiguration()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetVersionId()
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

// The jsii proxy struct for KinesisFirehoseDeliveryStream
type jsiiProxy_KinesisFirehoseDeliveryStream struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) DestinationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) DestinationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) DestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) ElasticsearchConfiguration() KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamElasticsearchConfigurationOutputReference
	_jsii_.Get(
		j,
		"elasticsearchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) ElasticsearchConfigurationInput() *KinesisFirehoseDeliveryStreamElasticsearchConfiguration {
	var returns *KinesisFirehoseDeliveryStreamElasticsearchConfiguration
	_jsii_.Get(
		j,
		"elasticsearchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) ExtendedS3Configuration() KinesisFirehoseDeliveryStreamExtendedS3ConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamExtendedS3ConfigurationOutputReference
	_jsii_.Get(
		j,
		"extendedS3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) ExtendedS3ConfigurationInput() *KinesisFirehoseDeliveryStreamExtendedS3Configuration {
	var returns *KinesisFirehoseDeliveryStreamExtendedS3Configuration
	_jsii_.Get(
		j,
		"extendedS3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) HttpEndpointConfiguration() KinesisFirehoseDeliveryStreamHttpEndpointConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamHttpEndpointConfigurationOutputReference
	_jsii_.Get(
		j,
		"httpEndpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) HttpEndpointConfigurationInput() *KinesisFirehoseDeliveryStreamHttpEndpointConfiguration {
	var returns *KinesisFirehoseDeliveryStreamHttpEndpointConfiguration
	_jsii_.Get(
		j,
		"httpEndpointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) IcebergConfiguration() KinesisFirehoseDeliveryStreamIcebergConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamIcebergConfigurationOutputReference
	_jsii_.Get(
		j,
		"icebergConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) IcebergConfigurationInput() *KinesisFirehoseDeliveryStreamIcebergConfiguration {
	var returns *KinesisFirehoseDeliveryStreamIcebergConfiguration
	_jsii_.Get(
		j,
		"icebergConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) KinesisSourceConfiguration() KinesisFirehoseDeliveryStreamKinesisSourceConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamKinesisSourceConfigurationOutputReference
	_jsii_.Get(
		j,
		"kinesisSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) KinesisSourceConfigurationInput() *KinesisFirehoseDeliveryStreamKinesisSourceConfiguration {
	var returns *KinesisFirehoseDeliveryStreamKinesisSourceConfiguration
	_jsii_.Get(
		j,
		"kinesisSourceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) MskSourceConfiguration() KinesisFirehoseDeliveryStreamMskSourceConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamMskSourceConfigurationOutputReference
	_jsii_.Get(
		j,
		"mskSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) MskSourceConfigurationInput() *KinesisFirehoseDeliveryStreamMskSourceConfiguration {
	var returns *KinesisFirehoseDeliveryStreamMskSourceConfiguration
	_jsii_.Get(
		j,
		"mskSourceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) OpensearchConfiguration() KinesisFirehoseDeliveryStreamOpensearchConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamOpensearchConfigurationOutputReference
	_jsii_.Get(
		j,
		"opensearchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) OpensearchConfigurationInput() *KinesisFirehoseDeliveryStreamOpensearchConfiguration {
	var returns *KinesisFirehoseDeliveryStreamOpensearchConfiguration
	_jsii_.Get(
		j,
		"opensearchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) OpensearchserverlessConfiguration() KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamOpensearchserverlessConfigurationOutputReference
	_jsii_.Get(
		j,
		"opensearchserverlessConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) OpensearchserverlessConfigurationInput() *KinesisFirehoseDeliveryStreamOpensearchserverlessConfiguration {
	var returns *KinesisFirehoseDeliveryStreamOpensearchserverlessConfiguration
	_jsii_.Get(
		j,
		"opensearchserverlessConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) RedshiftConfiguration() KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamRedshiftConfigurationOutputReference
	_jsii_.Get(
		j,
		"redshiftConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) RedshiftConfigurationInput() *KinesisFirehoseDeliveryStreamRedshiftConfiguration {
	var returns *KinesisFirehoseDeliveryStreamRedshiftConfiguration
	_jsii_.Get(
		j,
		"redshiftConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) ServerSideEncryption() KinesisFirehoseDeliveryStreamServerSideEncryptionOutputReference {
	var returns KinesisFirehoseDeliveryStreamServerSideEncryptionOutputReference
	_jsii_.Get(
		j,
		"serverSideEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) ServerSideEncryptionInput() *KinesisFirehoseDeliveryStreamServerSideEncryption {
	var returns *KinesisFirehoseDeliveryStreamServerSideEncryption
	_jsii_.Get(
		j,
		"serverSideEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) SnowflakeConfiguration() KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamSnowflakeConfigurationOutputReference
	_jsii_.Get(
		j,
		"snowflakeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) SnowflakeConfigurationInput() *KinesisFirehoseDeliveryStreamSnowflakeConfiguration {
	var returns *KinesisFirehoseDeliveryStreamSnowflakeConfiguration
	_jsii_.Get(
		j,
		"snowflakeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) SplunkConfiguration() KinesisFirehoseDeliveryStreamSplunkConfigurationOutputReference {
	var returns KinesisFirehoseDeliveryStreamSplunkConfigurationOutputReference
	_jsii_.Get(
		j,
		"splunkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) SplunkConfigurationInput() *KinesisFirehoseDeliveryStreamSplunkConfiguration {
	var returns *KinesisFirehoseDeliveryStreamSplunkConfiguration
	_jsii_.Get(
		j,
		"splunkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) Timeouts() KinesisFirehoseDeliveryStreamTimeoutsOutputReference {
	var returns KinesisFirehoseDeliveryStreamTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream) VersionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/kinesis_firehose_delivery_stream aws_kinesis_firehose_delivery_stream} Resource.
func NewKinesisFirehoseDeliveryStream(scope constructs.Construct, id *string, config *KinesisFirehoseDeliveryStreamConfig) KinesisFirehoseDeliveryStream {
	_init_.Initialize()

	if err := validateNewKinesisFirehoseDeliveryStreamParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisFirehoseDeliveryStream{}

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStream",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/kinesis_firehose_delivery_stream aws_kinesis_firehose_delivery_stream} Resource.
func NewKinesisFirehoseDeliveryStream_Override(k KinesisFirehoseDeliveryStream, scope constructs.Construct, id *string, config *KinesisFirehoseDeliveryStreamConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStream",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream)SetArn(val *string) {
	if err := j.validateSetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream)SetDestination(val *string) {
	if err := j.validateSetDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destination",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream)SetDestinationId(val *string) {
	if err := j.validateSetDestinationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationId",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStream)SetVersionId(val *string) {
	if err := j.validateSetVersionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionId",
		val,
	)
}

// Generates CDKTF code for importing a KinesisFirehoseDeliveryStream resource upon running "cdktf plan <stack-name>".
func KinesisFirehoseDeliveryStream_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateKinesisFirehoseDeliveryStream_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStream",
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
func KinesisFirehoseDeliveryStream_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKinesisFirehoseDeliveryStream_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStream",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KinesisFirehoseDeliveryStream_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKinesisFirehoseDeliveryStream_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStream",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KinesisFirehoseDeliveryStream_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKinesisFirehoseDeliveryStream_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStream",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KinesisFirehoseDeliveryStream_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStream",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) PutElasticsearchConfiguration(value *KinesisFirehoseDeliveryStreamElasticsearchConfiguration) {
	if err := k.validatePutElasticsearchConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putElasticsearchConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) PutExtendedS3Configuration(value *KinesisFirehoseDeliveryStreamExtendedS3Configuration) {
	if err := k.validatePutExtendedS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putExtendedS3Configuration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) PutHttpEndpointConfiguration(value *KinesisFirehoseDeliveryStreamHttpEndpointConfiguration) {
	if err := k.validatePutHttpEndpointConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putHttpEndpointConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) PutIcebergConfiguration(value *KinesisFirehoseDeliveryStreamIcebergConfiguration) {
	if err := k.validatePutIcebergConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putIcebergConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) PutKinesisSourceConfiguration(value *KinesisFirehoseDeliveryStreamKinesisSourceConfiguration) {
	if err := k.validatePutKinesisSourceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putKinesisSourceConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) PutMskSourceConfiguration(value *KinesisFirehoseDeliveryStreamMskSourceConfiguration) {
	if err := k.validatePutMskSourceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putMskSourceConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) PutOpensearchConfiguration(value *KinesisFirehoseDeliveryStreamOpensearchConfiguration) {
	if err := k.validatePutOpensearchConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putOpensearchConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) PutOpensearchserverlessConfiguration(value *KinesisFirehoseDeliveryStreamOpensearchserverlessConfiguration) {
	if err := k.validatePutOpensearchserverlessConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putOpensearchserverlessConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) PutRedshiftConfiguration(value *KinesisFirehoseDeliveryStreamRedshiftConfiguration) {
	if err := k.validatePutRedshiftConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putRedshiftConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) PutServerSideEncryption(value *KinesisFirehoseDeliveryStreamServerSideEncryption) {
	if err := k.validatePutServerSideEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putServerSideEncryption",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) PutSnowflakeConfiguration(value *KinesisFirehoseDeliveryStreamSnowflakeConfiguration) {
	if err := k.validatePutSnowflakeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSnowflakeConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) PutSplunkConfiguration(value *KinesisFirehoseDeliveryStreamSplunkConfiguration) {
	if err := k.validatePutSplunkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSplunkConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) PutTimeouts(value *KinesisFirehoseDeliveryStreamTimeouts) {
	if err := k.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetArn() {
	_jsii_.InvokeVoid(
		k,
		"resetArn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetDestinationId() {
	_jsii_.InvokeVoid(
		k,
		"resetDestinationId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetElasticsearchConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetElasticsearchConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetExtendedS3Configuration() {
	_jsii_.InvokeVoid(
		k,
		"resetExtendedS3Configuration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetHttpEndpointConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetHttpEndpointConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetIcebergConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetIcebergConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetKinesisSourceConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetKinesisSourceConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetMskSourceConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetMskSourceConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetOpensearchConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetOpensearchConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetOpensearchserverlessConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetOpensearchserverlessConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetRedshiftConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetRedshiftConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetRegion() {
	_jsii_.InvokeVoid(
		k,
		"resetRegion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetServerSideEncryption() {
	_jsii_.InvokeVoid(
		k,
		"resetServerSideEncryption",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetSnowflakeConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetSnowflakeConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetSplunkConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetSplunkConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetTags() {
	_jsii_.InvokeVoid(
		k,
		"resetTags",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetTagsAll() {
	_jsii_.InvokeVoid(
		k,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ResetVersionId() {
	_jsii_.InvokeVoid(
		k,
		"resetVersionId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStream) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

