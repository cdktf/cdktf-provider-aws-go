// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksmemcachedlayer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/opsworksmemcachedlayer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference interface {
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
	Downscaling() OpsworksMemcachedLayerLoadBasedAutoScalingDownscalingOutputReference
	DownscalingInput() *OpsworksMemcachedLayerLoadBasedAutoScalingDownscaling
	Enable() interface{}
	SetEnable(val interface{})
	EnableInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *OpsworksMemcachedLayerLoadBasedAutoScaling
	SetInternalValue(val *OpsworksMemcachedLayerLoadBasedAutoScaling)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Upscaling() OpsworksMemcachedLayerLoadBasedAutoScalingUpscalingOutputReference
	UpscalingInput() *OpsworksMemcachedLayerLoadBasedAutoScalingUpscaling
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
	PutDownscaling(value *OpsworksMemcachedLayerLoadBasedAutoScalingDownscaling)
	PutUpscaling(value *OpsworksMemcachedLayerLoadBasedAutoScalingUpscaling)
	ResetDownscaling()
	ResetEnable()
	ResetUpscaling()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference
type jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) Downscaling() OpsworksMemcachedLayerLoadBasedAutoScalingDownscalingOutputReference {
	var returns OpsworksMemcachedLayerLoadBasedAutoScalingDownscalingOutputReference
	_jsii_.Get(
		j,
		"downscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) DownscalingInput() *OpsworksMemcachedLayerLoadBasedAutoScalingDownscaling {
	var returns *OpsworksMemcachedLayerLoadBasedAutoScalingDownscaling
	_jsii_.Get(
		j,
		"downscalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) Enable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) EnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) InternalValue() *OpsworksMemcachedLayerLoadBasedAutoScaling {
	var returns *OpsworksMemcachedLayerLoadBasedAutoScaling
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) Upscaling() OpsworksMemcachedLayerLoadBasedAutoScalingUpscalingOutputReference {
	var returns OpsworksMemcachedLayerLoadBasedAutoScalingUpscalingOutputReference
	_jsii_.Get(
		j,
		"upscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) UpscalingInput() *OpsworksMemcachedLayerLoadBasedAutoScalingUpscaling {
	var returns *OpsworksMemcachedLayerLoadBasedAutoScalingUpscaling
	_jsii_.Get(
		j,
		"upscalingInput",
		&returns,
	)
	return returns
}


func NewOpsworksMemcachedLayerLoadBasedAutoScalingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference {
	_init_.Initialize()

	if err := validateNewOpsworksMemcachedLayerLoadBasedAutoScalingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksMemcachedLayer.OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOpsworksMemcachedLayerLoadBasedAutoScalingOutputReference_Override(o OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksMemcachedLayer.OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference)SetEnable(val interface{}) {
	if err := j.validateSetEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enable",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference)SetInternalValue(val *OpsworksMemcachedLayerLoadBasedAutoScaling) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) PutDownscaling(value *OpsworksMemcachedLayerLoadBasedAutoScalingDownscaling) {
	if err := o.validatePutDownscalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDownscaling",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) PutUpscaling(value *OpsworksMemcachedLayerLoadBasedAutoScalingUpscaling) {
	if err := o.validatePutUpscalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putUpscaling",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) ResetDownscaling() {
	_jsii_.InvokeVoid(
		o,
		"resetDownscaling",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) ResetEnable() {
	_jsii_.InvokeVoid(
		o,
		"resetEnable",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) ResetUpscaling() {
	_jsii_.InvokeVoid(
		o,
		"resetUpscaling",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksMemcachedLayerLoadBasedAutoScalingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

