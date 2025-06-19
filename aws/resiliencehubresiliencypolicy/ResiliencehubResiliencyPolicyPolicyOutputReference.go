// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resiliencehubresiliencypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/resiliencehubresiliencypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ResiliencehubResiliencyPolicyPolicyOutputReference interface {
	cdktf.ComplexObject
	Az() ResiliencehubResiliencyPolicyPolicyAzList
	AzInput() interface{}
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
	Hardware() ResiliencehubResiliencyPolicyPolicyHardwareList
	HardwareInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Region() ResiliencehubResiliencyPolicyPolicyRegionList
	RegionInput() interface{}
	SoftwareAttribute() ResiliencehubResiliencyPolicyPolicySoftwareList
	SoftwareAttributeInput() interface{}
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
	PutAz(value interface{})
	PutHardware(value interface{})
	PutRegion(value interface{})
	PutSoftwareAttribute(value interface{})
	ResetAz()
	ResetHardware()
	ResetRegion()
	ResetSoftwareAttribute()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ResiliencehubResiliencyPolicyPolicyOutputReference
type jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) Az() ResiliencehubResiliencyPolicyPolicyAzList {
	var returns ResiliencehubResiliencyPolicyPolicyAzList
	_jsii_.Get(
		j,
		"az",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) AzInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) Hardware() ResiliencehubResiliencyPolicyPolicyHardwareList {
	var returns ResiliencehubResiliencyPolicyPolicyHardwareList
	_jsii_.Get(
		j,
		"hardware",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) HardwareInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hardwareInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) Region() ResiliencehubResiliencyPolicyPolicyRegionList {
	var returns ResiliencehubResiliencyPolicyPolicyRegionList
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) RegionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) SoftwareAttribute() ResiliencehubResiliencyPolicyPolicySoftwareList {
	var returns ResiliencehubResiliencyPolicyPolicySoftwareList
	_jsii_.Get(
		j,
		"softwareAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) SoftwareAttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"softwareAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewResiliencehubResiliencyPolicyPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ResiliencehubResiliencyPolicyPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewResiliencehubResiliencyPolicyPolicyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.resiliencehubResiliencyPolicy.ResiliencehubResiliencyPolicyPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewResiliencehubResiliencyPolicyPolicyOutputReference_Override(r ResiliencehubResiliencyPolicyPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.resiliencehubResiliencyPolicy.ResiliencehubResiliencyPolicyPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) PutAz(value interface{}) {
	if err := r.validatePutAzParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAz",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) PutHardware(value interface{}) {
	if err := r.validatePutHardwareParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putHardware",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) PutRegion(value interface{}) {
	if err := r.validatePutRegionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putRegion",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) PutSoftwareAttribute(value interface{}) {
	if err := r.validatePutSoftwareAttributeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSoftwareAttribute",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) ResetAz() {
	_jsii_.InvokeVoid(
		r,
		"resetAz",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) ResetHardware() {
	_jsii_.InvokeVoid(
		r,
		"resetHardware",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		r,
		"resetRegion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) ResetSoftwareAttribute() {
	_jsii_.InvokeVoid(
		r,
		"resetSoftwareAttribute",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

