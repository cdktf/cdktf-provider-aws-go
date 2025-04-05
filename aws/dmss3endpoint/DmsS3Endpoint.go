// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmss3endpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/dmss3endpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.1/docs/resources/dms_s3_endpoint aws_dms_s3_endpoint}.
type DmsS3Endpoint interface {
	cdktf.TerraformResource
	AddColumnName() interface{}
	SetAddColumnName(val interface{})
	AddColumnNameInput() interface{}
	AddTrailingPaddingCharacter() interface{}
	SetAddTrailingPaddingCharacter(val interface{})
	AddTrailingPaddingCharacterInput() interface{}
	BucketFolder() *string
	SetBucketFolder(val *string)
	BucketFolderInput() *string
	BucketName() *string
	SetBucketName(val *string)
	BucketNameInput() *string
	CannedAclForObjects() *string
	SetCannedAclForObjects(val *string)
	CannedAclForObjectsInput() *string
	CdcInsertsAndUpdates() interface{}
	SetCdcInsertsAndUpdates(val interface{})
	CdcInsertsAndUpdatesInput() interface{}
	CdcInsertsOnly() interface{}
	SetCdcInsertsOnly(val interface{})
	CdcInsertsOnlyInput() interface{}
	CdcMaxBatchInterval() *float64
	SetCdcMaxBatchInterval(val *float64)
	CdcMaxBatchIntervalInput() *float64
	CdcMinFileSize() *float64
	SetCdcMinFileSize(val *float64)
	CdcMinFileSizeInput() *float64
	CdcPath() *string
	SetCdcPath(val *string)
	CdcPathInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateArn() *string
	SetCertificateArn(val *string)
	CertificateArnInput() *string
	CompressionType() *string
	SetCompressionType(val *string)
	CompressionTypeInput() *string
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
	CsvDelimiter() *string
	SetCsvDelimiter(val *string)
	CsvDelimiterInput() *string
	CsvNoSupValue() *string
	SetCsvNoSupValue(val *string)
	CsvNoSupValueInput() *string
	CsvNullValue() *string
	SetCsvNullValue(val *string)
	CsvNullValueInput() *string
	CsvRowDelimiter() *string
	SetCsvRowDelimiter(val *string)
	CsvRowDelimiterInput() *string
	DataFormat() *string
	SetDataFormat(val *string)
	DataFormatInput() *string
	DataPageSize() *float64
	SetDataPageSize(val *float64)
	DataPageSizeInput() *float64
	DatePartitionDelimiter() *string
	SetDatePartitionDelimiter(val *string)
	DatePartitionDelimiterInput() *string
	DatePartitionEnabled() interface{}
	SetDatePartitionEnabled(val interface{})
	DatePartitionEnabledInput() interface{}
	DatePartitionSequence() *string
	SetDatePartitionSequence(val *string)
	DatePartitionSequenceInput() *string
	DatePartitionTimezone() *string
	SetDatePartitionTimezone(val *string)
	DatePartitionTimezoneInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DetachTargetOnLobLookupFailureParquet() interface{}
	SetDetachTargetOnLobLookupFailureParquet(val interface{})
	DetachTargetOnLobLookupFailureParquetInput() interface{}
	DictPageSizeLimit() *float64
	SetDictPageSizeLimit(val *float64)
	DictPageSizeLimitInput() *float64
	EnableStatistics() interface{}
	SetEnableStatistics(val interface{})
	EnableStatisticsInput() interface{}
	EncodingType() *string
	SetEncodingType(val *string)
	EncodingTypeInput() *string
	EncryptionMode() *string
	SetEncryptionMode(val *string)
	EncryptionModeInput() *string
	EndpointArn() *string
	EndpointId() *string
	SetEndpointId(val *string)
	EndpointIdInput() *string
	EndpointType() *string
	SetEndpointType(val *string)
	EndpointTypeInput() *string
	EngineDisplayName() *string
	ExpectedBucketOwner() *string
	SetExpectedBucketOwner(val *string)
	ExpectedBucketOwnerInput() *string
	ExternalId() *string
	ExternalTableDefinition() *string
	SetExternalTableDefinition(val *string)
	ExternalTableDefinitionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GlueCatalogGeneration() interface{}
	SetGlueCatalogGeneration(val interface{})
	GlueCatalogGenerationInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	IgnoreHeaderRows() *float64
	SetIgnoreHeaderRows(val *float64)
	IgnoreHeaderRowsInput() *float64
	IncludeOpForFullLoad() interface{}
	SetIncludeOpForFullLoad(val interface{})
	IncludeOpForFullLoadInput() interface{}
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxFileSize() *float64
	SetMaxFileSize(val *float64)
	MaxFileSizeInput() *float64
	// The tree node.
	Node() constructs.Node
	ParquetTimestampInMillisecond() interface{}
	SetParquetTimestampInMillisecond(val interface{})
	ParquetTimestampInMillisecondInput() interface{}
	ParquetVersion() *string
	SetParquetVersion(val *string)
	ParquetVersionInput() *string
	PreserveTransactions() interface{}
	SetPreserveTransactions(val interface{})
	PreserveTransactionsInput() interface{}
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
	Rfc4180() interface{}
	SetRfc4180(val interface{})
	Rfc4180Input() interface{}
	RowGroupLength() *float64
	SetRowGroupLength(val *float64)
	RowGroupLengthInput() *float64
	ServerSideEncryptionKmsKeyId() *string
	SetServerSideEncryptionKmsKeyId(val *string)
	ServerSideEncryptionKmsKeyIdInput() *string
	ServiceAccessRoleArn() *string
	SetServiceAccessRoleArn(val *string)
	ServiceAccessRoleArnInput() *string
	SslMode() *string
	SetSslMode(val *string)
	SslModeInput() *string
	Status() *string
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
	Timeouts() DmsS3EndpointTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimestampColumnName() *string
	SetTimestampColumnName(val *string)
	TimestampColumnNameInput() *string
	UseCsvNoSupValue() interface{}
	SetUseCsvNoSupValue(val interface{})
	UseCsvNoSupValueInput() interface{}
	UseTaskStartTimeForFullLoadTimestamp() interface{}
	SetUseTaskStartTimeForFullLoadTimestamp(val interface{})
	UseTaskStartTimeForFullLoadTimestampInput() interface{}
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
	PutTimeouts(value *DmsS3EndpointTimeouts)
	ResetAddColumnName()
	ResetAddTrailingPaddingCharacter()
	ResetBucketFolder()
	ResetCannedAclForObjects()
	ResetCdcInsertsAndUpdates()
	ResetCdcInsertsOnly()
	ResetCdcMaxBatchInterval()
	ResetCdcMinFileSize()
	ResetCdcPath()
	ResetCertificateArn()
	ResetCompressionType()
	ResetCsvDelimiter()
	ResetCsvNoSupValue()
	ResetCsvNullValue()
	ResetCsvRowDelimiter()
	ResetDataFormat()
	ResetDataPageSize()
	ResetDatePartitionDelimiter()
	ResetDatePartitionEnabled()
	ResetDatePartitionSequence()
	ResetDatePartitionTimezone()
	ResetDetachTargetOnLobLookupFailureParquet()
	ResetDictPageSizeLimit()
	ResetEnableStatistics()
	ResetEncodingType()
	ResetEncryptionMode()
	ResetExpectedBucketOwner()
	ResetExternalTableDefinition()
	ResetGlueCatalogGeneration()
	ResetId()
	ResetIgnoreHeaderRows()
	ResetIncludeOpForFullLoad()
	ResetKmsKeyArn()
	ResetMaxFileSize()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParquetTimestampInMillisecond()
	ResetParquetVersion()
	ResetPreserveTransactions()
	ResetRfc4180()
	ResetRowGroupLength()
	ResetServerSideEncryptionKmsKeyId()
	ResetSslMode()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetTimestampColumnName()
	ResetUseCsvNoSupValue()
	ResetUseTaskStartTimeForFullLoadTimestamp()
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

// The jsii proxy struct for DmsS3Endpoint
type jsiiProxy_DmsS3Endpoint struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DmsS3Endpoint) AddColumnName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) AddColumnNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) AddTrailingPaddingCharacter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addTrailingPaddingCharacter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) AddTrailingPaddingCharacterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addTrailingPaddingCharacterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) BucketFolder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketFolder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) BucketFolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketFolderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CannedAclForObjects() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cannedAclForObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CannedAclForObjectsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cannedAclForObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CdcInsertsAndUpdates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdcInsertsAndUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CdcInsertsAndUpdatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdcInsertsAndUpdatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CdcInsertsOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdcInsertsOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CdcInsertsOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdcInsertsOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CdcMaxBatchInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdcMaxBatchInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CdcMaxBatchIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdcMaxBatchIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CdcMinFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdcMinFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CdcMinFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdcMinFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CdcPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdcPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CdcPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdcPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CsvDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CsvDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CsvNoSupValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvNoSupValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CsvNoSupValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvNoSupValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CsvNullValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvNullValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CsvNullValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvNullValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CsvRowDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvRowDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) CsvRowDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvRowDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) DataFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) DataFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) DataPageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataPageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) DataPageSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataPageSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) DatePartitionDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datePartitionDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) DatePartitionDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datePartitionDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) DatePartitionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datePartitionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) DatePartitionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datePartitionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) DatePartitionSequence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datePartitionSequence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) DatePartitionSequenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datePartitionSequenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) DatePartitionTimezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datePartitionTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) DatePartitionTimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datePartitionTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) DetachTargetOnLobLookupFailureParquet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detachTargetOnLobLookupFailureParquet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) DetachTargetOnLobLookupFailureParquetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detachTargetOnLobLookupFailureParquetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) DictPageSizeLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dictPageSizeLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) DictPageSizeLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dictPageSizeLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) EnableStatistics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStatistics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) EnableStatisticsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStatisticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) EncodingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) EncodingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) EncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) EncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) EndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) EndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) EndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) EndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) EngineDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) ExpectedBucketOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedBucketOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) ExpectedBucketOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedBucketOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) ExternalTableDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalTableDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) ExternalTableDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalTableDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) GlueCatalogGeneration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"glueCatalogGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) GlueCatalogGenerationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"glueCatalogGenerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) IgnoreHeaderRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ignoreHeaderRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) IgnoreHeaderRowsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ignoreHeaderRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) IncludeOpForFullLoad() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeOpForFullLoad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) IncludeOpForFullLoadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeOpForFullLoadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) MaxFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) ParquetTimestampInMillisecond() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parquetTimestampInMillisecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) ParquetTimestampInMillisecondInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parquetTimestampInMillisecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) ParquetVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parquetVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) ParquetVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parquetVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) PreserveTransactions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveTransactions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) PreserveTransactionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveTransactionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) Rfc4180() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rfc4180",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) Rfc4180Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rfc4180Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) RowGroupLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowGroupLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) RowGroupLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowGroupLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) ServerSideEncryptionKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) ServerSideEncryptionKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) ServiceAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) SslMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) SslModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) Timeouts() DmsS3EndpointTimeoutsOutputReference {
	var returns DmsS3EndpointTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) TimestampColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) TimestampColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) UseCsvNoSupValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCsvNoSupValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) UseCsvNoSupValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCsvNoSupValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) UseTaskStartTimeForFullLoadTimestamp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTaskStartTimeForFullLoadTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsS3Endpoint) UseTaskStartTimeForFullLoadTimestampInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTaskStartTimeForFullLoadTimestampInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.1/docs/resources/dms_s3_endpoint aws_dms_s3_endpoint} Resource.
