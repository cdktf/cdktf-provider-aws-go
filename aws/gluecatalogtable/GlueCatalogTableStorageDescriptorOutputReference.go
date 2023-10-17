// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/gluecatalogtable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlueCatalogTableStorageDescriptorOutputReference interface {
	cdktf.ComplexObject
	BucketColumns() *[]*string
	SetBucketColumns(val *[]*string)
	BucketColumnsInput() *[]*string
	Columns() GlueCatalogTableStorageDescriptorColumnsList
	ColumnsInput() interface{}
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
	Compressed() interface{}
	SetCompressed(val interface{})
	CompressedInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InputFormat() *string
	SetInputFormat(val *string)
	InputFormatInput() *string
	InternalValue() *GlueCatalogTableStorageDescriptor
	SetInternalValue(val *GlueCatalogTableStorageDescriptor)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	NumberOfBuckets() *float64
	SetNumberOfBuckets(val *float64)
	NumberOfBucketsInput() *float64
	OutputFormat() *string
	SetOutputFormat(val *string)
	OutputFormatInput() *string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	SchemaReference() GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference
	SchemaReferenceInput() *GlueCatalogTableStorageDescriptorSchemaReference
	SerDeInfo() GlueCatalogTableStorageDescriptorSerDeInfoOutputReference
	SerDeInfoInput() *GlueCatalogTableStorageDescriptorSerDeInfo
	SkewedInfo() GlueCatalogTableStorageDescriptorSkewedInfoOutputReference
	SkewedInfoInput() *GlueCatalogTableStorageDescriptorSkewedInfo
	SortColumns() GlueCatalogTableStorageDescriptorSortColumnsList
	SortColumnsInput() interface{}
	StoredAsSubDirectories() interface{}
	SetStoredAsSubDirectories(val interface{})
	StoredAsSubDirectoriesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutColumns(value interface{})
	PutSchemaReference(value *GlueCatalogTableStorageDescriptorSchemaReference)
	PutSerDeInfo(value *GlueCatalogTableStorageDescriptorSerDeInfo)
	PutSkewedInfo(value *GlueCatalogTableStorageDescriptorSkewedInfo)
	PutSortColumns(value interface{})
	ResetBucketColumns()
	ResetColumns()
	ResetCompressed()
	ResetInputFormat()
	ResetLocation()
	ResetNumberOfBuckets()
	ResetOutputFormat()
	ResetParameters()
	ResetSchemaReference()
	ResetSerDeInfo()
	ResetSkewedInfo()
	ResetSortColumns()
	ResetStoredAsSubDirectories()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueCatalogTableStorageDescriptorOutputReference
type jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) BucketColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) BucketColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) Columns() GlueCatalogTableStorageDescriptorColumnsList {
	var returns GlueCatalogTableStorageDescriptorColumnsList
	_jsii_.Get(
		j,
		"columns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) Compressed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) CompressedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) InputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) InputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) InternalValue() *GlueCatalogTableStorageDescriptor {
	var returns *GlueCatalogTableStorageDescriptor
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) NumberOfBuckets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) NumberOfBucketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) OutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) OutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SchemaReference() GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference {
	var returns GlueCatalogTableStorageDescriptorSchemaReferenceOutputReference
	_jsii_.Get(
		j,
		"schemaReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SchemaReferenceInput() *GlueCatalogTableStorageDescriptorSchemaReference {
	var returns *GlueCatalogTableStorageDescriptorSchemaReference
	_jsii_.Get(
		j,
		"schemaReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SerDeInfo() GlueCatalogTableStorageDescriptorSerDeInfoOutputReference {
	var returns GlueCatalogTableStorageDescriptorSerDeInfoOutputReference
	_jsii_.Get(
		j,
		"serDeInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SerDeInfoInput() *GlueCatalogTableStorageDescriptorSerDeInfo {
	var returns *GlueCatalogTableStorageDescriptorSerDeInfo
	_jsii_.Get(
		j,
		"serDeInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SkewedInfo() GlueCatalogTableStorageDescriptorSkewedInfoOutputReference {
	var returns GlueCatalogTableStorageDescriptorSkewedInfoOutputReference
	_jsii_.Get(
		j,
		"skewedInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SkewedInfoInput() *GlueCatalogTableStorageDescriptorSkewedInfo {
	var returns *GlueCatalogTableStorageDescriptorSkewedInfo
	_jsii_.Get(
		j,
		"skewedInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SortColumns() GlueCatalogTableStorageDescriptorSortColumnsList {
	var returns GlueCatalogTableStorageDescriptorSortColumnsList
	_jsii_.Get(
		j,
		"sortColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) SortColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sortColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) StoredAsSubDirectories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storedAsSubDirectories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) StoredAsSubDirectoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storedAsSubDirectoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueCatalogTableStorageDescriptorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GlueCatalogTableStorageDescriptorOutputReference {
	_init_.Initialize()

	if err := validateNewGlueCatalogTableStorageDescriptorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.glueCatalogTable.GlueCatalogTableStorageDescriptorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueCatalogTableStorageDescriptorOutputReference_Override(g GlueCatalogTableStorageDescriptorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.glueCatalogTable.GlueCatalogTableStorageDescriptorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference)SetBucketColumns(val *[]*string) {
	if err := j.validateSetBucketColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketColumns",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference)SetCompressed(val interface{}) {
	if err := j.validateSetCompressedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressed",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference)SetInputFormat(val *string) {
	if err := j.validateSetInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputFormat",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference)SetInternalValue(val *GlueCatalogTableStorageDescriptor) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference)SetNumberOfBuckets(val *float64) {
	if err := j.validateSetNumberOfBucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfBuckets",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference)SetOutputFormat(val *string) {
	if err := j.validateSetOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputFormat",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference)SetStoredAsSubDirectories(val interface{}) {
	if err := j.validateSetStoredAsSubDirectoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storedAsSubDirectories",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) PutColumns(value interface{}) {
	if err := g.validatePutColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putColumns",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) PutSchemaReference(value *GlueCatalogTableStorageDescriptorSchemaReference) {
	if err := g.validatePutSchemaReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSchemaReference",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) PutSerDeInfo(value *GlueCatalogTableStorageDescriptorSerDeInfo) {
	if err := g.validatePutSerDeInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSerDeInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) PutSkewedInfo(value *GlueCatalogTableStorageDescriptorSkewedInfo) {
	if err := g.validatePutSkewedInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSkewedInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) PutSortColumns(value interface{}) {
	if err := g.validatePutSortColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSortColumns",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetBucketColumns() {
	_jsii_.InvokeVoid(
		g,
		"resetBucketColumns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetColumns() {
	_jsii_.InvokeVoid(
		g,
		"resetColumns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetCompressed() {
	_jsii_.InvokeVoid(
		g,
		"resetCompressed",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetInputFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetInputFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetNumberOfBuckets() {
	_jsii_.InvokeVoid(
		g,
		"resetNumberOfBuckets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetOutputFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetOutputFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetSchemaReference() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaReference",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetSerDeInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetSerDeInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetSkewedInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetSkewedInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetSortColumns() {
	_jsii_.InvokeVoid(
		g,
		"resetSortColumns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ResetStoredAsSubDirectories() {
	_jsii_.InvokeVoid(
		g,
		"resetStoredAsSubDirectories",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

