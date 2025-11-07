// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataexchangerevisionassets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/dataexchangerevisionassets/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference interface {
	cdktf.ComplexObject
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyPrefixes() *[]*string
	SetKeyPrefixes(val *[]*string)
	KeyPrefixesInput() *[]*string
	Keys() *[]*string
	SetKeys(val *[]*string)
	KeysInput() *[]*string
	KmsKeysToGrant() DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceKmsKeysToGrantList
	KmsKeysToGrantInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutKmsKeysToGrant(value interface{})
	ResetKeyPrefixes()
	ResetKeys()
	ResetKmsKeysToGrant()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference
type jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) KeyPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) KeyPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) Keys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) KeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) KmsKeysToGrant() DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceKmsKeysToGrantList {
	var returns DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceKmsKeysToGrantList
	_jsii_.Get(
		j,
		"kmsKeysToGrant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) KmsKeysToGrantInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsKeysToGrantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference {
	_init_.Initialize()

	if err := validateNewDataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataexchangeRevisionAssets.DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference_Override(d DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataexchangeRevisionAssets.DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference)SetKeyPrefixes(val *[]*string) {
	if err := j.validateSetKeyPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPrefixes",
		val,
	)
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference)SetKeys(val *[]*string) {
	if err := j.validateSetKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keys",
		val,
	)
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) PutKmsKeysToGrant(value interface{}) {
	if err := d.validatePutKmsKeysToGrantParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKmsKeysToGrant",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) ResetKeyPrefixes() {
	_jsii_.InvokeVoid(
		d,
		"resetKeyPrefixes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) ResetKeys() {
	_jsii_.InvokeVoid(
		d,
		"resetKeys",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) ResetKmsKeysToGrant() {
	_jsii_.InvokeVoid(
		d,
		"resetKmsKeysToGrant",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

