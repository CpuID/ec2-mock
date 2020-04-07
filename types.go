package main

type Region string

// Key = Region
type AvailabilityZones map[Region]AvailabilityZone

type AvailabilityZone struct {
	Name   string
	ZoneId string
}

////////////////////

// Key = Instance ID
type Instances map[string]Instance

type Instance struct {
	// TODO: some fields
	ImageId          string
	InstanceType     string
	PublicIpAddress  string
	PrivateIpAddress string
	SubnetId         string
	VpcId            string
	SecurityGroups   SecurityGroups
	Tags             Tags
}

////////////////////

type SecurityGroups map[string]SecurityGroup

type SecurityGroup struct {
	Id   string
	Name string
}

////////////////////

// Key = Instance ID (for now, could be other objects in future?)
type Tags map[string][]Tag

type Tag struct {
	Key   string
	Value string
}
