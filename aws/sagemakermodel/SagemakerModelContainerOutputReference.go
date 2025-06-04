// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakermodel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/sagemakermodel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerModelContainerOutputReference interface {
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
	ContainerHostname() *string
	SetContainerHostname(val *string)
	ContainerHostnameInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Environment() *map[string]*string
	SetEnvironment(val *map[string]*string)
	EnvironmentInput() *map[string]*string
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageConfig() SagemakerModelContainerImageConfigOutputReference
	ImageConfigInput() *SagemakerModelContainerImageConfig
	ImageInput() *string
	InferenceSpecificationName() *string
	SetInferenceSpecificationName(val *string)
	InferenceSpecificationNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	ModelDataSource() SagemakerModelContainerModelDataSourceOutputReference
	ModelDataSourceInput() *SagemakerModelContainerModelDataSource
	ModelDataUrl() *string
	SetModelDataUrl(val *string)
	ModelDataUrlInput() *string
	ModelPackageName() *string
	SetModelPackageName(val *string)
	ModelPackageNameInput() *string
	MultiModelConfig() SagemakerModelContainerMultiModelConfigOutputReference
	MultiModelConfigInput() *SagemakerModelContainerMultiModelConfig
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
	PutImageConfig(value *SagemakerModelContainerImageConfig)
	PutModelDataSource(value *SagemakerModelContainerModelDataSource)
	PutMultiModelConfig(value *SagemakerModelContainerMultiModelConfig)
	ResetContainerHostname()
	ResetEnvironment()
	ResetImage()
	ResetImageConfig()
	ResetInferenceSpecificationName()
	ResetMode()
	ResetModelDataSource()
	ResetModelDataUrl()
	ResetModelPackageName()
	ResetMultiModelConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerModelContainerOutputReference
type jsiiProxy_SagemakerModelContainerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) ContainerHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) ContainerHostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerHostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) EnvironmentInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) ImageConfig() SagemakerModelContainerImageConfigOutputReference {
	var returns SagemakerModelContainerImageConfigOutputReference
	_jsii_.Get(
		j,
		"imageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) ImageConfigInput() *SagemakerModelContainerImageConfig {
	var returns *SagemakerModelContainerImageConfig
	_jsii_.Get(
		j,
		"imageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) InferenceSpecificationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceSpecificationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) InferenceSpecificationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceSpecificationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) ModelDataSource() SagemakerModelContainerModelDataSourceOutputReference {
	var returns SagemakerModelContainerModelDataSourceOutputReference
	_jsii_.Get(
		j,
		"modelDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) ModelDataSourceInput() *SagemakerModelContainerModelDataSource {
	var returns *SagemakerModelContainerModelDataSource
	_jsii_.Get(
		j,
		"modelDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) ModelDataUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDataUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) ModelDataUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDataUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) ModelPackageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) ModelPackageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) MultiModelConfig() SagemakerModelContainerMultiModelConfigOutputReference {
	var returns SagemakerModelContainerMultiModelConfigOutputReference
	_jsii_.Get(
		j,
		"multiModelConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) MultiModelConfigInput() *SagemakerModelContainerMultiModelConfig {
	var returns *SagemakerModelContainerMultiModelConfig
	_jsii_.Get(
		j,
		"multiModelConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerModelContainerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerModelContainerOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerModelContainerOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerModelContainerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerModel.SagemakerModelContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerModelContainerOutputReference_Override(s SagemakerModelContainerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerModel.SagemakerModelContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference)SetContainerHostname(val *string) {
	if err := j.validateSetContainerHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerHostname",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference)SetEnvironment(val *map[string]*string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference)SetInferenceSpecificationName(val *string) {
	if err := j.validateSetInferenceSpecificationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceSpecificationName",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference)SetModelDataUrl(val *string) {
	if err := j.validateSetModelDataUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDataUrl",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference)SetModelPackageName(val *string) {
	if err := j.validateSetModelPackageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelPackageName",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelContainerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) PutImageConfig(value *SagemakerModelContainerImageConfig) {
	if err := s.validatePutImageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putImageConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) PutModelDataSource(value *SagemakerModelContainerModelDataSource) {
	if err := s.validatePutModelDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelDataSource",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) PutMultiModelConfig(value *SagemakerModelContainerMultiModelConfig) {
	if err := s.validatePutMultiModelConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMultiModelConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) ResetContainerHostname() {
	_jsii_.InvokeVoid(
		s,
		"resetContainerHostname",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		s,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		s,
		"resetImage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) ResetImageConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetImageConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) ResetInferenceSpecificationName() {
	_jsii_.InvokeVoid(
		s,
		"resetInferenceSpecificationName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		s,
		"resetMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) ResetModelDataSource() {
	_jsii_.InvokeVoid(
		s,
		"resetModelDataSource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) ResetModelDataUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetModelDataUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) ResetModelPackageName() {
	_jsii_.InvokeVoid(
		s,
		"resetModelPackageName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) ResetMultiModelConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetMultiModelConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelContainerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

