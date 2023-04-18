package opensearchdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/opensearchdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpensearchDomainClusterConfigOutputReference interface {
	cdktf.ComplexObject
	ColdStorageOptions() OpensearchDomainClusterConfigColdStorageOptionsOutputReference
	ColdStorageOptionsInput() *OpensearchDomainClusterConfigColdStorageOptions
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
	DedicatedMasterCount() *float64
	SetDedicatedMasterCount(val *float64)
	DedicatedMasterCountInput() *float64
	DedicatedMasterEnabled() interface{}
	SetDedicatedMasterEnabled(val interface{})
	DedicatedMasterEnabledInput() interface{}
	DedicatedMasterType() *string
	SetDedicatedMasterType(val *string)
	DedicatedMasterTypeInput() *string
	// Experimental.
	Fqn() *string
	InstanceCount() *float64
	SetInstanceCount(val *float64)
	InstanceCountInput() *float64
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() *OpensearchDomainClusterConfig
	SetInternalValue(val *OpensearchDomainClusterConfig)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WarmCount() *float64
	SetWarmCount(val *float64)
	WarmCountInput() *float64
	WarmEnabled() interface{}
	SetWarmEnabled(val interface{})
	WarmEnabledInput() interface{}
	WarmType() *string
	SetWarmType(val *string)
	WarmTypeInput() *string
	ZoneAwarenessConfig() OpensearchDomainClusterConfigZoneAwarenessConfigOutputReference
	ZoneAwarenessConfigInput() *OpensearchDomainClusterConfigZoneAwarenessConfig
	ZoneAwarenessEnabled() interface{}
	SetZoneAwarenessEnabled(val interface{})
	ZoneAwarenessEnabledInput() interface{}
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
	PutColdStorageOptions(value *OpensearchDomainClusterConfigColdStorageOptions)
	PutZoneAwarenessConfig(value *OpensearchDomainClusterConfigZoneAwarenessConfig)
	ResetColdStorageOptions()
	ResetDedicatedMasterCount()
	ResetDedicatedMasterEnabled()
	ResetDedicatedMasterType()
	ResetInstanceCount()
	ResetInstanceType()
	ResetWarmCount()
	ResetWarmEnabled()
	ResetWarmType()
	ResetZoneAwarenessConfig()
	ResetZoneAwarenessEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OpensearchDomainClusterConfigOutputReference
type jsiiProxy_OpensearchDomainClusterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ColdStorageOptions() OpensearchDomainClusterConfigColdStorageOptionsOutputReference {
	var returns OpensearchDomainClusterConfigColdStorageOptionsOutputReference
	_jsii_.Get(
		j,
		"coldStorageOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ColdStorageOptionsInput() *OpensearchDomainClusterConfigColdStorageOptions {
	var returns *OpensearchDomainClusterConfigColdStorageOptions
	_jsii_.Get(
		j,
		"coldStorageOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) DedicatedMasterCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dedicatedMasterCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) DedicatedMasterCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dedicatedMasterCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) DedicatedMasterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedMasterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) DedicatedMasterEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedMasterEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) DedicatedMasterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedMasterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) DedicatedMasterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedMasterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) InternalValue() *OpensearchDomainClusterConfig {
	var returns *OpensearchDomainClusterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) WarmCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warmCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) WarmCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warmCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) WarmEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"warmEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) WarmEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"warmEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) WarmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) WarmTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warmTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ZoneAwarenessConfig() OpensearchDomainClusterConfigZoneAwarenessConfigOutputReference {
	var returns OpensearchDomainClusterConfigZoneAwarenessConfigOutputReference
	_jsii_.Get(
		j,
		"zoneAwarenessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ZoneAwarenessConfigInput() *OpensearchDomainClusterConfigZoneAwarenessConfig {
	var returns *OpensearchDomainClusterConfigZoneAwarenessConfig
	_jsii_.Get(
		j,
		"zoneAwarenessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ZoneAwarenessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneAwarenessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ZoneAwarenessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneAwarenessEnabledInput",
		&returns,
	)
	return returns
}


func NewOpensearchDomainClusterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OpensearchDomainClusterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewOpensearchDomainClusterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpensearchDomainClusterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.opensearchDomain.OpensearchDomainClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOpensearchDomainClusterConfigOutputReference_Override(o OpensearchDomainClusterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.opensearchDomain.OpensearchDomainClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference)SetDedicatedMasterCount(val *float64) {
	if err := j.validateSetDedicatedMasterCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedMasterCount",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference)SetDedicatedMasterEnabled(val interface{}) {
	if err := j.validateSetDedicatedMasterEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedMasterEnabled",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference)SetDedicatedMasterType(val *string) {
	if err := j.validateSetDedicatedMasterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedMasterType",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference)SetInstanceCount(val *float64) {
	if err := j.validateSetInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference)SetInternalValue(val *OpensearchDomainClusterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference)SetWarmCount(val *float64) {
	if err := j.validateSetWarmCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warmCount",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference)SetWarmEnabled(val interface{}) {
	if err := j.validateSetWarmEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warmEnabled",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference)SetWarmType(val *string) {
	if err := j.validateSetWarmTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warmType",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainClusterConfigOutputReference)SetZoneAwarenessEnabled(val interface{}) {
	if err := j.validateSetZoneAwarenessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneAwarenessEnabled",
		val,
	)
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) PutColdStorageOptions(value *OpensearchDomainClusterConfigColdStorageOptions) {
	if err := o.validatePutColdStorageOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putColdStorageOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) PutZoneAwarenessConfig(value *OpensearchDomainClusterConfigZoneAwarenessConfig) {
	if err := o.validatePutZoneAwarenessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putZoneAwarenessConfig",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ResetColdStorageOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetColdStorageOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ResetDedicatedMasterCount() {
	_jsii_.InvokeVoid(
		o,
		"resetDedicatedMasterCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ResetDedicatedMasterEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetDedicatedMasterEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ResetDedicatedMasterType() {
	_jsii_.InvokeVoid(
		o,
		"resetDedicatedMasterType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ResetInstanceCount() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ResetWarmCount() {
	_jsii_.InvokeVoid(
		o,
		"resetWarmCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ResetWarmEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetWarmEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ResetWarmType() {
	_jsii_.InvokeVoid(
		o,
		"resetWarmType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ResetZoneAwarenessConfig() {
	_jsii_.InvokeVoid(
		o,
		"resetZoneAwarenessConfig",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ResetZoneAwarenessEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetZoneAwarenessEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainClusterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

