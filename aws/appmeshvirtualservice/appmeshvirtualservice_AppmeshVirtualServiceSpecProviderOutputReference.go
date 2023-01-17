package appmeshvirtualservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/appmeshvirtualservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppmeshVirtualServiceSpecProviderOutputReference interface {
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
	InternalValue() *AppmeshVirtualServiceSpecProvider
	SetInternalValue(val *AppmeshVirtualServiceSpecProvider)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualNode() AppmeshVirtualServiceSpecProviderVirtualNodeOutputReference
	VirtualNodeInput() *AppmeshVirtualServiceSpecProviderVirtualNode
	VirtualRouter() AppmeshVirtualServiceSpecProviderVirtualRouterOutputReference
	VirtualRouterInput() *AppmeshVirtualServiceSpecProviderVirtualRouter
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
	PutVirtualNode(value *AppmeshVirtualServiceSpecProviderVirtualNode)
	PutVirtualRouter(value *AppmeshVirtualServiceSpecProviderVirtualRouter)
	ResetVirtualNode()
	ResetVirtualRouter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppmeshVirtualServiceSpecProviderOutputReference
type jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) InternalValue() *AppmeshVirtualServiceSpecProvider {
	var returns *AppmeshVirtualServiceSpecProvider
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) VirtualNode() AppmeshVirtualServiceSpecProviderVirtualNodeOutputReference {
	var returns AppmeshVirtualServiceSpecProviderVirtualNodeOutputReference
	_jsii_.Get(
		j,
		"virtualNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) VirtualNodeInput() *AppmeshVirtualServiceSpecProviderVirtualNode {
	var returns *AppmeshVirtualServiceSpecProviderVirtualNode
	_jsii_.Get(
		j,
		"virtualNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) VirtualRouter() AppmeshVirtualServiceSpecProviderVirtualRouterOutputReference {
	var returns AppmeshVirtualServiceSpecProviderVirtualRouterOutputReference
	_jsii_.Get(
		j,
		"virtualRouter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) VirtualRouterInput() *AppmeshVirtualServiceSpecProviderVirtualRouter {
	var returns *AppmeshVirtualServiceSpecProviderVirtualRouter
	_jsii_.Get(
		j,
		"virtualRouterInput",
		&returns,
	)
	return returns
}


func NewAppmeshVirtualServiceSpecProviderOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppmeshVirtualServiceSpecProviderOutputReference {
	_init_.Initialize()

	if err := validateNewAppmeshVirtualServiceSpecProviderOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshVirtualService.AppmeshVirtualServiceSpecProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppmeshVirtualServiceSpecProviderOutputReference_Override(a AppmeshVirtualServiceSpecProviderOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshVirtualService.AppmeshVirtualServiceSpecProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference)SetInternalValue(val *AppmeshVirtualServiceSpecProvider) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) PutVirtualNode(value *AppmeshVirtualServiceSpecProviderVirtualNode) {
	if err := a.validatePutVirtualNodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVirtualNode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) PutVirtualRouter(value *AppmeshVirtualServiceSpecProviderVirtualRouter) {
	if err := a.validatePutVirtualRouterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVirtualRouter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) ResetVirtualNode() {
	_jsii_.InvokeVoid(
		a,
		"resetVirtualNode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) ResetVirtualRouter() {
	_jsii_.InvokeVoid(
		a,
		"resetVirtualRouter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

