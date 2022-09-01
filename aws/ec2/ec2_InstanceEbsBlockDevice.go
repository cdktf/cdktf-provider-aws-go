package ec2


type InstanceEbsBlockDevice struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/instance#device_name Instance#device_name}.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/instance#delete_on_termination Instance#delete_on_termination}.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/instance#encrypted Instance#encrypted}.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/instance#iops Instance#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/instance#kms_key_id Instance#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/instance#snapshot_id Instance#snapshot_id}.
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/instance#tags Instance#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/instance#throughput Instance#throughput}.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/instance#volume_size Instance#volume_size}.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/instance#volume_type Instance#volume_type}.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

