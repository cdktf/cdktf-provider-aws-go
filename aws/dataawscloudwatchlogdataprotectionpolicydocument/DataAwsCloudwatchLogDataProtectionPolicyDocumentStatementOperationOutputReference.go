// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawscloudwatchlogdataprotectionpolicydocument

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/dataawscloudwatchlogdataprotectionpolicydocument/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference interface {
	cdktf.ComplexObject
	Audit() DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationAuditOutputReference
	AuditInput() *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationAudit
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
	Deidentify() DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationDeidentifyOutputReference
	DeidentifyInput() *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationDeidentify
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperation
	SetInternalValue(val *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperation)
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
	PutAudit(value *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationAudit)
	PutDeidentify(value *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationDeidentify)
	ResetAudit()
	ResetDeidentify()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference
type jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) Audit() DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationAuditOutputReference {
	var returns DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationAuditOutputReference
	_jsii_.Get(
		j,
		"audit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) AuditInput() *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationAudit {
	var returns *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationAudit
	_jsii_.Get(
		j,
		"auditInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) Deidentify() DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationDeidentifyOutputReference {
	var returns DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationDeidentifyOutputReference
	_jsii_.Get(
		j,
		"deidentify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) DeidentifyInput() *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationDeidentify {
	var returns *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationDeidentify
	_jsii_.Get(
		j,
		"deidentifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) InternalValue() *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperation {
	var returns *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsCloudwatchLogDataProtectionPolicyDocument.DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference_Override(d DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsCloudwatchLogDataProtectionPolicyDocument.DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference)SetInternalValue(val *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) PutAudit(value *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationAudit) {
	if err := d.validatePutAuditParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAudit",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) PutDeidentify(value *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationDeidentify) {
	if err := d.validatePutDeidentifyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDeidentify",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) ResetAudit() {
	_jsii_.InvokeVoid(
		d,
		"resetAudit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) ResetDeidentify() {
	_jsii_.InvokeVoid(
		d,
		"resetDeidentify",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

