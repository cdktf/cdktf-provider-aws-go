package servicecatalogprovisionedproduct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/servicecatalogprovisionedproduct/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference interface {
	cdktf.ComplexObject
	Accounts() *[]*string
	SetAccounts(val *[]*string)
	AccountsInput() *[]*string
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
	FailureToleranceCount() *float64
	SetFailureToleranceCount(val *float64)
	FailureToleranceCountInput() *float64
	FailureTolerancePercentage() *float64
	SetFailureTolerancePercentage(val *float64)
	FailureTolerancePercentageInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *ServicecatalogProvisionedProductStackSetProvisioningPreferences
	SetInternalValue(val *ServicecatalogProvisionedProductStackSetProvisioningPreferences)
	MaxConcurrencyCount() *float64
	SetMaxConcurrencyCount(val *float64)
	MaxConcurrencyCountInput() *float64
	MaxConcurrencyPercentage() *float64
	SetMaxConcurrencyPercentage(val *float64)
	MaxConcurrencyPercentageInput() *float64
	Regions() *[]*string
	SetRegions(val *[]*string)
	RegionsInput() *[]*string
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
	ResetAccounts()
	ResetFailureToleranceCount()
	ResetFailureTolerancePercentage()
	ResetMaxConcurrencyCount()
	ResetMaxConcurrencyPercentage()
	ResetRegions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference
type jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) Accounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) AccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) FailureToleranceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureToleranceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) FailureToleranceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureToleranceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) FailureTolerancePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureTolerancePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) FailureTolerancePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureTolerancePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) InternalValue() *ServicecatalogProvisionedProductStackSetProvisioningPreferences {
	var returns *ServicecatalogProvisionedProductStackSetProvisioningPreferences
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) MaxConcurrencyCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) MaxConcurrencyCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) MaxConcurrencyPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) MaxConcurrencyPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference {
	_init_.Initialize()

	if err := validateNewServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.servicecatalogProvisionedProduct.ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference_Override(s ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.servicecatalogProvisionedProduct.ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference)SetAccounts(val *[]*string) {
	if err := j.validateSetAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accounts",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference)SetFailureToleranceCount(val *float64) {
	if err := j.validateSetFailureToleranceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureToleranceCount",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference)SetFailureTolerancePercentage(val *float64) {
	if err := j.validateSetFailureTolerancePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureTolerancePercentage",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference)SetInternalValue(val *ServicecatalogProvisionedProductStackSetProvisioningPreferences) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference)SetMaxConcurrencyCount(val *float64) {
	if err := j.validateSetMaxConcurrencyCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrencyCount",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference)SetMaxConcurrencyPercentage(val *float64) {
	if err := j.validateSetMaxConcurrencyPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrencyPercentage",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference)SetRegions(val *[]*string) {
	if err := j.validateSetRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) ResetAccounts() {
	_jsii_.InvokeVoid(
		s,
		"resetAccounts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) ResetFailureToleranceCount() {
	_jsii_.InvokeVoid(
		s,
		"resetFailureToleranceCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) ResetFailureTolerancePercentage() {
	_jsii_.InvokeVoid(
		s,
		"resetFailureTolerancePercentage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) ResetMaxConcurrencyCount() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxConcurrencyCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) ResetMaxConcurrencyPercentage() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxConcurrencyPercentage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) ResetRegions() {
	_jsii_.InvokeVoid(
		s,
		"resetRegions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogProvisionedProductStackSetProvisioningPreferencesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

