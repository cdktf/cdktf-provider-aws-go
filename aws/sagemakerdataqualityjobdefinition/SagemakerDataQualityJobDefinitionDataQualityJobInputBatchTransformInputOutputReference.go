// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdataqualityjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/sagemakerdataqualityjobdefinition/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference interface {
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
	DataCapturedDestinationS3Uri() *string
	SetDataCapturedDestinationS3Uri(val *string)
	DataCapturedDestinationS3UriInput() *string
	DatasetFormat() SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputDatasetFormatOutputReference
	DatasetFormatInput() *SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputDatasetFormat
	// Experimental.
	Fqn() *string
	InternalValue() *SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInput
	SetInternalValue(val *SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInput)
	LocalPath() *string
	SetLocalPath(val *string)
	LocalPathInput() *string
	S3DataDistributionType() *string
	SetS3DataDistributionType(val *string)
	S3DataDistributionTypeInput() *string
	S3InputMode() *string
	SetS3InputMode(val *string)
	S3InputModeInput() *string
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
	PutDatasetFormat(value *SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputDatasetFormat)
	ResetLocalPath()
	ResetS3DataDistributionType()
	ResetS3InputMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference
type jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) DataCapturedDestinationS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCapturedDestinationS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) DataCapturedDestinationS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCapturedDestinationS3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) DatasetFormat() SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputDatasetFormatOutputReference {
	var returns SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputDatasetFormatOutputReference
	_jsii_.Get(
		j,
		"datasetFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) DatasetFormatInput() *SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputDatasetFormat {
	var returns *SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputDatasetFormat
	_jsii_.Get(
		j,
		"datasetFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) InternalValue() *SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInput {
	var returns *SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) LocalPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) LocalPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) S3DataDistributionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataDistributionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) S3DataDistributionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataDistributionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) S3InputMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3InputMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) S3InputModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3InputModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerDataQualityJobDefinition.SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference_Override(s SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerDataQualityJobDefinition.SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference)SetDataCapturedDestinationS3Uri(val *string) {
	if err := j.validateSetDataCapturedDestinationS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataCapturedDestinationS3Uri",
		val,
	)
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference)SetInternalValue(val *SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference)SetLocalPath(val *string) {
	if err := j.validateSetLocalPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localPath",
		val,
	)
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference)SetS3DataDistributionType(val *string) {
	if err := j.validateSetS3DataDistributionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3DataDistributionType",
		val,
	)
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference)SetS3InputMode(val *string) {
	if err := j.validateSetS3InputModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3InputMode",
		val,
	)
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) PutDatasetFormat(value *SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputDatasetFormat) {
	if err := s.validatePutDatasetFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDatasetFormat",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) ResetLocalPath() {
	_jsii_.InvokeVoid(
		s,
		"resetLocalPath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) ResetS3DataDistributionType() {
	_jsii_.InvokeVoid(
		s,
		"resetS3DataDistributionType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) ResetS3InputMode() {
	_jsii_.InvokeVoid(
		s,
		"resetS3InputMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

