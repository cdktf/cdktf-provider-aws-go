package opsworkscustomlayer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/opsworkscustomlayer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpsworksCustomLayerLoadBasedAutoScalingOutputReference interface {
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
	Downscaling() OpsworksCustomLayerLoadBasedAutoScalingDownscalingOutputReference
	DownscalingInput() *OpsworksCustomLayerLoadBasedAutoScalingDownscaling
	Enable() interface{}
	SetEnable(val interface{})
	EnableInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *OpsworksCustomLayerLoadBasedAutoScaling
	SetInternalValue(val *OpsworksCustomLayerLoadBasedAutoScaling)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Upscaling() OpsworksCustomLayerLoadBasedAutoScalingUpscalingOutputReference
	UpscalingInput() *OpsworksCustomLayerLoadBasedAutoScalingUpscaling
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
	PutDownscaling(value *OpsworksCustomLayerLoadBasedAutoScalingDownscaling)
	PutUpscaling(value *OpsworksCustomLayerLoadBasedAutoScalingUpscaling)
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

// The jsii proxy struct for OpsworksCustomLayerLoadBasedAutoScalingOutputReference
type jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) Downscaling() OpsworksCustomLayerLoadBasedAutoScalingDownscalingOutputReference {
	var returns OpsworksCustomLayerLoadBasedAutoScalingDownscalingOutputReference
	_jsii_.Get(
		j,
		"downscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) DownscalingInput() *OpsworksCustomLayerLoadBasedAutoScalingDownscaling {
	var returns *OpsworksCustomLayerLoadBasedAutoScalingDownscaling
	_jsii_.Get(
		j,
		"downscalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) Enable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) EnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) InternalValue() *OpsworksCustomLayerLoadBasedAutoScaling {
	var returns *OpsworksCustomLayerLoadBasedAutoScaling
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) Upscaling() OpsworksCustomLayerLoadBasedAutoScalingUpscalingOutputReference {
	var returns OpsworksCustomLayerLoadBasedAutoScalingUpscalingOutputReference
	_jsii_.Get(
		j,
		"upscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) UpscalingInput() *OpsworksCustomLayerLoadBasedAutoScalingUpscaling {
	var returns *OpsworksCustomLayerLoadBasedAutoScalingUpscaling
	_jsii_.Get(
		j,
		"upscalingInput",
		&returns,
	)
	return returns
}


func NewOpsworksCustomLayerLoadBasedAutoScalingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OpsworksCustomLayerLoadBasedAutoScalingOutputReference {
	_init_.Initialize()

	if err := validateNewOpsworksCustomLayerLoadBasedAutoScalingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksCustomLayer.OpsworksCustomLayerLoadBasedAutoScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOpsworksCustomLayerLoadBasedAutoScalingOutputReference_Override(o OpsworksCustomLayerLoadBasedAutoScalingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksCustomLayer.OpsworksCustomLayerLoadBasedAutoScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference)SetEnable(val interface{}) {
	if err := j.validateSetEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enable",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference)SetInternalValue(val *OpsworksCustomLayerLoadBasedAutoScaling) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) PutDownscaling(value *OpsworksCustomLayerLoadBasedAutoScalingDownscaling) {
	if err := o.validatePutDownscalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDownscaling",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) PutUpscaling(value *OpsworksCustomLayerLoadBasedAutoScalingUpscaling) {
	if err := o.validatePutUpscalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putUpscaling",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) ResetDownscaling() {
	_jsii_.InvokeVoid(
		o,
		"resetDownscaling",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) ResetEnable() {
	_jsii_.InvokeVoid(
		o,
		"resetEnable",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) ResetUpscaling() {
	_jsii_.InvokeVoid(
		o,
		"resetUpscaling",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OpsworksCustomLayerLoadBasedAutoScalingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

