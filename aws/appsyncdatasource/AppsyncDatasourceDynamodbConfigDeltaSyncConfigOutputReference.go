// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/appsyncdatasource/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference interface {
	cdktf.ComplexObject
	BaseTableTtl() *float64
	SetBaseTableTtl(val *float64)
	BaseTableTtlInput() *float64
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
	DeltaSyncTableName() *string
	SetDeltaSyncTableName(val *string)
	DeltaSyncTableNameInput() *string
	DeltaSyncTableTtl() *float64
	SetDeltaSyncTableTtl(val *float64)
	DeltaSyncTableTtlInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *AppsyncDatasourceDynamodbConfigDeltaSyncConfig
	SetInternalValue(val *AppsyncDatasourceDynamodbConfigDeltaSyncConfig)
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
	ResetBaseTableTtl()
	ResetDeltaSyncTableTtl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference
type jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) BaseTableTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseTableTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) BaseTableTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseTableTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) DeltaSyncTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaSyncTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) DeltaSyncTableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaSyncTableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) DeltaSyncTableTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deltaSyncTableTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) DeltaSyncTableTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deltaSyncTableTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) InternalValue() *AppsyncDatasourceDynamodbConfigDeltaSyncConfig {
	var returns *AppsyncDatasourceDynamodbConfigDeltaSyncConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appsyncDatasource.AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference_Override(a AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appsyncDatasource.AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference)SetBaseTableTtl(val *float64) {
	if err := j.validateSetBaseTableTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseTableTtl",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference)SetDeltaSyncTableName(val *string) {
	if err := j.validateSetDeltaSyncTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deltaSyncTableName",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference)SetDeltaSyncTableTtl(val *float64) {
	if err := j.validateSetDeltaSyncTableTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deltaSyncTableTtl",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference)SetInternalValue(val *AppsyncDatasourceDynamodbConfigDeltaSyncConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) ResetBaseTableTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetBaseTableTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) ResetDeltaSyncTableTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetDeltaSyncTableTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

