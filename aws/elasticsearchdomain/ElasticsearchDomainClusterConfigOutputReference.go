// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticsearchdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/elasticsearchdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElasticsearchDomainClusterConfigOutputReference interface {
	cdktf.ComplexObject
	ColdStorageOptions() ElasticsearchDomainClusterConfigColdStorageOptionsOutputReference
	ColdStorageOptionsInput() *ElasticsearchDomainClusterConfigColdStorageOptions
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
	DedicatedMasterCount() *float64
	SetDedicatedMasterCount(val *float64)
	DedicatedMasterCountInput() *float64
	DedicatedMasterEnabled() interface{}
	SetDedicatedMasterEnabled(val interface{})
	DedicatedMasterEnabledInput() interface{}
	DedicatedMasterType() *string
	SetDedicatedMasterType(val *string)
	DedicatedMasterTypeInput() *string
	// Experimental.
	Fqn() *string
	InstanceCount() *float64
	SetInstanceCount(val *float64)
	InstanceCountInput() *float64
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() *ElasticsearchDomainClusterConfig
	SetInternalValue(val *ElasticsearchDomainClusterConfig)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WarmCount() *float64
	SetWarmCount(val *float64)
	WarmCountInput() *float64
	WarmEnabled() interface{}
	SetWarmEnabled(val interface{})
	WarmEnabledInput() interface{}
	WarmType() *string
	SetWarmType(val *string)
	WarmTypeInput() *string
	ZoneAwarenessConfig() ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference
	ZoneAwarenessConfigInput() *ElasticsearchDomainClusterConfigZoneAwarenessConfig
	ZoneAwarenessEnabled() interface{}
	SetZoneAwarenessEnabled(val interface{})
	ZoneAwarenessEnabledInput() interface{}
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
	PutColdStorageOptions(value *ElasticsearchDomainClusterConfigColdStorageOptions)
	PutZoneAwarenessConfig(value *ElasticsearchDomainClusterConfigZoneAwarenessConfig)
	ResetColdStorageOptions()
	ResetDedicatedMasterCount()
	ResetDedicatedMasterEnabled()
	ResetDedicatedMasterType()
	ResetInstanceCount()
	ResetInstanceType()
	ResetWarmCount()
	ResetWarmEnabled()
	ResetWarmType()
	ResetZoneAwarenessConfig()
	ResetZoneAwarenessEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElasticsearchDomainClusterConfigOutputReference
type jsiiProxy_ElasticsearchDomainClusterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ColdStorageOptions() ElasticsearchDomainClusterConfigColdStorageOptionsOutputReference {
	var returns ElasticsearchDomainClusterConfigColdStorageOptionsOutputReference
	_jsii_.Get(
		j,
		"coldStorageOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ColdStorageOptionsInput() *ElasticsearchDomainClusterConfigColdStorageOptions {
	var returns *ElasticsearchDomainClusterConfigColdStorageOptions
	_jsii_.Get(
		j,
		"coldStorageOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) DedicatedMasterCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dedicatedMasterCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) DedicatedMasterCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dedicatedMasterCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) DedicatedMasterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedMasterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) DedicatedMasterEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedMasterEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) DedicatedMasterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedMasterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) DedicatedMasterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedMasterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) InternalValue() *ElasticsearchDomainClusterConfig {
	var returns *ElasticsearchDomainClusterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) WarmCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warmCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) WarmCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warmCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) WarmEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"warmEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) WarmEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"warmEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) WarmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) WarmTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warmTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ZoneAwarenessConfig() ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference {
	var returns ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference
	_jsii_.Get(
		j,
		"zoneAwarenessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ZoneAwarenessConfigInput() *ElasticsearchDomainClusterConfigZoneAwarenessConfig {
	var returns *ElasticsearchDomainClusterConfigZoneAwarenessConfig
	_jsii_.Get(
		j,
		"zoneAwarenessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ZoneAwarenessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneAwarenessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ZoneAwarenessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneAwarenessEnabledInput",
		&returns,
	)
	return returns
}


func NewElasticsearchDomainClusterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElasticsearchDomainClusterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewElasticsearchDomainClusterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticsearchDomainClusterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElasticsearchDomainClusterConfigOutputReference_Override(e ElasticsearchDomainClusterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.elasticsearchDomain.ElasticsearchDomainClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference)SetDedicatedMasterCount(val *float64) {
	if err := j.validateSetDedicatedMasterCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedMasterCount",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference)SetDedicatedMasterEnabled(val interface{}) {
	if err := j.validateSetDedicatedMasterEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedMasterEnabled",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference)SetDedicatedMasterType(val *string) {
	if err := j.validateSetDedicatedMasterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedMasterType",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference)SetInstanceCount(val *float64) {
	if err := j.validateSetInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference)SetInternalValue(val *ElasticsearchDomainClusterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference)SetWarmCount(val *float64) {
	if err := j.validateSetWarmCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warmCount",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference)SetWarmEnabled(val interface{}) {
	if err := j.validateSetWarmEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warmEnabled",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference)SetWarmType(val *string) {
	if err := j.validateSetWarmTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warmType",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference)SetZoneAwarenessEnabled(val interface{}) {
	if err := j.validateSetZoneAwarenessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneAwarenessEnabled",
		val,
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) PutColdStorageOptions(value *ElasticsearchDomainClusterConfigColdStorageOptions) {
	if err := e.validatePutColdStorageOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putColdStorageOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) PutZoneAwarenessConfig(value *ElasticsearchDomainClusterConfigZoneAwarenessConfig) {
	if err := e.validatePutZoneAwarenessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putZoneAwarenessConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetColdStorageOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetColdStorageOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetDedicatedMasterCount() {
	_jsii_.InvokeVoid(
		e,
		"resetDedicatedMasterCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetDedicatedMasterEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetDedicatedMasterEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetDedicatedMasterType() {
	_jsii_.InvokeVoid(
		e,
		"resetDedicatedMasterType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetInstanceCount() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetWarmCount() {
	_jsii_.InvokeVoid(
		e,
		"resetWarmCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetWarmEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetWarmEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetWarmType() {
	_jsii_.InvokeVoid(
		e,
		"resetWarmType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetZoneAwarenessConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetZoneAwarenessConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetZoneAwarenessEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetZoneAwarenessEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

