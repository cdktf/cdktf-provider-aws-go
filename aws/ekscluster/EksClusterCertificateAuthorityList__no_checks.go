// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ekscluster

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EksClusterCertificateAuthorityList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EksClusterCertificateAuthorityList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EksClusterCertificateAuthorityList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EksClusterCertificateAuthorityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksClusterCertificateAuthorityList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EksClusterCertificateAuthorityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEksClusterCertificateAuthorityListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

