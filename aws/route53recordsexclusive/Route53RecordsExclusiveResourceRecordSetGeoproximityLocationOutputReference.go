// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53recordsexclusive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/route53recordsexclusive/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference interface {
	cdktf.ComplexObject
	AwsRegion() *string
	SetAwsRegion(val *string)
	AwsRegionInput() *string
	Bias() *float64
	SetBias(val *float64)
	BiasInput() *float64
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
	Coordinates() Route53RecordsExclusiveResourceRecordSetGeoproximityLocationCoordinatesList
	CoordinatesInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LocalZoneGroup() *string
	SetLocalZoneGroup(val *string)
	LocalZoneGroupInput() *string
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
	PutCoordinates(value interface{})
	ResetAwsRegion()
	ResetBias()
	ResetCoordinates()
	ResetLocalZoneGroup()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference
type jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) AwsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) AwsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) Bias() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) BiasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"biasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) Coordinates() Route53RecordsExclusiveResourceRecordSetGeoproximityLocationCoordinatesList {
	var returns Route53RecordsExclusiveResourceRecordSetGeoproximityLocationCoordinatesList
	_jsii_.Get(
		j,
		"coordinates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) CoordinatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"coordinatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) LocalZoneGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localZoneGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) LocalZoneGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localZoneGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRoute53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference {
	_init_.Initialize()

	if err := validateNewRoute53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.route53RecordsExclusive.Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewRoute53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference_Override(r Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.route53RecordsExclusive.Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference)SetAwsRegion(val *string) {
	if err := j.validateSetAwsRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsRegion",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference)SetBias(val *float64) {
	if err := j.validateSetBiasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bias",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference)SetLocalZoneGroup(val *string) {
	if err := j.validateSetLocalZoneGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localZoneGroup",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) PutCoordinates(value interface{}) {
	if err := r.validatePutCoordinatesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putCoordinates",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) ResetAwsRegion() {
	_jsii_.InvokeVoid(
		r,
		"resetAwsRegion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) ResetBias() {
	_jsii_.InvokeVoid(
		r,
		"resetBias",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) ResetCoordinates() {
	_jsii_.InvokeVoid(
		r,
		"resetCoordinates",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) ResetLocalZoneGroup() {
	_jsii_.InvokeVoid(
		r,
		"resetLocalZoneGroup",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetGeoproximityLocationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

