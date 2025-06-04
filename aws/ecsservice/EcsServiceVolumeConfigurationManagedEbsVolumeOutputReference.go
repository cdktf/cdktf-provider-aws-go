// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/ecsservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference interface {
	cdktf.ComplexObject
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
	Encrypted() interface{}
	SetEncrypted(val interface{})
	EncryptedInput() interface{}
	FileSystemType() *string
	SetFileSystemType(val *string)
	FileSystemTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *EcsServiceVolumeConfigurationManagedEbsVolume
	SetInternalValue(val *EcsServiceVolumeConfigurationManagedEbsVolume)
	Iops() *float64
	SetIops(val *float64)
	IopsInput() *float64
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	SizeInGb() *float64
	SetSizeInGb(val *float64)
	SizeInGbInput() *float64
	SnapshotId() *string
	SetSnapshotId(val *string)
	SnapshotIdInput() *string
	TagSpecifications() EcsServiceVolumeConfigurationManagedEbsVolumeTagSpecificationsList
	TagSpecificationsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Throughput() *float64
	SetThroughput(val *float64)
	ThroughputInput() *float64
	VolumeInitializationRate() *float64
	SetVolumeInitializationRate(val *float64)
	VolumeInitializationRateInput() *float64
	VolumeType() *string
	SetVolumeType(val *string)
	VolumeTypeInput() *string
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
	PutTagSpecifications(value interface{})
	ResetEncrypted()
	ResetFileSystemType()
	ResetIops()
	ResetKmsKeyId()
	ResetSizeInGb()
	ResetSnapshotId()
	ResetTagSpecifications()
	ResetThroughput()
	ResetVolumeInitializationRate()
	ResetVolumeType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference
type jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) Encrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) EncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) FileSystemType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) FileSystemTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) InternalValue() *EcsServiceVolumeConfigurationManagedEbsVolume {
	var returns *EcsServiceVolumeConfigurationManagedEbsVolume
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) IopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) SizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) SizeInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) SnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) SnapshotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) TagSpecifications() EcsServiceVolumeConfigurationManagedEbsVolumeTagSpecificationsList {
	var returns EcsServiceVolumeConfigurationManagedEbsVolumeTagSpecificationsList
	_jsii_.Get(
		j,
		"tagSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) TagSpecificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) Throughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) ThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) VolumeInitializationRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeInitializationRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) VolumeInitializationRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeInitializationRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) VolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) VolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeTypeInput",
		&returns,
	)
	return returns
}


func NewEcsServiceVolumeConfigurationManagedEbsVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewEcsServiceVolumeConfigurationManagedEbsVolumeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ecsService.EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcsServiceVolumeConfigurationManagedEbsVolumeOutputReference_Override(e EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ecsService.EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference)SetEncrypted(val interface{}) {
	if err := j.validateSetEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encrypted",
		val,
	)
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference)SetFileSystemType(val *string) {
	if err := j.validateSetFileSystemTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileSystemType",
		val,
	)
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference)SetInternalValue(val *EcsServiceVolumeConfigurationManagedEbsVolume) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference)SetIops(val *float64) {
	if err := j.validateSetIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference)SetSizeInGb(val *float64) {
	if err := j.validateSetSizeInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeInGb",
		val,
	)
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference)SetSnapshotId(val *string) {
	if err := j.validateSetSnapshotIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotId",
		val,
	)
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference)SetThroughput(val *float64) {
	if err := j.validateSetThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throughput",
		val,
	)
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference)SetVolumeInitializationRate(val *float64) {
	if err := j.validateSetVolumeInitializationRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeInitializationRate",
		val,
	)
}

func (j *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference)SetVolumeType(val *string) {
	if err := j.validateSetVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeType",
		val,
	)
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) PutTagSpecifications(value interface{}) {
	if err := e.validatePutTagSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTagSpecifications",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) ResetEncrypted() {
	_jsii_.InvokeVoid(
		e,
		"resetEncrypted",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) ResetFileSystemType() {
	_jsii_.InvokeVoid(
		e,
		"resetFileSystemType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) ResetIops() {
	_jsii_.InvokeVoid(
		e,
		"resetIops",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		e,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) ResetSizeInGb() {
	_jsii_.InvokeVoid(
		e,
		"resetSizeInGb",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) ResetSnapshotId() {
	_jsii_.InvokeVoid(
		e,
		"resetSnapshotId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) ResetTagSpecifications() {
	_jsii_.InvokeVoid(
		e,
		"resetTagSpecifications",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) ResetThroughput() {
	_jsii_.InvokeVoid(
		e,
		"resetThroughput",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) ResetVolumeInitializationRate() {
	_jsii_.InvokeVoid(
		e,
		"resetVolumeInitializationRate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) ResetVolumeType() {
	_jsii_.InvokeVoid(
		e,
		"resetVolumeType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceVolumeConfigurationManagedEbsVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

