// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/sagemakerendpointconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerEndpointConfigurationProductionVariantsOutputReference interface {
	cdktf.ComplexObject
	AcceleratorType() *string
	SetAcceleratorType(val *string)
	AcceleratorTypeInput() *string
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
	ContainerStartupHealthCheckTimeoutInSeconds() *float64
	SetContainerStartupHealthCheckTimeoutInSeconds(val *float64)
	ContainerStartupHealthCheckTimeoutInSecondsInput() *float64
	CoreDumpConfig() SagemakerEndpointConfigurationProductionVariantsCoreDumpConfigOutputReference
	CoreDumpConfigInput() *SagemakerEndpointConfigurationProductionVariantsCoreDumpConfig
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnableSsmAccess() interface{}
	SetEnableSsmAccess(val interface{})
	EnableSsmAccessInput() interface{}
	// Experimental.
	Fqn() *string
	InferenceAmiVersion() *string
	SetInferenceAmiVersion(val *string)
	InferenceAmiVersionInput() *string
	InitialInstanceCount() *float64
	SetInitialInstanceCount(val *float64)
	InitialInstanceCountInput() *float64
	InitialVariantWeight() *float64
	SetInitialVariantWeight(val *float64)
	InitialVariantWeightInput() *float64
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ManagedInstanceScaling() SagemakerEndpointConfigurationProductionVariantsManagedInstanceScalingOutputReference
	ManagedInstanceScalingInput() *SagemakerEndpointConfigurationProductionVariantsManagedInstanceScaling
	ModelDataDownloadTimeoutInSeconds() *float64
	SetModelDataDownloadTimeoutInSeconds(val *float64)
	ModelDataDownloadTimeoutInSecondsInput() *float64
	ModelName() *string
	SetModelName(val *string)
	ModelNameInput() *string
	RoutingConfig() SagemakerEndpointConfigurationProductionVariantsRoutingConfigList
	RoutingConfigInput() interface{}
	ServerlessConfig() SagemakerEndpointConfigurationProductionVariantsServerlessConfigOutputReference
	ServerlessConfigInput() *SagemakerEndpointConfigurationProductionVariantsServerlessConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VariantName() *string
	SetVariantName(val *string)
	VariantNameInput() *string
	VolumeSizeInGb() *float64
	SetVolumeSizeInGb(val *float64)
	VolumeSizeInGbInput() *float64
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
	PutCoreDumpConfig(value *SagemakerEndpointConfigurationProductionVariantsCoreDumpConfig)
	PutManagedInstanceScaling(value *SagemakerEndpointConfigurationProductionVariantsManagedInstanceScaling)
	PutRoutingConfig(value interface{})
	PutServerlessConfig(value *SagemakerEndpointConfigurationProductionVariantsServerlessConfig)
	ResetAcceleratorType()
	ResetContainerStartupHealthCheckTimeoutInSeconds()
	ResetCoreDumpConfig()
	ResetEnableSsmAccess()
	ResetInferenceAmiVersion()
	ResetInitialInstanceCount()
	ResetInitialVariantWeight()
	ResetInstanceType()
	ResetManagedInstanceScaling()
	ResetModelDataDownloadTimeoutInSeconds()
	ResetRoutingConfig()
	ResetServerlessConfig()
	ResetVariantName()
	ResetVolumeSizeInGb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerEndpointConfigurationProductionVariantsOutputReference
type jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) AcceleratorType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) AcceleratorTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ContainerStartupHealthCheckTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerStartupHealthCheckTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ContainerStartupHealthCheckTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerStartupHealthCheckTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) CoreDumpConfig() SagemakerEndpointConfigurationProductionVariantsCoreDumpConfigOutputReference {
	var returns SagemakerEndpointConfigurationProductionVariantsCoreDumpConfigOutputReference
	_jsii_.Get(
		j,
		"coreDumpConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) CoreDumpConfigInput() *SagemakerEndpointConfigurationProductionVariantsCoreDumpConfig {
	var returns *SagemakerEndpointConfigurationProductionVariantsCoreDumpConfig
	_jsii_.Get(
		j,
		"coreDumpConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) EnableSsmAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSsmAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) EnableSsmAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSsmAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) InferenceAmiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceAmiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) InferenceAmiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceAmiVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) InitialInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) InitialInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) InitialVariantWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialVariantWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) InitialVariantWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialVariantWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ManagedInstanceScaling() SagemakerEndpointConfigurationProductionVariantsManagedInstanceScalingOutputReference {
	var returns SagemakerEndpointConfigurationProductionVariantsManagedInstanceScalingOutputReference
	_jsii_.Get(
		j,
		"managedInstanceScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ManagedInstanceScalingInput() *SagemakerEndpointConfigurationProductionVariantsManagedInstanceScaling {
	var returns *SagemakerEndpointConfigurationProductionVariantsManagedInstanceScaling
	_jsii_.Get(
		j,
		"managedInstanceScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ModelDataDownloadTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modelDataDownloadTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ModelDataDownloadTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modelDataDownloadTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) RoutingConfig() SagemakerEndpointConfigurationProductionVariantsRoutingConfigList {
	var returns SagemakerEndpointConfigurationProductionVariantsRoutingConfigList
	_jsii_.Get(
		j,
		"routingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) RoutingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ServerlessConfig() SagemakerEndpointConfigurationProductionVariantsServerlessConfigOutputReference {
	var returns SagemakerEndpointConfigurationProductionVariantsServerlessConfigOutputReference
	_jsii_.Get(
		j,
		"serverlessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ServerlessConfigInput() *SagemakerEndpointConfigurationProductionVariantsServerlessConfig {
	var returns *SagemakerEndpointConfigurationProductionVariantsServerlessConfig
	_jsii_.Get(
		j,
		"serverlessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) VariantName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variantName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) VariantNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variantNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) VolumeSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) VolumeSizeInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInGbInput",
		&returns,
	)
	return returns
}


func NewSagemakerEndpointConfigurationProductionVariantsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerEndpointConfigurationProductionVariantsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerEndpointConfigurationProductionVariantsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerEndpointConfiguration.SagemakerEndpointConfigurationProductionVariantsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerEndpointConfigurationProductionVariantsOutputReference_Override(s SagemakerEndpointConfigurationProductionVariantsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerEndpointConfiguration.SagemakerEndpointConfigurationProductionVariantsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference)SetAcceleratorType(val *string) {
	if err := j.validateSetAcceleratorTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorType",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference)SetContainerStartupHealthCheckTimeoutInSeconds(val *float64) {
	if err := j.validateSetContainerStartupHealthCheckTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerStartupHealthCheckTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference)SetEnableSsmAccess(val interface{}) {
	if err := j.validateSetEnableSsmAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSsmAccess",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference)SetInferenceAmiVersion(val *string) {
	if err := j.validateSetInferenceAmiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceAmiVersion",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference)SetInitialInstanceCount(val *float64) {
	if err := j.validateSetInitialInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialInstanceCount",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference)SetInitialVariantWeight(val *float64) {
	if err := j.validateSetInitialVariantWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialVariantWeight",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference)SetModelDataDownloadTimeoutInSeconds(val *float64) {
	if err := j.validateSetModelDataDownloadTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDataDownloadTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference)SetModelName(val *string) {
	if err := j.validateSetModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelName",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference)SetVariantName(val *string) {
	if err := j.validateSetVariantNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variantName",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference)SetVolumeSizeInGb(val *float64) {
	if err := j.validateSetVolumeSizeInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeSizeInGb",
		val,
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) PutCoreDumpConfig(value *SagemakerEndpointConfigurationProductionVariantsCoreDumpConfig) {
	if err := s.validatePutCoreDumpConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCoreDumpConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) PutManagedInstanceScaling(value *SagemakerEndpointConfigurationProductionVariantsManagedInstanceScaling) {
	if err := s.validatePutManagedInstanceScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putManagedInstanceScaling",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) PutRoutingConfig(value interface{}) {
	if err := s.validatePutRoutingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRoutingConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) PutServerlessConfig(value *SagemakerEndpointConfigurationProductionVariantsServerlessConfig) {
	if err := s.validatePutServerlessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putServerlessConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ResetAcceleratorType() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceleratorType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ResetContainerStartupHealthCheckTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetContainerStartupHealthCheckTimeoutInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ResetCoreDumpConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetCoreDumpConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ResetEnableSsmAccess() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableSsmAccess",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ResetInferenceAmiVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetInferenceAmiVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ResetInitialInstanceCount() {
	_jsii_.InvokeVoid(
		s,
		"resetInitialInstanceCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ResetInitialVariantWeight() {
	_jsii_.InvokeVoid(
		s,
		"resetInitialVariantWeight",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ResetManagedInstanceScaling() {
	_jsii_.InvokeVoid(
		s,
		"resetManagedInstanceScaling",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ResetModelDataDownloadTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetModelDataDownloadTimeoutInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ResetRoutingConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetRoutingConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ResetServerlessConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetServerlessConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ResetVariantName() {
	_jsii_.InvokeVoid(
		s,
		"resetVariantName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ResetVolumeSizeInGb() {
	_jsii_.InvokeVoid(
		s,
		"resetVolumeSizeInGb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigurationProductionVariantsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

