// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datazonedomain

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatazoneDomainSingleSignOnList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatazoneDomainSingleSignOnList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatazoneDomainSingleSignOnList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatazoneDomainSingleSignOnList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatazoneDomainSingleSignOnList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatazoneDomainSingleSignOnList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatazoneDomainSingleSignOnList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatazoneDomainSingleSignOnListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

