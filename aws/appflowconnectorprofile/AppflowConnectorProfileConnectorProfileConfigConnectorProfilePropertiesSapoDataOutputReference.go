// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/appflowconnectorprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference interface {
	cdktf.ComplexObject
	ApplicationHostUrl() *string
	SetApplicationHostUrl(val *string)
	ApplicationHostUrlInput() *string
	ApplicationServicePath() *string
	SetApplicationServicePath(val *string)
	ApplicationServicePathInput() *string
	ClientNumber() *string
	SetClientNumber(val *string)
	ClientNumberInput() *string
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
	InternalValue() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoData
	SetInternalValue(val *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoData)
	LogonLanguage() *string
	SetLogonLanguage(val *string)
	LogonLanguageInput() *string
	OauthProperties() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOauthPropertiesOutputReference
	OauthPropertiesInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOauthProperties
	PortNumber() *float64
	SetPortNumber(val *float64)
	PortNumberInput() *float64
	PrivateLinkServiceName() *string
	SetPrivateLinkServiceName(val *string)
	PrivateLinkServiceNameInput() *string
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
	PutOauthProperties(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOauthProperties)
	ResetLogonLanguage()
	ResetOauthProperties()
	ResetPrivateLinkServiceName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference
type jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) ApplicationHostUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationHostUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) ApplicationHostUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationHostUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) ApplicationServicePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationServicePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) ApplicationServicePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationServicePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) ClientNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) ClientNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) InternalValue() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoData {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoData
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) LogonLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logonLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) LogonLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logonLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) OauthProperties() AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOauthPropertiesOutputReference {
	var returns AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOauthPropertiesOutputReference
	_jsii_.Get(
		j,
		"oauthProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) OauthPropertiesInput() *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOauthProperties {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOauthProperties
	_jsii_.Get(
		j,
		"oauthPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) PortNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) PortNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) PrivateLinkServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateLinkServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) PrivateLinkServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateLinkServiceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference {
	_init_.Initialize()

	if err := validateNewAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appflowConnectorProfile.AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference_Override(a AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appflowConnectorProfile.AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference)SetApplicationHostUrl(val *string) {
	if err := j.validateSetApplicationHostUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationHostUrl",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference)SetApplicationServicePath(val *string) {
	if err := j.validateSetApplicationServicePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationServicePath",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference)SetClientNumber(val *string) {
	if err := j.validateSetClientNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientNumber",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference)SetInternalValue(val *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoData) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference)SetLogonLanguage(val *string) {
	if err := j.validateSetLogonLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logonLanguage",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference)SetPortNumber(val *float64) {
	if err := j.validateSetPortNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portNumber",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference)SetPrivateLinkServiceName(val *string) {
	if err := j.validateSetPrivateLinkServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateLinkServiceName",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) PutOauthProperties(value *AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOauthProperties) {
	if err := a.validatePutOauthPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOauthProperties",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) ResetLogonLanguage() {
	_jsii_.InvokeVoid(
		a,
		"resetLogonLanguage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) ResetOauthProperties() {
	_jsii_.InvokeVoid(
		a,
		"resetOauthProperties",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) ResetPrivateLinkServiceName() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivateLinkServiceName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

