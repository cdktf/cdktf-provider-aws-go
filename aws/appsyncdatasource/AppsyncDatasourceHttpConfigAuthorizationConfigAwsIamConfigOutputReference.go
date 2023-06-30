package appsyncdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/appsyncdatasource/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference interface {
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
	InternalValue() *AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfig
	SetInternalValue(val *AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfig)
	SigningRegion() *string
	SetSigningRegion(val *string)
	SigningRegionInput() *string
	SigningServiceName() *string
	SetSigningServiceName(val *string)
	SigningServiceNameInput() *string
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
	ResetSigningRegion()
	ResetSigningServiceName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference
type jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) InternalValue() *AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfig {
	var returns *AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) SigningRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) SigningRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) SigningServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) SigningServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingServiceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appsyncDatasource.AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference_Override(a AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appsyncDatasource.AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference)SetInternalValue(val *AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference)SetSigningRegion(val *string) {
	if err := j.validateSetSigningRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signingRegion",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference)SetSigningServiceName(val *string) {
	if err := j.validateSetSigningServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signingServiceName",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) ResetSigningRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetSigningRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) ResetSigningServiceName() {
	_jsii_.InvokeVoid(
		a,
		"resetSigningServiceName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

