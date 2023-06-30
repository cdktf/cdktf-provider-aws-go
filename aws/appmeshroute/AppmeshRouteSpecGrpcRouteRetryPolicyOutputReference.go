package appmeshroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/appmeshroute/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference interface {
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
	GrpcRetryEvents() *[]*string
	SetGrpcRetryEvents(val *[]*string)
	GrpcRetryEventsInput() *[]*string
	HttpRetryEvents() *[]*string
	SetHttpRetryEvents(val *[]*string)
	HttpRetryEventsInput() *[]*string
	InternalValue() *AppmeshRouteSpecGrpcRouteRetryPolicy
	SetInternalValue(val *AppmeshRouteSpecGrpcRouteRetryPolicy)
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	PerRetryTimeout() AppmeshRouteSpecGrpcRouteRetryPolicyPerRetryTimeoutOutputReference
	PerRetryTimeoutInput() *AppmeshRouteSpecGrpcRouteRetryPolicyPerRetryTimeout
	TcpRetryEvents() *[]*string
	SetTcpRetryEvents(val *[]*string)
	TcpRetryEventsInput() *[]*string
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
	PutPerRetryTimeout(value *AppmeshRouteSpecGrpcRouteRetryPolicyPerRetryTimeout)
	ResetGrpcRetryEvents()
	ResetHttpRetryEvents()
	ResetTcpRetryEvents()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference
type jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) GrpcRetryEvents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grpcRetryEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) GrpcRetryEventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grpcRetryEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) HttpRetryEvents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"httpRetryEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) HttpRetryEventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"httpRetryEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) InternalValue() *AppmeshRouteSpecGrpcRouteRetryPolicy {
	var returns *AppmeshRouteSpecGrpcRouteRetryPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) PerRetryTimeout() AppmeshRouteSpecGrpcRouteRetryPolicyPerRetryTimeoutOutputReference {
	var returns AppmeshRouteSpecGrpcRouteRetryPolicyPerRetryTimeoutOutputReference
	_jsii_.Get(
		j,
		"perRetryTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) PerRetryTimeoutInput() *AppmeshRouteSpecGrpcRouteRetryPolicyPerRetryTimeout {
	var returns *AppmeshRouteSpecGrpcRouteRetryPolicyPerRetryTimeout
	_jsii_.Get(
		j,
		"perRetryTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) TcpRetryEvents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tcpRetryEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) TcpRetryEventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tcpRetryEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppmeshRouteSpecGrpcRouteRetryPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewAppmeshRouteSpecGrpcRouteRetryPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshRoute.AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppmeshRouteSpecGrpcRouteRetryPolicyOutputReference_Override(a AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshRoute.AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference)SetGrpcRetryEvents(val *[]*string) {
	if err := j.validateSetGrpcRetryEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grpcRetryEvents",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference)SetHttpRetryEvents(val *[]*string) {
	if err := j.validateSetHttpRetryEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpRetryEvents",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference)SetInternalValue(val *AppmeshRouteSpecGrpcRouteRetryPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference)SetMaxRetries(val *float64) {
	if err := j.validateSetMaxRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference)SetTcpRetryEvents(val *[]*string) {
	if err := j.validateSetTcpRetryEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpRetryEvents",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) PutPerRetryTimeout(value *AppmeshRouteSpecGrpcRouteRetryPolicyPerRetryTimeout) {
	if err := a.validatePutPerRetryTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPerRetryTimeout",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) ResetGrpcRetryEvents() {
	_jsii_.InvokeVoid(
		a,
		"resetGrpcRetryEvents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) ResetHttpRetryEvents() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpRetryEvents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) ResetTcpRetryEvents() {
	_jsii_.InvokeVoid(
		a,
		"resetTcpRetryEvents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

