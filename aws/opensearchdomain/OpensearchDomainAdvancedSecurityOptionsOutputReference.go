// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/opensearchdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpensearchDomainAdvancedSecurityOptionsOutputReference interface {
	cdktf.ComplexObject
	AnonymousAuthEnabled() interface{}
	SetAnonymousAuthEnabled(val interface{})
	AnonymousAuthEnabledInput() interface{}
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalUserDatabaseEnabled() interface{}
	SetInternalUserDatabaseEnabled(val interface{})
	InternalUserDatabaseEnabledInput() interface{}
	InternalValue() *OpensearchDomainAdvancedSecurityOptions
	SetInternalValue(val *OpensearchDomainAdvancedSecurityOptions)
	MasterUserOptions() OpensearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference
	MasterUserOptionsInput() *OpensearchDomainAdvancedSecurityOptionsMasterUserOptions
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
	PutMasterUserOptions(value *OpensearchDomainAdvancedSecurityOptionsMasterUserOptions)
	ResetAnonymousAuthEnabled()
	ResetInternalUserDatabaseEnabled()
	ResetMasterUserOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OpensearchDomainAdvancedSecurityOptionsOutputReference
type jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) AnonymousAuthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anonymousAuthEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) AnonymousAuthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anonymousAuthEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) InternalUserDatabaseEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalUserDatabaseEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) InternalUserDatabaseEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalUserDatabaseEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) InternalValue() *OpensearchDomainAdvancedSecurityOptions {
	var returns *OpensearchDomainAdvancedSecurityOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) MasterUserOptions() OpensearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference {
	var returns OpensearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference
	_jsii_.Get(
		j,
		"masterUserOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) MasterUserOptionsInput() *OpensearchDomainAdvancedSecurityOptionsMasterUserOptions {
	var returns *OpensearchDomainAdvancedSecurityOptionsMasterUserOptions
	_jsii_.Get(
		j,
		"masterUserOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOpensearchDomainAdvancedSecurityOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OpensearchDomainAdvancedSecurityOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewOpensearchDomainAdvancedSecurityOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.opensearchDomain.OpensearchDomainAdvancedSecurityOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOpensearchDomainAdvancedSecurityOptionsOutputReference_Override(o OpensearchDomainAdvancedSecurityOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.opensearchDomain.OpensearchDomainAdvancedSecurityOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference)SetAnonymousAuthEnabled(val interface{}) {
	if err := j.validateSetAnonymousAuthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anonymousAuthEnabled",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference)SetInternalUserDatabaseEnabled(val interface{}) {
	if err := j.validateSetInternalUserDatabaseEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalUserDatabaseEnabled",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference)SetInternalValue(val *OpensearchDomainAdvancedSecurityOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) PutMasterUserOptions(value *OpensearchDomainAdvancedSecurityOptionsMasterUserOptions) {
	if err := o.validatePutMasterUserOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMasterUserOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) ResetAnonymousAuthEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetAnonymousAuthEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) ResetInternalUserDatabaseEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetInternalUserDatabaseEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) ResetMasterUserOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetMasterUserOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OpensearchDomainAdvancedSecurityOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

