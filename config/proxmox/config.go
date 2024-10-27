package proxmox

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("proxmox_vm_qemu", func(r *ujconfig.Resource) {
		r.ShortGroup = "compute"
		r.Kind = "VirtualMachine"
		r.Version = "v1alpha1"
		r.ExternalName = ujconfig.IdentifierFromProvider
	})
}
