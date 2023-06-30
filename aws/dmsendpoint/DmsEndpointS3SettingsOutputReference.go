package dmsendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/dmsendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsEndpointS3SettingsOutputReference interface {
	cdktf.ComplexObject
	AddColumnName() interface{}
	SetAddColumnName(val interface{})
	AddColumnNameInput() interface{}
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
	CompressionType() *string
	SetCompressionType(val *string)
	CompressionTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
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
	ExternalTableDefinition() *string
	SetExternalTableDefinition(val *string)
	ExternalTableDefinitionInput() *string
	// Experimental.
	Fqn() *string
	IgnoreHeaderRows() *float64
	SetIgnoreHeaderRows(val *float64)
	IgnoreHeaderRowsInput() *float64
	IncludeOpForFullLoad() interface{}
	SetIncludeOpForFullLoad(val interface{})
	IncludeOpForFullLoadInput() interface{}
	InternalValue() *DmsEndpointS3Settings
	SetInternalValue(val *DmsEndpointS3Settings)
	MaxFileSize() *float64
	SetMaxFileSize(val *float64)
	MaxFileSizeInput() *float64
	ParquetTimestampInMillisecond() interface{}
	SetParquetTimestampInMillisecond(val interface{})
	ParquetTimestampInMillisecondInput() interface{}
	ParquetVersion() *string
	SetParquetVersion(val *string)
	ParquetVersionInput() *string
	PreserveTransactions() interface{}
	SetPreserveTransactions(val interface{})
	PreserveTransactionsInput() interface{}
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
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimestampColumnName() *string
	SetTimestampColumnName(val *string)
	TimestampColumnNameInput() *string
	UseCsvNoSupValue() interface{}
	SetUseCsvNoSupValue(val interface{})
	UseCsvNoSupValueInput() interface{}
	UseTaskStartTimeForFullLoadTimestamp() interface{}
	SetUseTaskStartTimeForFullLoadTimestamp(val interface{})
	UseTaskStartTimeForFullLoadTimestampInput() interface{}
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
	ResetAddColumnName()
	ResetBucketFolder()
	ResetBucketName()
	ResetCannedAclForObjects()
	ResetCdcInsertsAndUpdates()
	ResetCdcInsertsOnly()
	ResetCdcMaxBatchInterval()
	ResetCdcMinFileSize()
	ResetCdcPath()
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
	ResetDictPageSizeLimit()
	ResetEnableStatistics()
	ResetEncodingType()
	ResetEncryptionMode()
	ResetExternalTableDefinition()
	ResetIgnoreHeaderRows()
	ResetIncludeOpForFullLoad()
	ResetMaxFileSize()
	ResetParquetTimestampInMillisecond()
	ResetParquetVersion()
	ResetPreserveTransactions()
	ResetRfc4180()
	ResetRowGroupLength()
	ResetServerSideEncryptionKmsKeyId()
	ResetServiceAccessRoleArn()
	ResetTimestampColumnName()
	ResetUseCsvNoSupValue()
	ResetUseTaskStartTimeForFullLoadTimestamp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsEndpointS3SettingsOutputReference
type jsiiProxy_DmsEndpointS3SettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) AddColumnName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) AddColumnNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) BucketFolder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketFolder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) BucketFolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketFolderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CannedAclForObjects() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cannedAclForObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CannedAclForObjectsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cannedAclForObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CdcInsertsAndUpdates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdcInsertsAndUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CdcInsertsAndUpdatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdcInsertsAndUpdatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CdcInsertsOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdcInsertsOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CdcInsertsOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdcInsertsOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CdcMaxBatchInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdcMaxBatchInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CdcMaxBatchIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdcMaxBatchIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CdcMinFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdcMinFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CdcMinFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdcMinFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CdcPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdcPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CdcPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdcPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CsvDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CsvDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CsvNoSupValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvNoSupValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CsvNoSupValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvNoSupValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CsvNullValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvNullValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CsvNullValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvNullValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CsvRowDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvRowDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CsvRowDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvRowDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) DataFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) DataFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) DataPageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataPageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) DataPageSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataPageSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) DatePartitionDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datePartitionDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) DatePartitionDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datePartitionDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) DatePartitionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datePartitionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) DatePartitionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datePartitionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) DatePartitionSequence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datePartitionSequence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) DatePartitionSequenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datePartitionSequenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) DictPageSizeLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dictPageSizeLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) DictPageSizeLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dictPageSizeLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) EnableStatistics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStatistics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) EnableStatisticsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStatisticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) EncodingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) EncodingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) EncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) EncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ExternalTableDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalTableDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ExternalTableDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalTableDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) IgnoreHeaderRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ignoreHeaderRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) IgnoreHeaderRowsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ignoreHeaderRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) IncludeOpForFullLoad() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeOpForFullLoad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) IncludeOpForFullLoadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeOpForFullLoadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) InternalValue() *DmsEndpointS3Settings {
	var returns *DmsEndpointS3Settings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) MaxFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ParquetTimestampInMillisecond() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parquetTimestampInMillisecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ParquetTimestampInMillisecondInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parquetTimestampInMillisecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ParquetVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parquetVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ParquetVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parquetVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) PreserveTransactions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveTransactions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) PreserveTransactionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveTransactionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) Rfc4180() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rfc4180",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) Rfc4180Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rfc4180Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) RowGroupLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowGroupLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) RowGroupLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowGroupLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ServerSideEncryptionKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ServerSideEncryptionKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ServiceAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) TimestampColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) TimestampColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) UseCsvNoSupValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCsvNoSupValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) UseCsvNoSupValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCsvNoSupValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) UseTaskStartTimeForFullLoadTimestamp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTaskStartTimeForFullLoadTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) UseTaskStartTimeForFullLoadTimestampInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTaskStartTimeForFullLoadTimestampInput",
		&returns,
	)
	return returns
}


