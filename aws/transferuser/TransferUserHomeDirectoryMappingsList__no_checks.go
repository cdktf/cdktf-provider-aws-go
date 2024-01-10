// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package transferuser

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransferUserHomeDirectoryMappingsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TransferUserHomeDirectoryMappingsList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TransferUserHomeDirectoryMappingsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TransferUserHomeDirectoryMappingsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TransferUserHomeDirectoryMappingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TransferUserHomeDirectoryMappingsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TransferUserHomeDirectoryMappingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTransferUserHomeDirectoryMappingsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

