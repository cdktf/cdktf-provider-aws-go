package appmeshvirtualnode

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/appmeshvirtualnode/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppmeshVirtualNodeSpecListenerOutputReference interface {
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
	ConnectionPool() AppmeshVirtualNodeSpecListenerConnectionPoolOutputReference
	ConnectionPoolInput() *AppmeshVirtualNodeSpecListenerConnectionPool
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	HealthCheck() AppmeshVirtualNodeSpecListenerHealthCheckOutputReference
	HealthCheckInput() *AppmeshVirtualNodeSpecListenerHealthCheck
	InternalValue() *AppmeshVirtualNodeSpecListener
	SetInternalValue(val *AppmeshVirtualNodeSpecListener)
	OutlierDetection() AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference
	OutlierDetectionInput() *AppmeshVirtualNodeSpecListenerOutlierDetection
	PortMapping() AppmeshVirtualNodeSpecListenerPortMappingOutputReference
	PortMappingInput() *AppmeshVirtualNodeSpecListenerPortMapping
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() AppmeshVirtualNodeSpecListenerTimeoutOutputReference
	TimeoutInput() *AppmeshVirtualNodeSpecListenerTimeout
	Tls() AppmeshVirtualNodeSpecListenerTlsOutputReference
	TlsInput() *AppmeshVirtualNodeSpecListenerTls
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
	PutConnectionPool(value *AppmeshVirtualNodeSpecListenerConnectionPool)
	PutHealthCheck(value *AppmeshVirtualNodeSpecListenerHealthCheck)
	PutOutlierDetection(value *AppmeshVirtualNodeSpecListenerOutlierDetection)
	PutPortMapping(value *AppmeshVirtualNodeSpecListenerPortMapping)
	PutTimeout(value *AppmeshVirtualNodeSpecListenerTimeout)
	PutTls(value *AppmeshVirtualNodeSpecListenerTls)
	ResetConnectionPool()
	ResetHealthCheck()
	ResetOutlierDetection()
	ResetTimeout()
	ResetTls()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppmeshVirtualNodeSpecListenerOutputReference
type jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) ConnectionPool() AppmeshVirtualNodeSpecListenerConnectionPoolOutputReference {
	var returns AppmeshVirtualNodeSpecListenerConnectionPoolOutputReference
	_jsii_.Get(
		j,
		"connectionPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) ConnectionPoolInput() *AppmeshVirtualNodeSpecListenerConnectionPool {
	var returns *AppmeshVirtualNodeSpecListenerConnectionPool
	_jsii_.Get(
		j,
		"connectionPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) HealthCheck() AppmeshVirtualNodeSpecListenerHealthCheckOutputReference {
	var returns AppmeshVirtualNodeSpecListenerHealthCheckOutputReference
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) HealthCheckInput() *AppmeshVirtualNodeSpecListenerHealthCheck {
	var returns *AppmeshVirtualNodeSpecListenerHealthCheck
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) InternalValue() *AppmeshVirtualNodeSpecListener {
	var returns *AppmeshVirtualNodeSpecListener
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) OutlierDetection() AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference {
	var returns AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference
	_jsii_.Get(
		j,
		"outlierDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) OutlierDetectionInput() *AppmeshVirtualNodeSpecListenerOutlierDetection {
	var returns *AppmeshVirtualNodeSpecListenerOutlierDetection
	_jsii_.Get(
		j,
		"outlierDetectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) PortMapping() AppmeshVirtualNodeSpecListenerPortMappingOutputReference {
	var returns AppmeshVirtualNodeSpecListenerPortMappingOutputReference
	_jsii_.Get(
		j,
		"portMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) PortMappingInput() *AppmeshVirtualNodeSpecListenerPortMapping {
	var returns *AppmeshVirtualNodeSpecListenerPortMapping
	_jsii_.Get(
		j,
		"portMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) Timeout() AppmeshVirtualNodeSpecListenerTimeoutOutputReference {
	var returns AppmeshVirtualNodeSpecListenerTimeoutOutputReference
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) TimeoutInput() *AppmeshVirtualNodeSpecListenerTimeout {
	var returns *AppmeshVirtualNodeSpecListenerTimeout
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) Tls() AppmeshVirtualNodeSpecListenerTlsOutputReference {
	var returns AppmeshVirtualNodeSpecListenerTlsOutputReference
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) TlsInput() *AppmeshVirtualNodeSpecListenerTls {
	var returns *AppmeshVirtualNodeSpecListenerTls
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}


func NewAppmeshVirtualNodeSpecListenerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppmeshVirtualNodeSpecListenerOutputReference {
	_init_.Initialize()

	if err := validateNewAppmeshVirtualNodeSpecListenerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshVirtualNode.AppmeshVirtualNodeSpecListenerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppmeshVirtualNodeSpecListenerOutputReference_Override(a AppmeshVirtualNodeSpecListenerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshVirtualNode.AppmeshVirtualNodeSpecListenerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference)SetInternalValue(val *AppmeshVirtualNodeSpecListener) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) PutConnectionPool(value *AppmeshVirtualNodeSpecListenerConnectionPool) {
	if err := a.validatePutConnectionPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConnectionPool",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) PutHealthCheck(value *AppmeshVirtualNodeSpecListenerHealthCheck) {
	if err := a.validatePutHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHealthCheck",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) PutOutlierDetection(value *AppmeshVirtualNodeSpecListenerOutlierDetection) {
	if err := a.validatePutOutlierDetectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOutlierDetection",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) PutPortMapping(value *AppmeshVirtualNodeSpecListenerPortMapping) {
	if err := a.validatePutPortMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPortMapping",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) PutTimeout(value *AppmeshVirtualNodeSpecListenerTimeout) {
	if err := a.validatePutTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeout",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) PutTls(value *AppmeshVirtualNodeSpecListenerTls) {
	if err := a.validatePutTlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTls",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) ResetConnectionPool() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionPool",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) ResetOutlierDetection() {
	_jsii_.InvokeVoid(
		a,
		"resetOutlierDetection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) ResetTls() {
	_jsii_.InvokeVoid(
		a,
		"resetTls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

