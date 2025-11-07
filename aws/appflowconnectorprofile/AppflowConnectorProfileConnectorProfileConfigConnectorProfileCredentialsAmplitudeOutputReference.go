// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/appflowconnectorprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference interface {
	cdktf.ComplexObject
	ApiKey() *string
	SetApiKey(val *string)
	ApiKeyInput() *string
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
	InternalValue() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitude
	SetInternalValue(val *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitude)
	SecretKey() *string
	SetSecretKey(val *string)
	SecretKeyInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference
type jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) ApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) ApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) InternalValue() *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitude {
	var returns *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitude
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) SecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) SecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference {
	_init_.Initialize()

	if err := validateNewAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appflowConnectorProfile.AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference_Override(a AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appflowConnectorProfile.AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference)SetApiKey(val *string) {
	if err := j.validateSetApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiKey",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference)SetInternalValue(val *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitude) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference)SetSecretKey(val *string) {
	if err := j.validateSetSecretKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretKey",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitudeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

