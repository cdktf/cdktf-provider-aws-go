package dataawsroute53trafficpolicydocument

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/dataawsroute53trafficpolicydocument/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference interface {
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
	EndpointReference() *string
	SetEndpointReference(val *string)
	EndpointReferenceInput() *string
	EvaluateTargetHealth() interface{}
	SetEvaluateTargetHealth(val interface{})
	EvaluateTargetHealthInput() interface{}
	// Experimental.
	Fqn() *string
	HealthCheck() *string
	SetHealthCheck(val *string)
	HealthCheckInput() *string
	InternalValue() *DataAwsRoute53TrafficPolicyDocumentRulePrimary
	SetInternalValue(val *DataAwsRoute53TrafficPolicyDocumentRulePrimary)
	RuleReference() *string
	SetRuleReference(val *string)
	RuleReferenceInput() *string
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
	ResetEndpointReference()
	ResetEvaluateTargetHealth()
	ResetHealthCheck()
	ResetRuleReference()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference
type jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) EndpointReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) EndpointReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) EvaluateTargetHealth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluateTargetHealth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) EvaluateTargetHealthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluateTargetHealthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) HealthCheck() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) HealthCheckInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) InternalValue() *DataAwsRoute53TrafficPolicyDocumentRulePrimary {
	var returns *DataAwsRoute53TrafficPolicyDocumentRulePrimary
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) RuleReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) RuleReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsRoute53TrafficPolicyDocument.DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference_Override(d DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsRoute53TrafficPolicyDocument.DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference)SetEndpointReference(val *string) {
	if err := j.validateSetEndpointReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointReference",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference)SetEvaluateTargetHealth(val interface{}) {
	if err := j.validateSetEvaluateTargetHealthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluateTargetHealth",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference)SetHealthCheck(val *string) {
	if err := j.validateSetHealthCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheck",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference)SetInternalValue(val *DataAwsRoute53TrafficPolicyDocumentRulePrimary) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference)SetRuleReference(val *string) {
	if err := j.validateSetRuleReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleReference",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) ResetEndpointReference() {
	_jsii_.InvokeVoid(
		d,
		"resetEndpointReference",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) ResetEvaluateTargetHealth() {
	_jsii_.InvokeVoid(
		d,
		"resetEvaluateTargetHealth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		d,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) ResetRuleReference() {
	_jsii_.InvokeVoid(
		d,
		"resetRuleReference",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocumentRulePrimaryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

