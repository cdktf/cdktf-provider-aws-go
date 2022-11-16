package cloudfrontcachepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/cloudfrontcachepolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference interface {
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
	CookieBehavior() *string
	SetCookieBehavior(val *string)
	CookieBehaviorInput() *string
	Cookies() CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference
	CookiesInput() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig
	SetInternalValue(val *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig)
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
	PutCookies(value *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies)
	ResetCookies()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference
type jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) CookieBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) CookieBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) Cookies() CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference {
	var returns CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) CookiesInput() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies {
	var returns *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) InternalValue() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig {
	var returns *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfrontCachePolicy.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference_Override(c CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfrontCachePolicy.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference)SetCookieBehavior(val *string) {
	if err := j.validateSetCookieBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cookieBehavior",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference)SetInternalValue(val *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) PutCookies(value *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies) {
	if err := c.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCookies",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		c,
		"resetCookies",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

