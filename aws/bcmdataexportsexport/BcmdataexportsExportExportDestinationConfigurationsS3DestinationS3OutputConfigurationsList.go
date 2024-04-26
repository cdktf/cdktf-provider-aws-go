// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bcmdataexportsexport

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/bcmdataexportsexport/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList interface {
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
	Get(index *float64) BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList
type jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList {
	_init_.Initialize()

	if err := validateNewBcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList{}

	_jsii_.Create(
		"@cdktf/provider-aws.bcmdataexportsExport.BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList_Override(b BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.bcmdataexportsExport.BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := b.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		b,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList) Get(index *float64) BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsOutputReference {
	if err := b.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

