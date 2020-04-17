// Copyright (c) 2018, Google, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package config

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/spinnaker/spin/config/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Config is the CLI configuration kept in '~/.spin/config'.
type Config struct {
	Gate struct {
		Endpoint string `yaml:"endpoint"`
	} `yaml:"gate"`
	Auth *auth.Config `yaml:"auth"`
}

// Resolves tries to figure out where the config resides.
func Resolve(flags *pflag.FlagSet) (string, error) {
	flag, err := flags.GetString("config")
	if err != nil {
		return "", err
	}
	if flag != "" {
		return flag, nil
	}
	home, err := findHome()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".spin", "config"), nil
}

func findHome() (string, error) {
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir, nil
	}
	// Could not get current user.
	home := os.Getenv("HOME")
	if home == "" {
		return "", status.Errorf(codes.NotFound, "current user not found")
	}
	return home, nil
}
