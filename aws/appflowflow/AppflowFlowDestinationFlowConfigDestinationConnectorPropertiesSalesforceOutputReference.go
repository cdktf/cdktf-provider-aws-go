// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/appflowflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference interface {
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
	DataTransferApi() *string
	SetDataTransferApi(val *string)
	DataTransferApiInput() *string
	ErrorHandlingConfig() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfigOutputReference
	ErrorHandlingConfigInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfig
	// Experimental.
	Fqn() *string
	IdFieldNames() *[]*string
	SetIdFieldNames(val *[]*string)
	IdFieldNamesInput() *[]*string
	InternalValue() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforce
	SetInternalValue(val *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforce)
	Object() *string
	SetObject(val *string)
	ObjectInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WriteOperationType() *string
	SetWriteOperationType(val *string)
	WriteOperationTypeInput() *string
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
	PutErrorHandlingConfig(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfig)
	ResetDataTransferApi()
	ResetErrorHandlingConfig()
	ResetIdFieldNames()
	ResetWriteOperationType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference
type jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) DataTransferApi() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTransferApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) DataTransferApiInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTransferApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) ErrorHandlingConfig() AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfigOutputReference {
	var returns AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfigOutputReference
	_jsii_.Get(
		j,
		"errorHandlingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) ErrorHandlingConfigInput() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfig {
	var returns *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfig
	_jsii_.Get(
		j,
		"errorHandlingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) IdFieldNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"idFieldNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) IdFieldNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"idFieldNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) InternalValue() *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforce {
	var returns *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforce
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) Object() *string {
	var returns *string
	_jsii_.Get(
		j,
		"object",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) ObjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) WriteOperationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeOperationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) WriteOperationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeOperationTypeInput",
		&returns,
	)
	return returns
}


func NewAppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference {
	_init_.Initialize()

	if err := validateNewAppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference_Override(a AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference)SetDataTransferApi(val *string) {
	if err := j.validateSetDataTransferApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataTransferApi",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference)SetIdFieldNames(val *[]*string) {
	if err := j.validateSetIdFieldNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idFieldNames",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference)SetInternalValue(val *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforce) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference)SetObject(val *string) {
	if err := j.validateSetObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"object",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference)SetWriteOperationType(val *string) {
	if err := j.validateSetWriteOperationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeOperationType",
		val,
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) PutErrorHandlingConfig(value *AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfig) {
	if err := a.validatePutErrorHandlingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putErrorHandlingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) ResetDataTransferApi() {
	_jsii_.InvokeVoid(
		a,
		"resetDataTransferApi",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) ResetErrorHandlingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetErrorHandlingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) ResetIdFieldNames() {
	_jsii_.InvokeVoid(
		a,
		"resetIdFieldNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) ResetWriteOperationType() {
	_jsii_.InvokeVoid(
		a,
		"resetWriteOperationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

