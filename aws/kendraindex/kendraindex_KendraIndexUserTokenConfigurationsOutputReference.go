package kendraindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v10/kendraindex/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KendraIndexUserTokenConfigurationsOutputReference interface {
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
	InternalValue() *KendraIndexUserTokenConfigurations
	SetInternalValue(val *KendraIndexUserTokenConfigurations)
	JsonTokenTypeConfiguration() KendraIndexUserTokenConfigurationsJsonTokenTypeConfigurationOutputReference
	JsonTokenTypeConfigurationInput() *KendraIndexUserTokenConfigurationsJsonTokenTypeConfiguration
	JwtTokenTypeConfiguration() KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference
	JwtTokenTypeConfigurationInput() *KendraIndexUserTokenConfigurationsJwtTokenTypeConfiguration
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
	PutJsonTokenTypeConfiguration(value *KendraIndexUserTokenConfigurationsJsonTokenTypeConfiguration)
	PutJwtTokenTypeConfiguration(value *KendraIndexUserTokenConfigurationsJwtTokenTypeConfiguration)
	ResetJsonTokenTypeConfiguration()
	ResetJwtTokenTypeConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraIndexUserTokenConfigurationsOutputReference
type jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) InternalValue() *KendraIndexUserTokenConfigurations {
	var returns *KendraIndexUserTokenConfigurations
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) JsonTokenTypeConfiguration() KendraIndexUserTokenConfigurationsJsonTokenTypeConfigurationOutputReference {
	var returns KendraIndexUserTokenConfigurationsJsonTokenTypeConfigurationOutputReference
	_jsii_.Get(
		j,
		"jsonTokenTypeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) JsonTokenTypeConfigurationInput() *KendraIndexUserTokenConfigurationsJsonTokenTypeConfiguration {
	var returns *KendraIndexUserTokenConfigurationsJsonTokenTypeConfiguration
	_jsii_.Get(
		j,
		"jsonTokenTypeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) JwtTokenTypeConfiguration() KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference {
	var returns KendraIndexUserTokenConfigurationsJwtTokenTypeConfigurationOutputReference
	_jsii_.Get(
		j,
		"jwtTokenTypeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) JwtTokenTypeConfigurationInput() *KendraIndexUserTokenConfigurationsJwtTokenTypeConfiguration {
	var returns *KendraIndexUserTokenConfigurationsJwtTokenTypeConfiguration
	_jsii_.Get(
		j,
		"jwtTokenTypeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKendraIndexUserTokenConfigurationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KendraIndexUserTokenConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewKendraIndexUserTokenConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kendraIndex.KendraIndexUserTokenConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraIndexUserTokenConfigurationsOutputReference_Override(k KendraIndexUserTokenConfigurationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kendraIndex.KendraIndexUserTokenConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference)SetInternalValue(val *KendraIndexUserTokenConfigurations) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) PutJsonTokenTypeConfiguration(value *KendraIndexUserTokenConfigurationsJsonTokenTypeConfiguration) {
	if err := k.validatePutJsonTokenTypeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putJsonTokenTypeConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) PutJwtTokenTypeConfiguration(value *KendraIndexUserTokenConfigurationsJwtTokenTypeConfiguration) {
	if err := k.validatePutJwtTokenTypeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putJwtTokenTypeConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) ResetJsonTokenTypeConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetJsonTokenTypeConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) ResetJwtTokenTypeConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetJwtTokenTypeConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraIndexUserTokenConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

