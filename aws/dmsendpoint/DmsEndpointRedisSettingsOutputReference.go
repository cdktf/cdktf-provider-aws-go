// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/dmsendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsEndpointRedisSettingsOutputReference interface {
	cdktf.ComplexObject
	AuthPassword() *string
	SetAuthPassword(val *string)
	AuthPasswordInput() *string
	AuthType() *string
	SetAuthType(val *string)
	AuthTypeInput() *string
	AuthUserName() *string
	SetAuthUserName(val *string)
	AuthUserNameInput() *string
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
	InternalValue() *DmsEndpointRedisSettings
	SetInternalValue(val *DmsEndpointRedisSettings)
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	ServerName() *string
	SetServerName(val *string)
	ServerNameInput() *string
	SslCaCertificateArn() *string
	SetSslCaCertificateArn(val *string)
	SslCaCertificateArnInput() *string
	SslSecurityProtocol() *string
	SetSslSecurityProtocol(val *string)
	SslSecurityProtocolInput() *string
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
	ResetAuthPassword()
	ResetAuthUserName()
	ResetSslCaCertificateArn()
	ResetSslSecurityProtocol()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsEndpointRedisSettingsOutputReference
type jsiiProxy_DmsEndpointRedisSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) AuthPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) AuthPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) AuthUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) AuthUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) InternalValue() *DmsEndpointRedisSettings {
	var returns *DmsEndpointRedisSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) ServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) ServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) SslCaCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) SslCaCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) SslSecurityProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslSecurityProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) SslSecurityProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslSecurityProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDmsEndpointRedisSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DmsEndpointRedisSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDmsEndpointRedisSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsEndpointRedisSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dmsEndpoint.DmsEndpointRedisSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDmsEndpointRedisSettingsOutputReference_Override(d DmsEndpointRedisSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dmsEndpoint.DmsEndpointRedisSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference)SetAuthPassword(val *string) {
	if err := j.validateSetAuthPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authPassword",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference)SetAuthType(val *string) {
	if err := j.validateSetAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference)SetAuthUserName(val *string) {
	if err := j.validateSetAuthUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authUserName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference)SetInternalValue(val *DmsEndpointRedisSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference)SetServerName(val *string) {
	if err := j.validateSetServerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference)SetSslCaCertificateArn(val *string) {
	if err := j.validateSetSslCaCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCaCertificateArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference)SetSslSecurityProtocol(val *string) {
	if err := j.validateSetSslSecurityProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslSecurityProtocol",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointRedisSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DmsEndpointRedisSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointRedisSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointRedisSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointRedisSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointRedisSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointRedisSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointRedisSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointRedisSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointRedisSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointRedisSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointRedisSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointRedisSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointRedisSettingsOutputReference) ResetAuthPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedisSettingsOutputReference) ResetAuthUserName() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthUserName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedisSettingsOutputReference) ResetSslCaCertificateArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSslCaCertificateArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedisSettingsOutputReference) ResetSslSecurityProtocol() {
	_jsii_.InvokeVoid(
		d,
		"resetSslSecurityProtocol",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointRedisSettingsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointRedisSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

