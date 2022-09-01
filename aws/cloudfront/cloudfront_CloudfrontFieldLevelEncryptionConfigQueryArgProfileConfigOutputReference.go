package cloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/cloudfront/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference interface {
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
	ForwardWhenQueryArgProfileIsUnknown() interface{}
	SetForwardWhenQueryArgProfileIsUnknown(val interface{})
	ForwardWhenQueryArgProfileIsUnknownInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfig
	SetInternalValue(val *CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfig)
	QueryArgProfiles() CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfilesOutputReference
	QueryArgProfilesInput() *CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfiles
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
	PutQueryArgProfiles(value *CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfiles)
	ResetQueryArgProfiles()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference
type jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) ForwardWhenQueryArgProfileIsUnknown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardWhenQueryArgProfileIsUnknown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) ForwardWhenQueryArgProfileIsUnknownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardWhenQueryArgProfileIsUnknownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) InternalValue() *CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfig {
	var returns *CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) QueryArgProfiles() CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfilesOutputReference {
	var returns CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfilesOutputReference
	_jsii_.Get(
		j,
		"queryArgProfiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) QueryArgProfilesInput() *CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfiles {
	var returns *CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfiles
	_jsii_.Get(
		j,
		"queryArgProfilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfront.CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference_Override(c CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfront.CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference)SetForwardWhenQueryArgProfileIsUnknown(val interface{}) {
	if err := j.validateSetForwardWhenQueryArgProfileIsUnknownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardWhenQueryArgProfileIsUnknown",
		val,
	)
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference)SetInternalValue(val *CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) PutQueryArgProfiles(value *CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfiles) {
	if err := c.validatePutQueryArgProfilesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putQueryArgProfiles",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) ResetQueryArgProfiles() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryArgProfiles",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionConfigQueryArgProfileConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

