// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/mskcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MskClusterOpenMonitoringPrometheusOutputReference interface {
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
	InternalValue() *MskClusterOpenMonitoringPrometheus
	SetInternalValue(val *MskClusterOpenMonitoringPrometheus)
	JmxExporter() MskClusterOpenMonitoringPrometheusJmxExporterOutputReference
	JmxExporterInput() *MskClusterOpenMonitoringPrometheusJmxExporter
	NodeExporter() MskClusterOpenMonitoringPrometheusNodeExporterOutputReference
	NodeExporterInput() *MskClusterOpenMonitoringPrometheusNodeExporter
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
	PutJmxExporter(value *MskClusterOpenMonitoringPrometheusJmxExporter)
	PutNodeExporter(value *MskClusterOpenMonitoringPrometheusNodeExporter)
	ResetJmxExporter()
	ResetNodeExporter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MskClusterOpenMonitoringPrometheusOutputReference
type jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) InternalValue() *MskClusterOpenMonitoringPrometheus {
	var returns *MskClusterOpenMonitoringPrometheus
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) JmxExporter() MskClusterOpenMonitoringPrometheusJmxExporterOutputReference {
	var returns MskClusterOpenMonitoringPrometheusJmxExporterOutputReference
	_jsii_.Get(
		j,
		"jmxExporter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) JmxExporterInput() *MskClusterOpenMonitoringPrometheusJmxExporter {
	var returns *MskClusterOpenMonitoringPrometheusJmxExporter
	_jsii_.Get(
		j,
		"jmxExporterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) NodeExporter() MskClusterOpenMonitoringPrometheusNodeExporterOutputReference {
	var returns MskClusterOpenMonitoringPrometheusNodeExporterOutputReference
	_jsii_.Get(
		j,
		"nodeExporter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) NodeExporterInput() *MskClusterOpenMonitoringPrometheusNodeExporter {
	var returns *MskClusterOpenMonitoringPrometheusNodeExporter
	_jsii_.Get(
		j,
		"nodeExporterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMskClusterOpenMonitoringPrometheusOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MskClusterOpenMonitoringPrometheusOutputReference {
	_init_.Initialize()

	if err := validateNewMskClusterOpenMonitoringPrometheusOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.mskCluster.MskClusterOpenMonitoringPrometheusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMskClusterOpenMonitoringPrometheusOutputReference_Override(m MskClusterOpenMonitoringPrometheusOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.mskCluster.MskClusterOpenMonitoringPrometheusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference)SetInternalValue(val *MskClusterOpenMonitoringPrometheus) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) PutJmxExporter(value *MskClusterOpenMonitoringPrometheusJmxExporter) {
	if err := m.validatePutJmxExporterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putJmxExporter",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) PutNodeExporter(value *MskClusterOpenMonitoringPrometheusNodeExporter) {
	if err := m.validatePutNodeExporterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putNodeExporter",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) ResetJmxExporter() {
	_jsii_.InvokeVoid(
		m,
		"resetJmxExporter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) ResetNodeExporter() {
	_jsii_.InvokeVoid(
		m,
		"resetNodeExporter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskClusterOpenMonitoringPrometheusOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

