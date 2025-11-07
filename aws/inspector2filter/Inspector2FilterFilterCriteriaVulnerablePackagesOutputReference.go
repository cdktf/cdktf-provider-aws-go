// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inspector2filter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/inspector2filter/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference interface {
	cdktf.ComplexObject
	Architecture() Inspector2FilterFilterCriteriaVulnerablePackagesArchitectureList
	ArchitectureInput() interface{}
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
	Epoch() Inspector2FilterFilterCriteriaVulnerablePackagesEpochList
	EpochInput() interface{}
	FilePath() Inspector2FilterFilterCriteriaVulnerablePackagesFilePathList
	FilePathInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() Inspector2FilterFilterCriteriaVulnerablePackagesNameList
	NameInput() interface{}
	Release() Inspector2FilterFilterCriteriaVulnerablePackagesReleaseList
	ReleaseInput() interface{}
	SourceLambdaLayerArn() Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList
	SourceLambdaLayerArnInput() interface{}
	SourceLayerHash() Inspector2FilterFilterCriteriaVulnerablePackagesSourceLayerHashList
	SourceLayerHashInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() Inspector2FilterFilterCriteriaVulnerablePackagesVersionList
	VersionInput() interface{}
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
	PutArchitecture(value interface{})
	PutEpoch(value interface{})
	PutFilePath(value interface{})
	PutName(value interface{})
	PutRelease(value interface{})
	PutSourceLambdaLayerArn(value interface{})
	PutSourceLayerHash(value interface{})
	PutVersion(value interface{})
	ResetArchitecture()
	ResetEpoch()
	ResetFilePath()
	ResetName()
	ResetRelease()
	ResetSourceLambdaLayerArn()
	ResetSourceLayerHash()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference
type jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) Architecture() Inspector2FilterFilterCriteriaVulnerablePackagesArchitectureList {
	var returns Inspector2FilterFilterCriteriaVulnerablePackagesArchitectureList
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) ArchitectureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"architectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) Epoch() Inspector2FilterFilterCriteriaVulnerablePackagesEpochList {
	var returns Inspector2FilterFilterCriteriaVulnerablePackagesEpochList
	_jsii_.Get(
		j,
		"epoch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) EpochInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"epochInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) FilePath() Inspector2FilterFilterCriteriaVulnerablePackagesFilePathList {
	var returns Inspector2FilterFilterCriteriaVulnerablePackagesFilePathList
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) FilePathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) Name() Inspector2FilterFilterCriteriaVulnerablePackagesNameList {
	var returns Inspector2FilterFilterCriteriaVulnerablePackagesNameList
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) NameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) Release() Inspector2FilterFilterCriteriaVulnerablePackagesReleaseList {
	var returns Inspector2FilterFilterCriteriaVulnerablePackagesReleaseList
	_jsii_.Get(
		j,
		"release",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) ReleaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"releaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) SourceLambdaLayerArn() Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList {
	var returns Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArnList
	_jsii_.Get(
		j,
		"sourceLambdaLayerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) SourceLambdaLayerArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceLambdaLayerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) SourceLayerHash() Inspector2FilterFilterCriteriaVulnerablePackagesSourceLayerHashList {
	var returns Inspector2FilterFilterCriteriaVulnerablePackagesSourceLayerHashList
	_jsii_.Get(
		j,
		"sourceLayerHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) SourceLayerHashInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceLayerHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) Version() Inspector2FilterFilterCriteriaVulnerablePackagesVersionList {
	var returns Inspector2FilterFilterCriteriaVulnerablePackagesVersionList
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) VersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewInspector2FilterFilterCriteriaVulnerablePackagesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference {
	_init_.Initialize()

	if err := validateNewInspector2FilterFilterCriteriaVulnerablePackagesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.inspector2Filter.Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewInspector2FilterFilterCriteriaVulnerablePackagesOutputReference_Override(i Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.inspector2Filter.Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) PutArchitecture(value interface{}) {
	if err := i.validatePutArchitectureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putArchitecture",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) PutEpoch(value interface{}) {
	if err := i.validatePutEpochParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEpoch",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) PutFilePath(value interface{}) {
	if err := i.validatePutFilePathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFilePath",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) PutName(value interface{}) {
	if err := i.validatePutNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putName",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) PutRelease(value interface{}) {
	if err := i.validatePutReleaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putRelease",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) PutSourceLambdaLayerArn(value interface{}) {
	if err := i.validatePutSourceLambdaLayerArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSourceLambdaLayerArn",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) PutSourceLayerHash(value interface{}) {
	if err := i.validatePutSourceLayerHashParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSourceLayerHash",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) PutVersion(value interface{}) {
	if err := i.validatePutVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putVersion",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) ResetArchitecture() {
	_jsii_.InvokeVoid(
		i,
		"resetArchitecture",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) ResetEpoch() {
	_jsii_.InvokeVoid(
		i,
		"resetEpoch",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) ResetFilePath() {
	_jsii_.InvokeVoid(
		i,
		"resetFilePath",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		i,
		"resetName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) ResetRelease() {
	_jsii_.InvokeVoid(
		i,
		"resetRelease",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) ResetSourceLambdaLayerArn() {
	_jsii_.InvokeVoid(
		i,
		"resetSourceLambdaLayerArn",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) ResetSourceLayerHash() {
	_jsii_.InvokeVoid(
		i,
		"resetSourceLayerHash",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		i,
		"resetVersion",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaVulnerablePackagesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

