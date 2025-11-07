// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dsqlcluster

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DsqlClusterMultiRegionPropertiesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DsqlClusterMultiRegionPropertiesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DsqlClusterMultiRegionPropertiesList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DsqlClusterMultiRegionPropertiesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DsqlClusterMultiRegionPropertiesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DsqlClusterMultiRegionPropertiesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DsqlClusterMultiRegionPropertiesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDsqlClusterMultiRegionPropertiesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

