package route53recoveryreadinessresourceset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v13/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v13/route53recoveryreadinessresourceset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference interface {
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
	InternalValue() *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResource
	SetInternalValue(val *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResource)
	NlbResource() Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceNlbResourceOutputReference
	NlbResourceInput() *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceNlbResource
	R53Resource() Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceR53ResourceOutputReference
	R53ResourceInput() *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceR53Resource
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
	PutNlbResource(value *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceNlbResource)
	PutR53Resource(value *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceR53Resource)
	ResetNlbResource()
	ResetR53Resource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference
type jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) InternalValue() *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResource {
	var returns *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) NlbResource() Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceNlbResourceOutputReference {
	var returns Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceNlbResourceOutputReference
	_jsii_.Get(
		j,
		"nlbResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) NlbResourceInput() *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceNlbResource {
	var returns *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceNlbResource
	_jsii_.Get(
		j,
		"nlbResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) R53Resource() Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceR53ResourceOutputReference {
	var returns Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceR53ResourceOutputReference
	_jsii_.Get(
		j,
		"r53Resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) R53ResourceInput() *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceR53Resource {
	var returns *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceR53Resource
	_jsii_.Get(
		j,
		"r53ResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRoute53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference {
	_init_.Initialize()

	if err := validateNewRoute53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.route53RecoveryreadinessResourceSet.Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRoute53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference_Override(r Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.route53RecoveryreadinessResourceSet.Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference)SetInternalValue(val *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) PutNlbResource(value *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceNlbResource) {
	if err := r.validatePutNlbResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putNlbResource",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) PutR53Resource(value *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceR53Resource) {
	if err := r.validatePutR53ResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putR53Resource",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) ResetNlbResource() {
	_jsii_.InvokeVoid(
		r,
		"resetNlbResource",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) ResetR53Resource() {
	_jsii_.InvokeVoid(
		r,
		"resetR53Resource",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

