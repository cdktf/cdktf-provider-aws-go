// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package s3tablestable

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3TablesTableMetadataList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_S3TablesTableMetadataList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3TablesTableMetadataList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3TablesTableMetadataList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3TablesTableMetadataList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3TablesTableMetadataList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3TablesTableMetadataList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3TablesTableMetadataListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

