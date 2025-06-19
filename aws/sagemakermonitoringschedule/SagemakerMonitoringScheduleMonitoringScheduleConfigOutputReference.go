// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakermonitoringschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/sagemakermonitoringschedule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference interface {
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
	InternalValue() *SagemakerMonitoringScheduleMonitoringScheduleConfig
	SetInternalValue(val *SagemakerMonitoringScheduleMonitoringScheduleConfig)
	MonitoringJobDefinitionName() *string
	SetMonitoringJobDefinitionName(val *string)
	MonitoringJobDefinitionNameInput() *string
	MonitoringType() *string
	SetMonitoringType(val *string)
	MonitoringTypeInput() *string
	ScheduleConfig() SagemakerMonitoringScheduleMonitoringScheduleConfigScheduleConfigOutputReference
	ScheduleConfigInput() *SagemakerMonitoringScheduleMonitoringScheduleConfigScheduleConfig
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
	PutScheduleConfig(value *SagemakerMonitoringScheduleMonitoringScheduleConfigScheduleConfig)
	ResetScheduleConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference
type jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) InternalValue() *SagemakerMonitoringScheduleMonitoringScheduleConfig {
	var returns *SagemakerMonitoringScheduleMonitoringScheduleConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) MonitoringJobDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringJobDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) MonitoringJobDefinitionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringJobDefinitionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) MonitoringType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) MonitoringTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) ScheduleConfig() SagemakerMonitoringScheduleMonitoringScheduleConfigScheduleConfigOutputReference {
	var returns SagemakerMonitoringScheduleMonitoringScheduleConfigScheduleConfigOutputReference
	_jsii_.Get(
		j,
		"scheduleConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) ScheduleConfigInput() *SagemakerMonitoringScheduleMonitoringScheduleConfigScheduleConfig {
	var returns *SagemakerMonitoringScheduleMonitoringScheduleConfigScheduleConfig
	_jsii_.Get(
		j,
		"scheduleConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerMonitoringScheduleMonitoringScheduleConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerMonitoringSchedule.SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference_Override(s SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerMonitoringSchedule.SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference)SetInternalValue(val *SagemakerMonitoringScheduleMonitoringScheduleConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference)SetMonitoringJobDefinitionName(val *string) {
	if err := j.validateSetMonitoringJobDefinitionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringJobDefinitionName",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference)SetMonitoringType(val *string) {
	if err := j.validateSetMonitoringTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringType",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) PutScheduleConfig(value *SagemakerMonitoringScheduleMonitoringScheduleConfigScheduleConfig) {
	if err := s.validatePutScheduleConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putScheduleConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) ResetScheduleConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetScheduleConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerMonitoringScheduleMonitoringScheduleConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

