/* Copyright (c) Edgeless Systems GmbH

   This program is free software; you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation; version 2 of the License.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program; if not, write to the Free Software
   Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1335  USA */

package core

import (
	"os"
)

// Config is an EDB config.
type Config struct {
	DataPath           string `json:",omitempty"`
	DatabaseAddress    string `json:",omitempty"`
	APIAddress         string `json:",omitempty"`
	CertificateDNSName string `json:",omitempty"`
	Debug              bool   `json:",omitempty"`
	LogDir             string `json:",omitempty"`
}

// EnvDataPath is the name of the optional environment variable holding the data path for edb
const EnvDataPath = "EDG_EDB_DATA_PATH"

// EnvDatabaseAddress is the name of the optional environment variable holding the database address
const EnvDatabaseAddress = "EDG_EDB_DATABASE_ADDR"

// EnvAPIAddress is the name of the optional environment variable holding the API address
const EnvAPIAddress = "EDG_EDB_API_ADDR"

// EnvCertificateDNSName is the name of the optional environment variable holding the DNS Name used for the certificate generated by edb
const EnvCertificateDNSName = "EDG_EDB_CERT_DNS"

// EnvDebug is a flag to enable debug logging for edb
const EnvDebug = "EDG_EDB_DEBUG"

// EnvLogDir is the name of the optional environment variable holding the path for storing the log files
const EnvLogDir = "EDG_EDB_LOG_DIR"

// FillConfigFromEnvironment takes an existing config filled with defaults and replaces single values based on environment variables.
func FillConfigFromEnvironment(config Config) Config {
	envDataPath := os.Getenv(EnvDataPath)
	envDatabaseAddress := os.Getenv(EnvDatabaseAddress)
	envAPIAddress := os.Getenv(EnvAPIAddress)
	envCertificateDNSName := os.Getenv(EnvCertificateDNSName)
	envDebug := os.Getenv(EnvDebug)
	envLogDir := os.Getenv(EnvLogDir)

	if envDataPath != "" {
		config.DataPath = envDataPath
	}

	if envDatabaseAddress != "" {
		config.DatabaseAddress = envDatabaseAddress
	}

	if envAPIAddress != "" {
		config.APIAddress = envAPIAddress
	}

	if envCertificateDNSName != "" {
		config.CertificateDNSName = envCertificateDNSName
	}

	if envDebug != "" {
		config.Debug = true
	}

	if envLogDir != "" {
		config.LogDir = envLogDir
		config.Debug = true
	}

	return config
}
