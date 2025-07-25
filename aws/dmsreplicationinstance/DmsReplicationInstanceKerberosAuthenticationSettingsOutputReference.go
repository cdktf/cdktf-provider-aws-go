// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsreplicationinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/dmsreplicationinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference interface {
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
	InternalValue() *DmsReplicationInstanceKerberosAuthenticationSettings
	SetInternalValue(val *DmsReplicationInstanceKerberosAuthenticationSettings)
	KeyCacheSecretIamArn() *string
	SetKeyCacheSecretIamArn(val *string)
	KeyCacheSecretIamArnInput() *string
	KeyCacheSecretId() *string
	SetKeyCacheSecretId(val *string)
	KeyCacheSecretIdInput() *string
	Krb5FileContents() *string
	SetKrb5FileContents(val *string)
	Krb5FileContentsInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference
type jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) InternalValue() *DmsReplicationInstanceKerberosAuthenticationSettings {
	var returns *DmsReplicationInstanceKerberosAuthenticationSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) KeyCacheSecretIamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyCacheSecretIamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) KeyCacheSecretIamArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyCacheSecretIamArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) KeyCacheSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyCacheSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) KeyCacheSecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyCacheSecretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) Krb5FileContents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"krb5FileContents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) Krb5FileContentsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"krb5FileContentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDmsReplicationInstanceKerberosAuthenticationSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDmsReplicationInstanceKerberosAuthenticationSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dmsReplicationInstance.DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDmsReplicationInstanceKerberosAuthenticationSettingsOutputReference_Override(d DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dmsReplicationInstance.DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference)SetInternalValue(val *DmsReplicationInstanceKerberosAuthenticationSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference)SetKeyCacheSecretIamArn(val *string) {
	if err := j.validateSetKeyCacheSecretIamArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyCacheSecretIamArn",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference)SetKeyCacheSecretId(val *string) {
	if err := j.validateSetKeyCacheSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyCacheSecretId",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference)SetKrb5FileContents(val *string) {
	if err := j.validateSetKrb5FileContentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"krb5FileContents",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationInstanceKerberosAuthenticationSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

