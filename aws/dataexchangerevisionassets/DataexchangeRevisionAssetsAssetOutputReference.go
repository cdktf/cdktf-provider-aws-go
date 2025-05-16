// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataexchangerevisionassets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/dataexchangerevisionassets/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataexchangeRevisionAssetsAssetOutputReference interface {
	cdktf.ComplexObject
	Arn() *string
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
	CreatedAt() *string
	CreateS3DataAccessFromS3Bucket() DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketList
	CreateS3DataAccessFromS3BucketInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Id() *string
	ImportAssetsFromS3() DataexchangeRevisionAssetsAssetImportAssetsFromS3List
	ImportAssetsFromS3Input() interface{}
	ImportAssetsFromSignedUrl() DataexchangeRevisionAssetsAssetImportAssetsFromSignedUrlList
	ImportAssetsFromSignedUrlInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpdatedAt() *string
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
	PutCreateS3DataAccessFromS3Bucket(value interface{})
	PutImportAssetsFromS3(value interface{})
	PutImportAssetsFromSignedUrl(value interface{})
	ResetCreateS3DataAccessFromS3Bucket()
	ResetImportAssetsFromS3()
	ResetImportAssetsFromSignedUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataexchangeRevisionAssetsAssetOutputReference
type jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) CreateS3DataAccessFromS3Bucket() DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketList {
	var returns DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketList
	_jsii_.Get(
		j,
		"createS3DataAccessFromS3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) CreateS3DataAccessFromS3BucketInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createS3DataAccessFromS3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) ImportAssetsFromS3() DataexchangeRevisionAssetsAssetImportAssetsFromS3List {
	var returns DataexchangeRevisionAssetsAssetImportAssetsFromS3List
	_jsii_.Get(
		j,
		"importAssetsFromS3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) ImportAssetsFromS3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importAssetsFromS3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) ImportAssetsFromSignedUrl() DataexchangeRevisionAssetsAssetImportAssetsFromSignedUrlList {
	var returns DataexchangeRevisionAssetsAssetImportAssetsFromSignedUrlList
	_jsii_.Get(
		j,
		"importAssetsFromSignedUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) ImportAssetsFromSignedUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importAssetsFromSignedUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


func NewDataexchangeRevisionAssetsAssetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataexchangeRevisionAssetsAssetOutputReference {
	_init_.Initialize()

	if err := validateNewDataexchangeRevisionAssetsAssetOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataexchangeRevisionAssets.DataexchangeRevisionAssetsAssetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataexchangeRevisionAssetsAssetOutputReference_Override(d DataexchangeRevisionAssetsAssetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataexchangeRevisionAssets.DataexchangeRevisionAssetsAssetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) PutCreateS3DataAccessFromS3Bucket(value interface{}) {
	if err := d.validatePutCreateS3DataAccessFromS3BucketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCreateS3DataAccessFromS3Bucket",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) PutImportAssetsFromS3(value interface{}) {
	if err := d.validatePutImportAssetsFromS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putImportAssetsFromS3",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) PutImportAssetsFromSignedUrl(value interface{}) {
	if err := d.validatePutImportAssetsFromSignedUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putImportAssetsFromSignedUrl",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) ResetCreateS3DataAccessFromS3Bucket() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateS3DataAccessFromS3Bucket",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) ResetImportAssetsFromS3() {
	_jsii_.InvokeVoid(
		d,
		"resetImportAssetsFromS3",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) ResetImportAssetsFromSignedUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetImportAssetsFromSignedUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

