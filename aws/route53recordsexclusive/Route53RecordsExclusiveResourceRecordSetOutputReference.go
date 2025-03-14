// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53recordsexclusive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/route53recordsexclusive/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53RecordsExclusiveResourceRecordSetOutputReference interface {
	cdktf.ComplexObject
	AliasTarget() Route53RecordsExclusiveResourceRecordSetAliasTargetList
	AliasTargetInput() interface{}
	CidrRoutingConfig() Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList
	CidrRoutingConfigInput() interface{}
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
	Failover() *string
	SetFailover(val *string)
	FailoverInput() *string
	// Experimental.
	Fqn() *string
	Geolocation() Route53RecordsExclusiveResourceRecordSetGeolocationList
	GeolocationInput() interface{}
	GeoproximityLocation() Route53RecordsExclusiveResourceRecordSetGeoproximityLocationList
	GeoproximityLocationInput() interface{}
	HealthCheckId() *string
	SetHealthCheckId(val *string)
	HealthCheckIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MultiValueAnswer() interface{}
	SetMultiValueAnswer(val interface{})
	MultiValueAnswerInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ResourceRecords() Route53RecordsExclusiveResourceRecordSetResourceRecordsList
	ResourceRecordsInput() interface{}
	SetIdentifier() *string
	SetSetIdentifier(val *string)
	SetIdentifierInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrafficPolicyInstanceId() *string
	SetTrafficPolicyInstanceId(val *string)
	TrafficPolicyInstanceIdInput() *string
	Ttl() *float64
	SetTtl(val *float64)
	TtlInput() *float64
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Weight() *float64
	SetWeight(val *float64)
	WeightInput() *float64
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
	PutAliasTarget(value interface{})
	PutCidrRoutingConfig(value interface{})
	PutGeolocation(value interface{})
	PutGeoproximityLocation(value interface{})
	PutResourceRecords(value interface{})
	ResetAliasTarget()
	ResetCidrRoutingConfig()
	ResetFailover()
	ResetGeolocation()
	ResetGeoproximityLocation()
	ResetHealthCheckId()
	ResetMultiValueAnswer()
	ResetRegion()
	ResetResourceRecords()
	ResetSetIdentifier()
	ResetTrafficPolicyInstanceId()
	ResetTtl()
	ResetType()
	ResetWeight()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Route53RecordsExclusiveResourceRecordSetOutputReference
type jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) AliasTarget() Route53RecordsExclusiveResourceRecordSetAliasTargetList {
	var returns Route53RecordsExclusiveResourceRecordSetAliasTargetList
	_jsii_.Get(
		j,
		"aliasTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) AliasTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aliasTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) CidrRoutingConfig() Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList {
	var returns Route53RecordsExclusiveResourceRecordSetCidrRoutingConfigList
	_jsii_.Get(
		j,
		"cidrRoutingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) CidrRoutingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cidrRoutingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) Failover() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) FailoverInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) Geolocation() Route53RecordsExclusiveResourceRecordSetGeolocationList {
	var returns Route53RecordsExclusiveResourceRecordSetGeolocationList
	_jsii_.Get(
		j,
		"geolocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) GeolocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geolocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) GeoproximityLocation() Route53RecordsExclusiveResourceRecordSetGeoproximityLocationList {
	var returns Route53RecordsExclusiveResourceRecordSetGeoproximityLocationList
	_jsii_.Get(
		j,
		"geoproximityLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) GeoproximityLocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoproximityLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) HealthCheckId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) HealthCheckIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) MultiValueAnswer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiValueAnswer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) MultiValueAnswerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiValueAnswerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) ResourceRecords() Route53RecordsExclusiveResourceRecordSetResourceRecordsList {
	var returns Route53RecordsExclusiveResourceRecordSetResourceRecordsList
	_jsii_.Get(
		j,
		"resourceRecords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) ResourceRecordsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceRecordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) SetIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) SetIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) TrafficPolicyInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficPolicyInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) TrafficPolicyInstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficPolicyInstanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) TtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) WeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightInput",
		&returns,
	)
	return returns
}


func NewRoute53RecordsExclusiveResourceRecordSetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Route53RecordsExclusiveResourceRecordSetOutputReference {
	_init_.Initialize()

	if err := validateNewRoute53RecordsExclusiveResourceRecordSetOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.route53RecordsExclusive.Route53RecordsExclusiveResourceRecordSetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewRoute53RecordsExclusiveResourceRecordSetOutputReference_Override(r Route53RecordsExclusiveResourceRecordSetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.route53RecordsExclusive.Route53RecordsExclusiveResourceRecordSetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference)SetFailover(val *string) {
	if err := j.validateSetFailoverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failover",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference)SetHealthCheckId(val *string) {
	if err := j.validateSetHealthCheckIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckId",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference)SetMultiValueAnswer(val interface{}) {
	if err := j.validateSetMultiValueAnswerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiValueAnswer",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference)SetSetIdentifier(val *string) {
	if err := j.validateSetSetIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"setIdentifier",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference)SetTrafficPolicyInstanceId(val *string) {
	if err := j.validateSetTrafficPolicyInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficPolicyInstanceId",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference)SetTtl(val *float64) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference)SetWeight(val *float64) {
	if err := j.validateSetWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) PutAliasTarget(value interface{}) {
	if err := r.validatePutAliasTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAliasTarget",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) PutCidrRoutingConfig(value interface{}) {
	if err := r.validatePutCidrRoutingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putCidrRoutingConfig",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) PutGeolocation(value interface{}) {
	if err := r.validatePutGeolocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putGeolocation",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) PutGeoproximityLocation(value interface{}) {
	if err := r.validatePutGeoproximityLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putGeoproximityLocation",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) PutResourceRecords(value interface{}) {
	if err := r.validatePutResourceRecordsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putResourceRecords",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) ResetAliasTarget() {
	_jsii_.InvokeVoid(
		r,
		"resetAliasTarget",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) ResetCidrRoutingConfig() {
	_jsii_.InvokeVoid(
		r,
		"resetCidrRoutingConfig",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) ResetFailover() {
	_jsii_.InvokeVoid(
		r,
		"resetFailover",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) ResetGeolocation() {
	_jsii_.InvokeVoid(
		r,
		"resetGeolocation",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) ResetGeoproximityLocation() {
	_jsii_.InvokeVoid(
		r,
		"resetGeoproximityLocation",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) ResetHealthCheckId() {
	_jsii_.InvokeVoid(
		r,
		"resetHealthCheckId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) ResetMultiValueAnswer() {
	_jsii_.InvokeVoid(
		r,
		"resetMultiValueAnswer",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		r,
		"resetRegion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) ResetResourceRecords() {
	_jsii_.InvokeVoid(
		r,
		"resetResourceRecords",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) ResetSetIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetSetIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) ResetTrafficPolicyInstanceId() {
	_jsii_.InvokeVoid(
		r,
		"resetTrafficPolicyInstanceId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) ResetTtl() {
	_jsii_.InvokeVoid(
		r,
		"resetTtl",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		r,
		"resetType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) ResetWeight() {
	_jsii_.InvokeVoid(
		r,
		"resetWeight",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_Route53RecordsExclusiveResourceRecordSetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