func NewDmsEndpointS3SettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DmsEndpointS3SettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDmsEndpointS3SettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsEndpointS3SettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dmsEndpoint.DmsEndpointS3SettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDmsEndpointS3SettingsOutputReference_Override(d DmsEndpointS3SettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dmsEndpoint.DmsEndpointS3SettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetAddColumnName(val interface{}) {
	if err := j.validateSetAddColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addColumnName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetBucketFolder(val *string) {
	if err := j.validateSetBucketFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketFolder",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetCannedAclForObjects(val *string) {
	if err := j.validateSetCannedAclForObjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cannedAclForObjects",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetCdcInsertsAndUpdates(val interface{}) {
	if err := j.validateSetCdcInsertsAndUpdatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdcInsertsAndUpdates",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetCdcInsertsOnly(val interface{}) {
	if err := j.validateSetCdcInsertsOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdcInsertsOnly",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetCdcMaxBatchInterval(val *float64) {
	if err := j.validateSetCdcMaxBatchIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdcMaxBatchInterval",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetCdcMinFileSize(val *float64) {
	if err := j.validateSetCdcMinFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdcMinFileSize",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetCdcPath(val *string) {
	if err := j.validateSetCdcPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdcPath",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetCompressionType(val *string) {
	if err := j.validateSetCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionType",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetCsvDelimiter(val *string) {
	if err := j.validateSetCsvDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"csvDelimiter",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetCsvNoSupValue(val *string) {
	if err := j.validateSetCsvNoSupValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"csvNoSupValue",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetCsvNullValue(val *string) {
	if err := j.validateSetCsvNullValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"csvNullValue",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetCsvRowDelimiter(val *string) {
	if err := j.validateSetCsvRowDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"csvRowDelimiter",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetDataFormat(val *string) {
	if err := j.validateSetDataFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFormat",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetDataPageSize(val *float64) {
	if err := j.validateSetDataPageSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataPageSize",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetDatePartitionDelimiter(val *string) {
	if err := j.validateSetDatePartitionDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datePartitionDelimiter",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetDatePartitionEnabled(val interface{}) {
	if err := j.validateSetDatePartitionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datePartitionEnabled",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetDatePartitionSequence(val *string) {
	if err := j.validateSetDatePartitionSequenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datePartitionSequence",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetDictPageSizeLimit(val *float64) {
	if err := j.validateSetDictPageSizeLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dictPageSizeLimit",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetEnableStatistics(val interface{}) {
	if err := j.validateSetEnableStatisticsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableStatistics",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetEncodingType(val *string) {
	if err := j.validateSetEncodingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encodingType",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetEncryptionMode(val *string) {
	if err := j.validateSetEncryptionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionMode",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetExternalTableDefinition(val *string) {
	if err := j.validateSetExternalTableDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalTableDefinition",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetIgnoreHeaderRows(val *float64) {
	if err := j.validateSetIgnoreHeaderRowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreHeaderRows",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetIncludeOpForFullLoad(val interface{}) {
	if err := j.validateSetIncludeOpForFullLoadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeOpForFullLoad",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetInternalValue(val *DmsEndpointS3Settings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetMaxFileSize(val *float64) {
	if err := j.validateSetMaxFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFileSize",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetParquetTimestampInMillisecond(val interface{}) {
	if err := j.validateSetParquetTimestampInMillisecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parquetTimestampInMillisecond",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetParquetVersion(val *string) {
	if err := j.validateSetParquetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parquetVersion",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetPreserveTransactions(val interface{}) {
	if err := j.validateSetPreserveTransactionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveTransactions",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetRfc4180(val interface{}) {
	if err := j.validateSetRfc4180Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rfc4180",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetRowGroupLength(val *float64) {
	if err := j.validateSetRowGroupLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowGroupLength",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetServerSideEncryptionKmsKeyId(val *string) {
	if err := j.validateSetServerSideEncryptionKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverSideEncryptionKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetServiceAccessRoleArn(val *string) {
	if err := j.validateSetServiceAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetTimestampColumnName(val *string) {
	if err := j.validateSetTimestampColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampColumnName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetUseCsvNoSupValue(val interface{}) {
	if err := j.validateSetUseCsvNoSupValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCsvNoSupValue",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference)SetUseTaskStartTimeForFullLoadTimestamp(val interface{}) {
	if err := j.validateSetUseTaskStartTimeForFullLoadTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useTaskStartTimeForFullLoadTimestamp",
		val,
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetAddColumnName() {
	_jsii_.InvokeVoid(
		d,
		"resetAddColumnName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetBucketFolder() {
	_jsii_.InvokeVoid(
		d,
		"resetBucketFolder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetBucketName() {
	_jsii_.InvokeVoid(
		d,
		"resetBucketName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetCannedAclForObjects() {
	_jsii_.InvokeVoid(
		d,
		"resetCannedAclForObjects",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetCdcInsertsAndUpdates() {
	_jsii_.InvokeVoid(
		d,
		"resetCdcInsertsAndUpdates",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetCdcInsertsOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetCdcInsertsOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetCdcMaxBatchInterval() {
	_jsii_.InvokeVoid(
		d,
		"resetCdcMaxBatchInterval",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetCdcMinFileSize() {
	_jsii_.InvokeVoid(
		d,
		"resetCdcMinFileSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetCdcPath() {
	_jsii_.InvokeVoid(
		d,
		"resetCdcPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetCompressionType() {
	_jsii_.InvokeVoid(
		d,
		"resetCompressionType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetCsvDelimiter() {
	_jsii_.InvokeVoid(
		d,
		"resetCsvDelimiter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetCsvNoSupValue() {
	_jsii_.InvokeVoid(
		d,
		"resetCsvNoSupValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetCsvNullValue() {
	_jsii_.InvokeVoid(
		d,
		"resetCsvNullValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetCsvRowDelimiter() {
	_jsii_.InvokeVoid(
		d,
		"resetCsvRowDelimiter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetDataFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetDataFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetDataPageSize() {
	_jsii_.InvokeVoid(
		d,
		"resetDataPageSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetDatePartitionDelimiter() {
	_jsii_.InvokeVoid(
		d,
		"resetDatePartitionDelimiter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetDatePartitionEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetDatePartitionEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetDatePartitionSequence() {
	_jsii_.InvokeVoid(
		d,
		"resetDatePartitionSequence",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetDictPageSizeLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetDictPageSizeLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetEnableStatistics() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableStatistics",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetEncodingType() {
	_jsii_.InvokeVoid(
		d,
		"resetEncodingType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetEncryptionMode() {
	_jsii_.InvokeVoid(
		d,
		"resetEncryptionMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetExternalTableDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetExternalTableDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetIgnoreHeaderRows() {
	_jsii_.InvokeVoid(
		d,
		"resetIgnoreHeaderRows",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetIncludeOpForFullLoad() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeOpForFullLoad",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetMaxFileSize() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxFileSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetParquetTimestampInMillisecond() {
	_jsii_.InvokeVoid(
		d,
		"resetParquetTimestampInMillisecond",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetParquetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetParquetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetPreserveTransactions() {
	_jsii_.InvokeVoid(
		d,
		"resetPreserveTransactions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetRfc4180() {
	_jsii_.InvokeVoid(
		d,
		"resetRfc4180",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetRowGroupLength() {
	_jsii_.InvokeVoid(
		d,
		"resetRowGroupLength",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetServerSideEncryptionKmsKeyId() {
	_jsii_.InvokeVoid(
		d,
		"resetServerSideEncryptionKmsKeyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetServiceAccessRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccessRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetTimestampColumnName() {
	_jsii_.InvokeVoid(
		d,
		"resetTimestampColumnName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetUseCsvNoSupValue() {
	_jsii_.InvokeVoid(
		d,
		"resetUseCsvNoSupValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetUseTaskStartTimeForFullLoadTimestamp() {
	_jsii_.InvokeVoid(
		d,
		"resetUseTaskStartTimeForFullLoadTimestamp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

