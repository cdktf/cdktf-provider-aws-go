// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resiliencehubresiliencypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/resiliencehubresiliencypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ResiliencehubResiliencyPolicyPolicyRegionOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Rpo() *string
	SetRpo(val *string)
	RpoInput() *string
	Rto() *string
	SetRto(val *string)
	RtoInput() *string
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
	ResetRpo()
	ResetRto()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ResiliencehubResiliencyPolicyPolicyRegionOutputReference
type jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) Rpo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rpo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) RpoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rpoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) Rto() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rto",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) RtoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rtoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewResiliencehubResiliencyPolicyPolicyRegionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ResiliencehubResiliencyPolicyPolicyRegionOutputReference {
	_init_.Initialize()

	if err := validateNewResiliencehubResiliencyPolicyPolicyRegionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.resiliencehubResiliencyPolicy.ResiliencehubResiliencyPolicyPolicyRegionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewResiliencehubResiliencyPolicyPolicyRegionOutputReference_Override(r ResiliencehubResiliencyPolicyPolicyRegionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.resiliencehubResiliencyPolicy.ResiliencehubResiliencyPolicyPolicyRegionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference)SetRpo(val *string) {
	if err := j.validateSetRpoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rpo",
		val,
	)
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference)SetRto(val *string) {
	if err := j.validateSetRtoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rto",
		val,
	)
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) ResetRpo() {
	_jsii_.InvokeVoid(
		r,
		"resetRpo",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) ResetRto() {
	_jsii_.InvokeVoid(
		r,
		"resetRto",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_ResiliencehubResiliencyPolicyPolicyRegionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

