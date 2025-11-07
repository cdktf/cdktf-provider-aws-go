// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxs3accesspointattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/fsxs3accesspointattachment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference interface {
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
	// Experimental.
	Fqn() *string
	Gid() *float64
	SetGid(val *float64)
	GidInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SecondaryGids() *[]*float64
	SetSecondaryGids(val *[]*float64)
	SecondaryGidsInput() *[]*float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uid() *float64
	SetUid(val *float64)
	UidInput() *float64
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
	ResetSecondaryGids()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference
type jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) Gid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) GidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) SecondaryGids() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"secondaryGids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) SecondaryGidsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"secondaryGidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) Uid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) UidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uidInput",
		&returns,
	)
	return returns
}


func NewFsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference {
	_init_.Initialize()

	if err := validateNewFsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.fsxS3AccessPointAttachment.FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference_Override(f FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.fsxS3AccessPointAttachment.FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference)SetGid(val *float64) {
	if err := j.validateSetGidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gid",
		val,
	)
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference)SetSecondaryGids(val *[]*float64) {
	if err := j.validateSetSecondaryGidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryGids",
		val,
	)
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference)SetUid(val *float64) {
	if err := j.validateSetUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uid",
		val,
	)
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) ResetSecondaryGids() {
	_jsii_.InvokeVoid(
		f,
		"resetSecondaryGids",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxS3AccessPointAttachmentOpenzfsConfigurationFileSystemIdentityPosixUserOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