func NewDmsS3Endpoint(scope constructs.Construct, id *string, config *DmsS3EndpointConfig) DmsS3Endpoint {
	_init_.Initialize()

	if err := validateNewDmsS3EndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsS3Endpoint{}

	_jsii_.Create(
		"@cdktf/provider-aws.dmsS3Endpoint.DmsS3Endpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.1/docs/resources/dms_s3_endpoint aws_dms_s3_endpoint} Resource.
func NewDmsS3Endpoint_Override(d DmsS3Endpoint, scope constructs.Construct, id *string, config *DmsS3EndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dmsS3Endpoint.DmsS3Endpoint",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetAddColumnName(val interface{}) {
	if err := j.validateSetAddColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addColumnName",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetAddTrailingPaddingCharacter(val interface{}) {
	if err := j.validateSetAddTrailingPaddingCharacterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addTrailingPaddingCharacter",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetBucketFolder(val *string) {
	if err := j.validateSetBucketFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketFolder",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetCannedAclForObjects(val *string) {
	if err := j.validateSetCannedAclForObjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cannedAclForObjects",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetCdcInsertsAndUpdates(val interface{}) {
	if err := j.validateSetCdcInsertsAndUpdatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdcInsertsAndUpdates",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetCdcInsertsOnly(val interface{}) {
	if err := j.validateSetCdcInsertsOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdcInsertsOnly",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetCdcMaxBatchInterval(val *float64) {
	if err := j.validateSetCdcMaxBatchIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdcMaxBatchInterval",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetCdcMinFileSize(val *float64) {
	if err := j.validateSetCdcMinFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdcMinFileSize",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetCdcPath(val *string) {
	if err := j.validateSetCdcPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdcPath",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetCertificateArn(val *string) {
	if err := j.validateSetCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetCompressionType(val *string) {
	if err := j.validateSetCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionType",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetCsvDelimiter(val *string) {
	if err := j.validateSetCsvDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"csvDelimiter",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetCsvNoSupValue(val *string) {
	if err := j.validateSetCsvNoSupValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"csvNoSupValue",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetCsvNullValue(val *string) {
	if err := j.validateSetCsvNullValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"csvNullValue",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetCsvRowDelimiter(val *string) {
	if err := j.validateSetCsvRowDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"csvRowDelimiter",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetDataFormat(val *string) {
	if err := j.validateSetDataFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFormat",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetDataPageSize(val *float64) {
	if err := j.validateSetDataPageSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataPageSize",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetDatePartitionDelimiter(val *string) {
	if err := j.validateSetDatePartitionDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datePartitionDelimiter",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetDatePartitionEnabled(val interface{}) {
	if err := j.validateSetDatePartitionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datePartitionEnabled",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetDatePartitionSequence(val *string) {
	if err := j.validateSetDatePartitionSequenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datePartitionSequence",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetDatePartitionTimezone(val *string) {
	if err := j.validateSetDatePartitionTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datePartitionTimezone",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetDetachTargetOnLobLookupFailureParquet(val interface{}) {
	if err := j.validateSetDetachTargetOnLobLookupFailureParquetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detachTargetOnLobLookupFailureParquet",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetDictPageSizeLimit(val *float64) {
	if err := j.validateSetDictPageSizeLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dictPageSizeLimit",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetEnableStatistics(val interface{}) {
	if err := j.validateSetEnableStatisticsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableStatistics",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetEncodingType(val *string) {
	if err := j.validateSetEncodingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encodingType",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetEncryptionMode(val *string) {
	if err := j.validateSetEncryptionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionMode",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetEndpointId(val *string) {
	if err := j.validateSetEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointId",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetEndpointType(val *string) {
	if err := j.validateSetEndpointTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetExpectedBucketOwner(val *string) {
	if err := j.validateSetExpectedBucketOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expectedBucketOwner",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetExternalTableDefinition(val *string) {
	if err := j.validateSetExternalTableDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalTableDefinition",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetGlueCatalogGeneration(val interface{}) {
	if err := j.validateSetGlueCatalogGenerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"glueCatalogGeneration",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetIgnoreHeaderRows(val *float64) {
	if err := j.validateSetIgnoreHeaderRowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreHeaderRows",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetIncludeOpForFullLoad(val interface{}) {
	if err := j.validateSetIncludeOpForFullLoadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeOpForFullLoad",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetMaxFileSize(val *float64) {
	if err := j.validateSetMaxFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFileSize",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetParquetTimestampInMillisecond(val interface{}) {
	if err := j.validateSetParquetTimestampInMillisecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parquetTimestampInMillisecond",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetParquetVersion(val *string) {
	if err := j.validateSetParquetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parquetVersion",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetPreserveTransactions(val interface{}) {
	if err := j.validateSetPreserveTransactionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveTransactions",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetRfc4180(val interface{}) {
	if err := j.validateSetRfc4180Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rfc4180",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetRowGroupLength(val *float64) {
	if err := j.validateSetRowGroupLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowGroupLength",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetServerSideEncryptionKmsKeyId(val *string) {
	if err := j.validateSetServerSideEncryptionKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverSideEncryptionKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetServiceAccessRoleArn(val *string) {
	if err := j.validateSetServiceAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetSslMode(val *string) {
	if err := j.validateSetSslModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslMode",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetTimestampColumnName(val *string) {
	if err := j.validateSetTimestampColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampColumnName",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetUseCsvNoSupValue(val interface{}) {
	if err := j.validateSetUseCsvNoSupValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCsvNoSupValue",
		val,
	)
}

func (j *jsiiProxy_DmsS3Endpoint)SetUseTaskStartTimeForFullLoadTimestamp(val interface{}) {
	if err := j.validateSetUseTaskStartTimeForFullLoadTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useTaskStartTimeForFullLoadTimestamp",
		val,
	)
}

// Generates CDKTF code for importing a DmsS3Endpoint resource upon running "cdktf plan <stack-name>".
func DmsS3Endpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDmsS3Endpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dmsS3Endpoint.DmsS3Endpoint",
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
func DmsS3Endpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsS3Endpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dmsS3Endpoint.DmsS3Endpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DmsS3Endpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsS3Endpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dmsS3Endpoint.DmsS3Endpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DmsS3Endpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsS3Endpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dmsS3Endpoint.DmsS3Endpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DmsS3Endpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dmsS3Endpoint.DmsS3Endpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DmsS3Endpoint) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DmsS3Endpoint) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DmsS3Endpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsS3Endpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsS3Endpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsS3Endpoint) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsS3Endpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsS3Endpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsS3Endpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsS3Endpoint) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsS3Endpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsS3Endpoint) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsS3Endpoint) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DmsS3Endpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsS3Endpoint) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DmsS3Endpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DmsS3Endpoint) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DmsS3Endpoint) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DmsS3Endpoint) PutTimeouts(value *DmsS3EndpointTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetAddColumnName() {
	_jsii_.InvokeVoid(
		d,
		"resetAddColumnName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetAddTrailingPaddingCharacter() {
	_jsii_.InvokeVoid(
		d,
		"resetAddTrailingPaddingCharacter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetBucketFolder() {
	_jsii_.InvokeVoid(
		d,
		"resetBucketFolder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetCannedAclForObjects() {
	_jsii_.InvokeVoid(
		d,
		"resetCannedAclForObjects",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetCdcInsertsAndUpdates() {
	_jsii_.InvokeVoid(
		d,
		"resetCdcInsertsAndUpdates",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetCdcInsertsOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetCdcInsertsOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetCdcMaxBatchInterval() {
	_jsii_.InvokeVoid(
		d,
		"resetCdcMaxBatchInterval",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetCdcMinFileSize() {
	_jsii_.InvokeVoid(
		d,
		"resetCdcMinFileSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetCdcPath() {
	_jsii_.InvokeVoid(
		d,
		"resetCdcPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetCertificateArn() {
	_jsii_.InvokeVoid(
		d,
		"resetCertificateArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetCompressionType() {
	_jsii_.InvokeVoid(
		d,
		"resetCompressionType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetCsvDelimiter() {
	_jsii_.InvokeVoid(
		d,
		"resetCsvDelimiter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetCsvNoSupValue() {
	_jsii_.InvokeVoid(
		d,
		"resetCsvNoSupValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetCsvNullValue() {
	_jsii_.InvokeVoid(
		d,
		"resetCsvNullValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetCsvRowDelimiter() {
	_jsii_.InvokeVoid(
		d,
		"resetCsvRowDelimiter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetDataFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetDataFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetDataPageSize() {
	_jsii_.InvokeVoid(
		d,
		"resetDataPageSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetDatePartitionDelimiter() {
	_jsii_.InvokeVoid(
		d,
		"resetDatePartitionDelimiter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetDatePartitionEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetDatePartitionEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetDatePartitionSequence() {
	_jsii_.InvokeVoid(
		d,
		"resetDatePartitionSequence",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetDatePartitionTimezone() {
	_jsii_.InvokeVoid(
		d,
		"resetDatePartitionTimezone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetDetachTargetOnLobLookupFailureParquet() {
	_jsii_.InvokeVoid(
		d,
		"resetDetachTargetOnLobLookupFailureParquet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetDictPageSizeLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetDictPageSizeLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetEnableStatistics() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableStatistics",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetEncodingType() {
	_jsii_.InvokeVoid(
		d,
		"resetEncodingType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetEncryptionMode() {
	_jsii_.InvokeVoid(
		d,
		"resetEncryptionMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetExpectedBucketOwner() {
	_jsii_.InvokeVoid(
		d,
		"resetExpectedBucketOwner",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetExternalTableDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetExternalTableDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetGlueCatalogGeneration() {
	_jsii_.InvokeVoid(
		d,
		"resetGlueCatalogGeneration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetIgnoreHeaderRows() {
	_jsii_.InvokeVoid(
		d,
		"resetIgnoreHeaderRows",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetIncludeOpForFullLoad() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeOpForFullLoad",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		d,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetMaxFileSize() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxFileSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetParquetTimestampInMillisecond() {
	_jsii_.InvokeVoid(
		d,
		"resetParquetTimestampInMillisecond",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetParquetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetParquetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetPreserveTransactions() {
	_jsii_.InvokeVoid(
		d,
		"resetPreserveTransactions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetRfc4180() {
	_jsii_.InvokeVoid(
		d,
		"resetRfc4180",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetRowGroupLength() {
	_jsii_.InvokeVoid(
		d,
		"resetRowGroupLength",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetServerSideEncryptionKmsKeyId() {
	_jsii_.InvokeVoid(
		d,
		"resetServerSideEncryptionKmsKeyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetSslMode() {
	_jsii_.InvokeVoid(
		d,
		"resetSslMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetTagsAll() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetTimestampColumnName() {
	_jsii_.InvokeVoid(
		d,
		"resetTimestampColumnName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetUseCsvNoSupValue() {
	_jsii_.InvokeVoid(
		d,
		"resetUseCsvNoSupValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) ResetUseTaskStartTimeForFullLoadTimestamp() {
	_jsii_.InvokeVoid(
		d,
		"resetUseTaskStartTimeForFullLoadTimestamp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsS3Endpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsS3Endpoint) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsS3Endpoint) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsS3Endpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsS3Endpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsS3Endpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

