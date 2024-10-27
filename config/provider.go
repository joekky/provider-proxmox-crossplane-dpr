// /*
// Copyright 2021 Upbound Inc.
// */

// package config

// import (
// 	// Note(turkenh): we are importing this to embed provider schema document
// 	_ "embed"

// 	ujconfig "github.com/crossplane/upjet/pkg/config"
// 	"github.com/joekky/provider-proxmox-crossplane/config/proxmox"
// )

// const (
// 	resourcePrefix = "proxmox"
// 	modulePath     = "github.com/joekky/provider-proxmox-crossplane"
// )

// //go:embed schema.json
// var providerSchema string

// //go:embed provider-metadata.yaml
// var providerMetadata string

// // GetProvider returns provider configuration
// func GetProvider() *ujconfig.Provider {
// 	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
// 		ujconfig.WithRootGroup("proxmox.crossplane.io"),
// 		ujconfig.WithIncludeList(ExternalNameConfigured()),
// 		ujconfig.WithFeaturesPackage("internal/features"),
// 		ujconfig.WithDefaultResourceOptions(
// 			ExternalNameConfigurations(),
// 		))

// 	for _, configure := range []func(provider *ujconfig.Provider){
// 		proxmox.Configure,
// 	} {
// 		configure(pc)
// 	}

// 	pc.ConfigureResources()
// 	return pc
// }

package config

import (
	"github.com/joekky/provider-proxmox-crossplane/config/container"
	"github.com/joekky/provider-proxmox-crossplane/config/virtualmachine"
	"github.com/upbound/upjet/pkg/config"
)

const (
	modulePath = "github.com/joekky/provider-proxmox-crossplane"
)

// GetProvider returns provider configuration
func GetProvider() *config.Provider {
	pc := config.NewProvider([]config.ExternalName{
		config.NameAsIdentifier,
	})

	for _, configure := range []func(provider *config.Provider){
		// Configure resource groups
		virtualmachine.Configure,
		container.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
