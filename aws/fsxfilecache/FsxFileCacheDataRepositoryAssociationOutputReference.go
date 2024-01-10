// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxfilecache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/fsxfilecache/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FsxFileCacheDataRepositoryAssociationOutputReference interface {
	cdktf.ComplexObject
	AssociationId() *string
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
	DataRepositoryPath() *string
	SetDataRepositoryPath(val *string)
	DataRepositoryPathInput() *string
	DataRepositorySubdirectories() *[]*string
	SetDataRepositorySubdirectories(val *[]*string)
	DataRepositorySubdirectoriesInput() *[]*string
	FileCacheId() *string
	FileCachePath() *string
	SetFileCachePath(val *string)
	FileCachePathInput() *string
	FileSystemId() *string
	FileSystemPath() *string
	// Experimental.
	Fqn() *string
	ImportedFileChunkSize() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Nfs() FsxFileCacheDataRepositoryAssociationNfsList
	NfsInput() interface{}
	ResourceArn() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
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
	PutNfs(value interface{})
	ResetDataRepositorySubdirectories()
	ResetNfs()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FsxFileCacheDataRepositoryAssociationOutputReference
type jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) AssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) DataRepositoryPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataRepositoryPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) DataRepositoryPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataRepositoryPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) DataRepositorySubdirectories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataRepositorySubdirectories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) DataRepositorySubdirectoriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataRepositorySubdirectoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) FileCacheId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileCacheId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) FileCachePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileCachePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) FileCachePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileCachePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) FileSystemPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) ImportedFileChunkSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importedFileChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) Nfs() FsxFileCacheDataRepositoryAssociationNfsList {
	var returns FsxFileCacheDataRepositoryAssociationNfsList
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) NfsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFsxFileCacheDataRepositoryAssociationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FsxFileCacheDataRepositoryAssociationOutputReference {
	_init_.Initialize()

	if err := validateNewFsxFileCacheDataRepositoryAssociationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.fsxFileCache.FsxFileCacheDataRepositoryAssociationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFsxFileCacheDataRepositoryAssociationOutputReference_Override(f FsxFileCacheDataRepositoryAssociationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.fsxFileCache.FsxFileCacheDataRepositoryAssociationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference)SetDataRepositoryPath(val *string) {
	if err := j.validateSetDataRepositoryPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataRepositoryPath",
		val,
	)
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference)SetDataRepositorySubdirectories(val *[]*string) {
	if err := j.validateSetDataRepositorySubdirectoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataRepositorySubdirectories",
		val,
	)
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference)SetFileCachePath(val *string) {
	if err := j.validateSetFileCachePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileCachePath",
		val,
	)
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) PutNfs(value interface{}) {
	if err := f.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putNfs",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) ResetDataRepositorySubdirectories() {
	_jsii_.InvokeVoid(
		f,
		"resetDataRepositorySubdirectories",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		f,
		"resetNfs",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxFileCacheDataRepositoryAssociationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

