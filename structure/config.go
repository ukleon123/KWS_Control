package structure

type Config struct {
	VmInternalSubnets []string `yaml:"vm_internal_subnets"`
	Cores             []string `yaml:"cores"`
}