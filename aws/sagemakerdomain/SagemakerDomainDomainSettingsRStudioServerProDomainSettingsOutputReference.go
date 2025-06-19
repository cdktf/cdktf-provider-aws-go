// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/sagemakerdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference interface {
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
	DefaultResourceSpec() SagemakerDomainDomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecOutputReference
	DefaultResourceSpecInput() *SagemakerDomainDomainSettingsRStudioServerProDomainSettingsDefaultResourceSpec
	DomainExecutionRoleArn() *string
	SetDomainExecutionRoleArn(val *string)
	DomainExecutionRoleArnInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *SagemakerDomainDomainSettingsRStudioServerProDomainSettings
	SetInternalValue(val *SagemakerDomainDomainSettingsRStudioServerProDomainSettings)
	RStudioConnectUrl() *string
	SetRStudioConnectUrl(val *string)
	RStudioConnectUrlInput() *string
	RStudioPackageManagerUrl() *string
	SetRStudioPackageManagerUrl(val *string)
	RStudioPackageManagerUrlInput() *string
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
	PutDefaultResourceSpec(value *SagemakerDomainDomainSettingsRStudioServerProDomainSettingsDefaultResourceSpec)
	ResetDefaultResourceSpec()
	ResetRStudioConnectUrl()
	ResetRStudioPackageManagerUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference
type jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) DefaultResourceSpec() SagemakerDomainDomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecOutputReference {
	var returns SagemakerDomainDomainSettingsRStudioServerProDomainSettingsDefaultResourceSpecOutputReference
	_jsii_.Get(
		j,
		"defaultResourceSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) DefaultResourceSpecInput() *SagemakerDomainDomainSettingsRStudioServerProDomainSettingsDefaultResourceSpec {
	var returns *SagemakerDomainDomainSettingsRStudioServerProDomainSettingsDefaultResourceSpec
	_jsii_.Get(
		j,
		"defaultResourceSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) DomainExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) DomainExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainExecutionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) InternalValue() *SagemakerDomainDomainSettingsRStudioServerProDomainSettings {
	var returns *SagemakerDomainDomainSettingsRStudioServerProDomainSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) RStudioConnectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rStudioConnectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) RStudioConnectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rStudioConnectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) RStudioPackageManagerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rStudioPackageManagerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) RStudioPackageManagerUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rStudioPackageManagerUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerDomain.SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference_Override(s SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerDomain.SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference)SetDomainExecutionRoleArn(val *string) {
	if err := j.validateSetDomainExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainExecutionRoleArn",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference)SetInternalValue(val *SagemakerDomainDomainSettingsRStudioServerProDomainSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference)SetRStudioConnectUrl(val *string) {
	if err := j.validateSetRStudioConnectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rStudioConnectUrl",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference)SetRStudioPackageManagerUrl(val *string) {
	if err := j.validateSetRStudioPackageManagerUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rStudioPackageManagerUrl",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) PutDefaultResourceSpec(value *SagemakerDomainDomainSettingsRStudioServerProDomainSettingsDefaultResourceSpec) {
	if err := s.validatePutDefaultResourceSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDefaultResourceSpec",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) ResetDefaultResourceSpec() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultResourceSpec",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) ResetRStudioConnectUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetRStudioConnectUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) ResetRStudioPackageManagerUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetRStudioPackageManagerUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

