// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"os"
	"testing"

	"github.com/tj/assert"

	"github.com/vmware-tanzu/tanzu-framework/cli/core/pkg/constants"
	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
)

func TestConfigPopulateDefaultStandaloneDiscovery(t *testing.T) {
	cfg := &configapi.ClientConfig{
		ClientOptions: &configapi.ClientOptions{
			CLI: &configapi.CLIOptions{
				DiscoverySources: []configapi.PluginDiscovery{},
			},
		},
	}
	configureTestDefaultStandaloneDiscoveryOCI()

	assert := assert.New(t)

	added := populateDefaultStandaloneDiscovery(cfg)
	assert.Equal(true, added)
	assert.Equal(len(cfg.ClientOptions.CLI.DiscoverySources), 1)
	assert.Equal(cfg.ClientOptions.CLI.DiscoverySources[0].OCI.Name, DefaultStandaloneDiscoveryName)
	assert.Equal(cfg.ClientOptions.CLI.DiscoverySources[0].OCI.Image, "fake.image.repo/package/standalone-plugins:v1.0.0")
}

func TestConfigPopulateDefaultStandaloneDiscoveryWhenDefaultDiscoveryExistsAndIsSame(t *testing.T) {
	cfg := &configapi.ClientConfig{
		ClientOptions: &configapi.ClientOptions{
			CLI: &configapi.CLIOptions{
				DiscoverySources: []configapi.PluginDiscovery{
					configapi.PluginDiscovery{
						OCI: &configapi.OCIDiscovery{
							Name:  DefaultStandaloneDiscoveryName,
							Image: "fake.image.repo/package/standalone-plugins:v1.0.0",
						},
					},
				},
			},
		},
	}
	configureTestDefaultStandaloneDiscoveryOCI()

	assert := assert.New(t)

	added := populateDefaultStandaloneDiscovery(cfg)
	assert.Equal(false, added)
	assert.Equal(len(cfg.ClientOptions.CLI.DiscoverySources), 1)
	assert.Equal(cfg.ClientOptions.CLI.DiscoverySources[0].OCI.Name, DefaultStandaloneDiscoveryName)
	assert.Equal(cfg.ClientOptions.CLI.DiscoverySources[0].OCI.Image, "fake.image.repo/package/standalone-plugins:v1.0.0")
}

func TestConfigPopulateDefaultStandaloneDiscoveryWhenDefaultDiscoveryExistsAndIsNotSame(t *testing.T) {
	cfg := &configapi.ClientConfig{
		ClientOptions: &configapi.ClientOptions{
			CLI: &configapi.CLIOptions{
				DiscoverySources: []configapi.PluginDiscovery{
					configapi.PluginDiscovery{
						OCI: &configapi.OCIDiscovery{
							Name:  DefaultStandaloneDiscoveryName,
							Image: "fake.image/path:v2.0.0",
						},
					},
					configapi.PluginDiscovery{
						OCI: &configapi.OCIDiscovery{
							Name:  "additional-discovery",
							Image: "additional-discovery/path:v1.0.0",
						},
					},
				},
			},
		},
	}
	configureTestDefaultStandaloneDiscoveryOCI()

	assert := assert.New(t)

	added := populateDefaultStandaloneDiscovery(cfg)
	assert.Equal(true, added)
	assert.Equal(len(cfg.ClientOptions.CLI.DiscoverySources), 2)
	assert.Equal(cfg.ClientOptions.CLI.DiscoverySources[0].OCI.Name, DefaultStandaloneDiscoveryName)
	assert.Equal(cfg.ClientOptions.CLI.DiscoverySources[0].OCI.Image, "fake.image.repo/package/standalone-plugins:v1.0.0")
	assert.Equal(cfg.ClientOptions.CLI.DiscoverySources[1].OCI.Name, "additional-discovery")
	assert.Equal(cfg.ClientOptions.CLI.DiscoverySources[1].OCI.Image, "additional-discovery/path:v1.0.0")
}

func TestConfigPopulateDefaultStandaloneDiscoveryLocal(t *testing.T) {
	cfg := &configapi.ClientConfig{
		ClientOptions: &configapi.ClientOptions{
			CLI: &configapi.CLIOptions{
				DiscoverySources: []configapi.PluginDiscovery{},
			},
		},
	}

	configureTestDefaultStandaloneDiscoveryLocal()

	assert := assert.New(t)

	added := populateDefaultStandaloneDiscovery(cfg)
	assert.Equal(true, added)
	assert.Equal(len(cfg.ClientOptions.CLI.DiscoverySources), 1)
	assert.Equal(cfg.ClientOptions.CLI.DiscoverySources[0].Local.Name, DefaultStandaloneDiscoveryNameLocal)
	assert.Equal(cfg.ClientOptions.CLI.DiscoverySources[0].Local.Path, "local/path")
}

func TestConfigPopulateDefaultStandaloneDiscoveryEnvVariables(t *testing.T) {
	cfg := &configapi.ClientConfig{
		ClientOptions: &configapi.ClientOptions{
			CLI: &configapi.CLIOptions{
				DiscoverySources: []configapi.PluginDiscovery{},
			},
		},
	}

	configureTestDefaultStandaloneDiscoveryOCI()

	os.Setenv(constants.ConfigVariableCustomImageRepository, "env.fake.image.repo")
	os.Setenv(constants.ConfigVariableDefaultStandaloneDiscoveryImagePath, "package/env/standalone-plugins")
	os.Setenv(constants.ConfigVariableDefaultStandaloneDiscoveryImageTag, "v2.0.0")

	assert := assert.New(t)

	added := populateDefaultStandaloneDiscovery(cfg)
	assert.Equal(true, added)
	assert.Equal(len(cfg.ClientOptions.CLI.DiscoverySources), 1)
	assert.Equal(cfg.ClientOptions.CLI.DiscoverySources[0].OCI.Name, DefaultStandaloneDiscoveryName)
	assert.Equal(cfg.ClientOptions.CLI.DiscoverySources[0].OCI.Image, "env.fake.image.repo/package/env/standalone-plugins:v2.0.0")
}

func configureTestDefaultStandaloneDiscoveryOCI() {
	DefaultStandaloneDiscoveryType = "oci"
	DefaultStandaloneDiscoveryRepository = "fake.image.repo"
	DefaultStandaloneDiscoveryImagePath = "package/standalone-plugins"
	DefaultStandaloneDiscoveryImageTag = "v1.0.0"
}

func configureTestDefaultStandaloneDiscoveryLocal() {
	DefaultStandaloneDiscoveryType = "local"
	DefaultStandaloneDiscoveryLocalPath = "local/path"
}
