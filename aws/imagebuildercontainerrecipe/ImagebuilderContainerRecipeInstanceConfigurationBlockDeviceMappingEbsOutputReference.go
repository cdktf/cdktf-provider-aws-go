// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuildercontainerrecipe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/imagebuildercontainerrecipe/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference interface {
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
	DeleteOnTermination() *string
	SetDeleteOnTermination(val *string)
	DeleteOnTerminationInput() *string
	Encrypted() *string
	SetEncrypted(val *string)
	EncryptedInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbs
	SetInternalValue(val *ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbs)
	Iops() *float64
	SetIops(val *float64)
	IopsInput() *float64
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	SnapshotId() *string
	SetSnapshotId(val *string)
	SnapshotIdInput() *string
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
	VolumeSize() *float64
	SetVolumeSize(val *float64)
	VolumeSizeInput() *float64
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
	ResetDeleteOnTermination()
	ResetEncrypted()
	ResetIops()
	ResetKmsKeyId()
	ResetSnapshotId()
	ResetThroughput()
	ResetVolumeSize()
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

// The jsii proxy struct for ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference
type jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) DeleteOnTermination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteOnTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) DeleteOnTerminationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteOnTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) Encrypted() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) EncryptedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) InternalValue() *ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbs {
	var returns *ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) IopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) SnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) SnapshotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) Throughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) ThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) VolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) VolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) VolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) VolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeTypeInput",
		&returns,
	)
	return returns
}


func NewImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference {
	_init_.Initialize()

	if err := validateNewImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.imagebuilderContainerRecipe.ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference_Override(i ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.imagebuilderContainerRecipe.ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference)SetDeleteOnTermination(val *string) {
	if err := j.validateSetDeleteOnTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteOnTermination",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference)SetEncrypted(val *string) {
	if err := j.validateSetEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encrypted",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference)SetInternalValue(val *ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbs) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference)SetIops(val *float64) {
	if err := j.validateSetIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference)SetSnapshotId(val *string) {
	if err := j.validateSetSnapshotIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotId",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference)SetThroughput(val *float64) {
	if err := j.validateSetThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throughput",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference)SetVolumeSize(val *float64) {
	if err := j.validateSetVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeSize",
		val,
	)
}

func (j *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference)SetVolumeType(val *string) {
	if err := j.validateSetVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeType",
		val,
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) ResetDeleteOnTermination() {
	_jsii_.InvokeVoid(
		i,
		"resetDeleteOnTermination",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) ResetEncrypted() {
	_jsii_.InvokeVoid(
		i,
		"resetEncrypted",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) ResetIops() {
	_jsii_.InvokeVoid(
		i,
		"resetIops",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		i,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) ResetSnapshotId() {
	_jsii_.InvokeVoid(
		i,
		"resetSnapshotId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) ResetThroughput() {
	_jsii_.InvokeVoid(
		i,
		"resetThroughput",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) ResetVolumeSize() {
	_jsii_.InvokeVoid(
		i,
		"resetVolumeSize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) ResetVolumeType() {
	_jsii_.InvokeVoid(
		i,
		"resetVolumeType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

