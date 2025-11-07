// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appfabricingestiondestination

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/appfabricingestiondestination/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList
type jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList {
	_init_.Initialize()

	if err := validateNewAppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList{}

	_jsii_.Create(
		"@cdktf/provider-aws.appfabricIngestionDestination.AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList_Override(a AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appfabricIngestionDestination.AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := a.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		a,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList) Get(index *float64) AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

