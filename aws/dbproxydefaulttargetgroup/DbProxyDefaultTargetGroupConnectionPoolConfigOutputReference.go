// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dbproxydefaulttargetgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/dbproxydefaulttargetgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference interface {
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
	ConnectionBorrowTimeout() *float64
	SetConnectionBorrowTimeout(val *float64)
	ConnectionBorrowTimeoutInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InitQuery() *string
	SetInitQuery(val *string)
	InitQueryInput() *string
	InternalValue() *DbProxyDefaultTargetGroupConnectionPoolConfig
	SetInternalValue(val *DbProxyDefaultTargetGroupConnectionPoolConfig)
	MaxConnectionsPercent() *float64
	SetMaxConnectionsPercent(val *float64)
	MaxConnectionsPercentInput() *float64
	MaxIdleConnectionsPercent() *float64
	SetMaxIdleConnectionsPercent(val *float64)
	MaxIdleConnectionsPercentInput() *float64
	SessionPinningFilters() *[]*string
	SetSessionPinningFilters(val *[]*string)
	SessionPinningFiltersInput() *[]*string
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
	ResetConnectionBorrowTimeout()
	ResetInitQuery()
	ResetMaxConnectionsPercent()
	ResetMaxIdleConnectionsPercent()
	ResetSessionPinningFilters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference
type jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) ConnectionBorrowTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionBorrowTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) ConnectionBorrowTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionBorrowTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) InitQuery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) InitQueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) InternalValue() *DbProxyDefaultTargetGroupConnectionPoolConfig {
	var returns *DbProxyDefaultTargetGroupConnectionPoolConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) MaxConnectionsPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) MaxConnectionsPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) MaxIdleConnectionsPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleConnectionsPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) MaxIdleConnectionsPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleConnectionsPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) SessionPinningFilters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sessionPinningFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) SessionPinningFiltersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sessionPinningFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDbProxyDefaultTargetGroupConnectionPoolConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDbProxyDefaultTargetGroupConnectionPoolConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dbProxyDefaultTargetGroup.DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDbProxyDefaultTargetGroupConnectionPoolConfigOutputReference_Override(d DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dbProxyDefaultTargetGroup.DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference)SetConnectionBorrowTimeout(val *float64) {
	if err := j.validateSetConnectionBorrowTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionBorrowTimeout",
		val,
	)
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference)SetInitQuery(val *string) {
	if err := j.validateSetInitQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initQuery",
		val,
	)
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference)SetInternalValue(val *DbProxyDefaultTargetGroupConnectionPoolConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference)SetMaxConnectionsPercent(val *float64) {
	if err := j.validateSetMaxConnectionsPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConnectionsPercent",
		val,
	)
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference)SetMaxIdleConnectionsPercent(val *float64) {
	if err := j.validateSetMaxIdleConnectionsPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIdleConnectionsPercent",
		val,
	)
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference)SetSessionPinningFilters(val *[]*string) {
	if err := j.validateSetSessionPinningFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionPinningFilters",
		val,
	)
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) ResetConnectionBorrowTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectionBorrowTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) ResetInitQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetInitQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) ResetMaxConnectionsPercent() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxConnectionsPercent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) ResetMaxIdleConnectionsPercent() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxIdleConnectionsPercent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) ResetSessionPinningFilters() {
	_jsii_.InvokeVoid(
		d,
		"resetSessionPinningFilters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DbProxyDefaultTargetGroupConnectionPoolConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

