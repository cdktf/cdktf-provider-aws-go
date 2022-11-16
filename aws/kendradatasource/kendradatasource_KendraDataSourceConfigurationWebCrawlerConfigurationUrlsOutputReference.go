package kendradatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/kendradatasource/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference interface {
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
	InternalValue() *KendraDataSourceConfigurationWebCrawlerConfigurationUrls
	SetInternalValue(val *KendraDataSourceConfigurationWebCrawlerConfigurationUrls)
	SeedUrlConfiguration() KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationOutputReference
	SeedUrlConfigurationInput() *KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfiguration
	SiteMapsConfiguration() KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSiteMapsConfigurationOutputReference
	SiteMapsConfigurationInput() *KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSiteMapsConfiguration
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
	PutSeedUrlConfiguration(value *KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfiguration)
	PutSiteMapsConfiguration(value *KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSiteMapsConfiguration)
	ResetSeedUrlConfiguration()
	ResetSiteMapsConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference
type jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) InternalValue() *KendraDataSourceConfigurationWebCrawlerConfigurationUrls {
	var returns *KendraDataSourceConfigurationWebCrawlerConfigurationUrls
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) SeedUrlConfiguration() KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationOutputReference {
	var returns KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationOutputReference
	_jsii_.Get(
		j,
		"seedUrlConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) SeedUrlConfigurationInput() *KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfiguration {
	var returns *KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfiguration
	_jsii_.Get(
		j,
		"seedUrlConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) SiteMapsConfiguration() KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSiteMapsConfigurationOutputReference {
	var returns KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSiteMapsConfigurationOutputReference
	_jsii_.Get(
		j,
		"siteMapsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) SiteMapsConfigurationInput() *KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSiteMapsConfiguration {
	var returns *KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSiteMapsConfiguration
	_jsii_.Get(
		j,
		"siteMapsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference {
	_init_.Initialize()

	if err := validateNewKendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kendraDataSource.KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference_Override(k KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kendraDataSource.KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference)SetInternalValue(val *KendraDataSourceConfigurationWebCrawlerConfigurationUrls) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) PutSeedUrlConfiguration(value *KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfiguration) {
	if err := k.validatePutSeedUrlConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSeedUrlConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) PutSiteMapsConfiguration(value *KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSiteMapsConfiguration) {
	if err := k.validatePutSiteMapsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSiteMapsConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) ResetSeedUrlConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetSeedUrlConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) ResetSiteMapsConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetSiteMapsConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KendraDataSourceConfigurationWebCrawlerConfigurationUrlsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

