// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticsearchdomainsamloptions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/elasticsearchdomainsamloptions/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElasticsearchDomainSamlOptionsSamlOptionsOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	Idp() ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference
	IdpInput() *ElasticsearchDomainSamlOptionsSamlOptionsIdp
	InternalValue() *ElasticsearchDomainSamlOptionsSamlOptions
	SetInternalValue(val *ElasticsearchDomainSamlOptionsSamlOptions)
	MasterBackendRole() *string
	SetMasterBackendRole(val *string)
	MasterBackendRoleInput() *string
	MasterUserName() *string
	SetMasterUserName(val *string)
	MasterUserNameInput() *string
	RolesKey() *string
	SetRolesKey(val *string)
	RolesKeyInput() *string
	SessionTimeoutMinutes() *float64
	SetSessionTimeoutMinutes(val *float64)
	SessionTimeoutMinutesInput() *float64
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
	PutIdp(value *ElasticsearchDomainSamlOptionsSamlOptionsIdp)
	ResetEnabled()
	ResetIdp()
	ResetMasterBackendRole()
	ResetMasterUserName()
	ResetRolesKey()
	ResetSessionTimeoutMinutes()
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

// The jsii proxy struct for ElasticsearchDomainSamlOptionsSamlOptionsOutputReference
type jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) Idp() ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference {
	var returns ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference
	_jsii_.Get(
		j,
		"idp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) IdpInput() *ElasticsearchDomainSamlOptionsSamlOptionsIdp {
	var returns *ElasticsearchDomainSamlOptionsSamlOptionsIdp
	_jsii_.Get(
		j,
		"idpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) InternalValue() *ElasticsearchDomainSamlOptionsSamlOptions {
	var returns *ElasticsearchDomainSamlOptionsSamlOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) MasterBackendRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterBackendRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) MasterBackendRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterBackendRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) MasterUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) MasterUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) RolesKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rolesKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) RolesKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rolesKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) SessionTimeoutMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) SessionTimeoutMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) SubjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) SubjectKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElasticsearchDomainSamlOptionsSamlOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElasticsearchDomainSamlOptionsSamlOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewElasticsearchDomainSamlOptionsSamlOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.elasticsearchDomainSamlOptions.ElasticsearchDomainSamlOptionsSamlOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElasticsearchDomainSamlOptionsSamlOptionsOutputReference_Override(e ElasticsearchDomainSamlOptionsSamlOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.elasticsearchDomainSamlOptions.ElasticsearchDomainSamlOptionsSamlOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference)SetInternalValue(val *ElasticsearchDomainSamlOptionsSamlOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference)SetMasterBackendRole(val *string) {
	if err := j.validateSetMasterBackendRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterBackendRole",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference)SetMasterUserName(val *string) {
	if err := j.validateSetMasterUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserName",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference)SetRolesKey(val *string) {
	if err := j.validateSetRolesKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rolesKey",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference)SetSessionTimeoutMinutes(val *float64) {
	if err := j.validateSetSessionTimeoutMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionTimeoutMinutes",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference)SetSubjectKey(val *string) {
	if err := j.validateSetSubjectKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectKey",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) PutIdp(value *ElasticsearchDomainSamlOptionsSamlOptionsIdp) {
	if err := e.validatePutIdpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIdp",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) ResetIdp() {
	_jsii_.InvokeVoid(
		e,
		"resetIdp",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) ResetMasterBackendRole() {
	_jsii_.InvokeVoid(
		e,
		"resetMasterBackendRole",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) ResetMasterUserName() {
	_jsii_.InvokeVoid(
		e,
		"resetMasterUserName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) ResetRolesKey() {
	_jsii_.InvokeVoid(
		e,
		"resetRolesKey",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) ResetSessionTimeoutMinutes() {
	_jsii_.InvokeVoid(
		e,
		"resetSessionTimeoutMinutes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) ResetSubjectKey() {
	_jsii_.InvokeVoid(
		e,
		"resetSubjectKey",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

