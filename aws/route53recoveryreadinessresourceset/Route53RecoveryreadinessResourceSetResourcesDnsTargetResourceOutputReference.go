// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53recoveryreadinessresourceset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/route53recoveryreadinessresourceset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference interface {
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
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	// Experimental.
	Fqn() *string
	HostedZoneArn() *string
	SetHostedZoneArn(val *string)
	HostedZoneArnInput() *string
	InternalValue() *Route53RecoveryreadinessResourceSetResourcesDnsTargetResource
	SetInternalValue(val *Route53RecoveryreadinessResourceSetResourcesDnsTargetResource)
	RecordSetId() *string
	SetRecordSetId(val *string)
	RecordSetIdInput() *string
	RecordType() *string
	SetRecordType(val *string)
	RecordTypeInput() *string
	TargetResource() Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference
	TargetResourceInput() *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResource
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
	PutTargetResource(value *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResource)
	ResetHostedZoneArn()
	ResetRecordSetId()
	ResetRecordType()
	ResetTargetResource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference
type jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) HostedZoneArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) HostedZoneArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) InternalValue() *Route53RecoveryreadinessResourceSetResourcesDnsTargetResource {
	var returns *Route53RecoveryreadinessResourceSetResourcesDnsTargetResource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) RecordSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) RecordSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) RecordType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) RecordTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) TargetResource() Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference {
	var returns Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference
	_jsii_.Get(
		j,
		"targetResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) TargetResourceInput() *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResource {
	var returns *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResource
	_jsii_.Get(
		j,
		"targetResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRoute53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference {
	_init_.Initialize()

	if err := validateNewRoute53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.route53RecoveryreadinessResourceSet.Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRoute53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference_Override(r Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.route53RecoveryreadinessResourceSet.Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference)SetHostedZoneArn(val *string) {
	if err := j.validateSetHostedZoneArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostedZoneArn",
		val,
	)
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference)SetInternalValue(val *Route53RecoveryreadinessResourceSetResourcesDnsTargetResource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference)SetRecordSetId(val *string) {
	if err := j.validateSetRecordSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordSetId",
		val,
	)
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference)SetRecordType(val *string) {
	if err := j.validateSetRecordTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordType",
		val,
	)
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) PutTargetResource(value *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResource) {
	if err := r.validatePutTargetResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTargetResource",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) ResetHostedZoneArn() {
	_jsii_.InvokeVoid(
		r,
		"resetHostedZoneArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) ResetRecordSetId() {
	_jsii_.InvokeVoid(
		r,
		"resetRecordSetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) ResetRecordType() {
	_jsii_.InvokeVoid(
		r,
		"resetRecordType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) ResetTargetResource() {
	_jsii_.InvokeVoid(
		r,
		"resetTargetResource",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

