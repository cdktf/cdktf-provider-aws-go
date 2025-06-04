// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/sagemakerdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference interface {
	cdktf.ComplexObject
	AssumableRoleArns() *[]*string
	SetAssumableRoleArns(val *[]*string)
	AssumableRoleArnsInput() *[]*string
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
	ExecutionRoleArns() *[]*string
	SetExecutionRoleArns(val *[]*string)
	ExecutionRoleArnsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettings
	SetInternalValue(val *SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettings)
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
	ResetAssumableRoleArns()
	ResetExecutionRoleArns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference
type jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) AssumableRoleArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assumableRoleArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) AssumableRoleArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assumableRoleArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) ExecutionRoleArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"executionRoleArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) ExecutionRoleArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"executionRoleArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) InternalValue() *SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettings {
	var returns *SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference_Override(s SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference)SetAssumableRoleArns(val *[]*string) {
	if err := j.validateSetAssumableRoleArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assumableRoleArns",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference)SetExecutionRoleArns(val *[]*string) {
	if err := j.validateSetExecutionRoleArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArns",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference)SetInternalValue(val *SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) ResetAssumableRoleArns() {
	_jsii_.InvokeVoid(
		s,
		"resetAssumableRoleArns",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) ResetExecutionRoleArns() {
	_jsii_.InvokeVoid(
		s,
		"resetExecutionRoleArns",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsEmrSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

