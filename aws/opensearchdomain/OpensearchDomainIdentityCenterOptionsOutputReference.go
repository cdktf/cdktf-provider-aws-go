// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/opensearchdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpensearchDomainIdentityCenterOptionsOutputReference interface {
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
	EnabledApiAccess() interface{}
	SetEnabledApiAccess(val interface{})
	EnabledApiAccessInput() interface{}
	// Experimental.
	Fqn() *string
	IdentityCenterInstanceArn() *string
	SetIdentityCenterInstanceArn(val *string)
	IdentityCenterInstanceArnInput() *string
	InternalValue() *OpensearchDomainIdentityCenterOptions
	SetInternalValue(val *OpensearchDomainIdentityCenterOptions)
	RolesKey() *string
	SetRolesKey(val *string)
	RolesKeyInput() *string
	SubjectKey() *string
	SetSubjectKey(val *string)
	SubjectKeyInput() *string
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
	ResetEnabledApiAccess()
	ResetIdentityCenterInstanceArn()
	ResetRolesKey()
	ResetSubjectKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OpensearchDomainIdentityCenterOptionsOutputReference
type jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) EnabledApiAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledApiAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) EnabledApiAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledApiAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) IdentityCenterInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityCenterInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) IdentityCenterInstanceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityCenterInstanceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) InternalValue() *OpensearchDomainIdentityCenterOptions {
	var returns *OpensearchDomainIdentityCenterOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) RolesKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rolesKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) RolesKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rolesKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) SubjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) SubjectKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOpensearchDomainIdentityCenterOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OpensearchDomainIdentityCenterOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewOpensearchDomainIdentityCenterOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.opensearchDomain.OpensearchDomainIdentityCenterOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOpensearchDomainIdentityCenterOptionsOutputReference_Override(o OpensearchDomainIdentityCenterOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.opensearchDomain.OpensearchDomainIdentityCenterOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference)SetEnabledApiAccess(val interface{}) {
	if err := j.validateSetEnabledApiAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledApiAccess",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference)SetIdentityCenterInstanceArn(val *string) {
	if err := j.validateSetIdentityCenterInstanceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityCenterInstanceArn",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference)SetInternalValue(val *OpensearchDomainIdentityCenterOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference)SetRolesKey(val *string) {
	if err := j.validateSetRolesKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rolesKey",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference)SetSubjectKey(val *string) {
	if err := j.validateSetSubjectKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectKey",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) ResetEnabledApiAccess() {
	_jsii_.InvokeVoid(
		o,
		"resetEnabledApiAccess",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) ResetIdentityCenterInstanceArn() {
	_jsii_.InvokeVoid(
		o,
		"resetIdentityCenterInstanceArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) ResetRolesKey() {
	_jsii_.InvokeVoid(
		o,
		"resetRolesKey",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) ResetSubjectKey() {
	_jsii_.InvokeVoid(
		o,
		"resetSubjectKey",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainIdentityCenterOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

