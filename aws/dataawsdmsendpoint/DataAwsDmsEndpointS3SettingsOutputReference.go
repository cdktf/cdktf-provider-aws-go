// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsdmsendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/dataawsdmsendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsDmsEndpointS3SettingsOutputReference interface {
	cdktf.ComplexObject
	AddColumnName() cdktf.IResolvable
	BucketFolder() *string
	BucketName() *string
	CannedAclForObjects() *string
	CdcInsertsAndUpdates() cdktf.IResolvable
	CdcInsertsOnly() cdktf.IResolvable
	CdcMaxBatchInterval() *float64
	CdcMinFileSize() *float64
	CdcPath() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CsvDelimiter() *string
	CsvNoSupValue() *string
	CsvNullValue() *string
	CsvRowDelimiter() *string
	DataFormat() *string
	DataPageSize() *float64
	DatePartitionDelimiter() *string
	DatePartitionEnabled() cdktf.IResolvable
	DatePartitionSequence() *string
	DictPageSizeLimit() *float64
	EnableStatistics() cdktf.IResolvable
	EncodingType() *string
	EncryptionMode() *string
	ExternalTableDefinition() *string
	// Experimental.
	Fqn() *string
	GlueCatalogGeneration() cdktf.IResolvable
	IgnoreHeaderRows() *float64
	IgnoreHeadersRow() *float64
	IncludeOpForFullLoad() cdktf.IResolvable
	InternalValue() *DataAwsDmsEndpointS3Settings
	SetInternalValue(val *DataAwsDmsEndpointS3Settings)
	MaxFileSize() *float64
	ParquetTimestampInMillisecond() cdktf.IResolvable
	ParquetVersion() *string
	PreserveTransactions() cdktf.IResolvable
	Rfc4180() cdktf.IResolvable
	RowGroupLength() *float64
	ServerSideEncryptionKmsKeyId() *string
	ServiceAccessRoleArn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimestampColumnName() *string
	UseCsvNoSupValue() cdktf.IResolvable
	UseTaskStartTimeForFullLoadTimestamp() cdktf.IResolvable
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsDmsEndpointS3SettingsOutputReference
type jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) AddColumnName() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"addColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) BucketFolder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketFolder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) CannedAclForObjects() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cannedAclForObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) CdcInsertsAndUpdates() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"cdcInsertsAndUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) CdcInsertsOnly() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"cdcInsertsOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) CdcMaxBatchInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdcMaxBatchInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) CdcMinFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdcMinFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) CdcPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdcPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) CsvDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) CsvNoSupValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvNoSupValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) CsvNullValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvNullValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) CsvRowDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvRowDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) DataFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) DataPageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataPageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) DatePartitionDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datePartitionDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) DatePartitionEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"datePartitionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) DatePartitionSequence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datePartitionSequence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) DictPageSizeLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dictPageSizeLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) EnableStatistics() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableStatistics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) EncodingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) EncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) ExternalTableDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalTableDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) GlueCatalogGeneration() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"glueCatalogGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) IgnoreHeaderRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ignoreHeaderRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) IgnoreHeadersRow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ignoreHeadersRow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) IncludeOpForFullLoad() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"includeOpForFullLoad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) InternalValue() *DataAwsDmsEndpointS3Settings {
	var returns *DataAwsDmsEndpointS3Settings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) ParquetTimestampInMillisecond() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"parquetTimestampInMillisecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) ParquetVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parquetVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) PreserveTransactions() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"preserveTransactions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) Rfc4180() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"rfc4180",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) RowGroupLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowGroupLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) ServerSideEncryptionKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) TimestampColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) UseCsvNoSupValue() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"useCsvNoSupValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) UseTaskStartTimeForFullLoadTimestamp() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"useTaskStartTimeForFullLoadTimestamp",
		&returns,
	)
	return returns
}


func NewDataAwsDmsEndpointS3SettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsDmsEndpointS3SettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsDmsEndpointS3SettingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsDmsEndpoint.DataAwsDmsEndpointS3SettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsDmsEndpointS3SettingsOutputReference_Override(d DataAwsDmsEndpointS3SettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsDmsEndpoint.DataAwsDmsEndpointS3SettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference)SetInternalValue(val *DataAwsDmsEndpointS3Settings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsDmsEndpointS3SettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

