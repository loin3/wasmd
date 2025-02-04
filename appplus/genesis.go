package appplus

import (
	wasmapp "github.com/line/wasmd/app"
)

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState() wasmapp.GenesisState {
	encodingConfig := wasmapp.MakeEncodingConfig()
	return ModuleBasics.DefaultGenesis(encodingConfig.Marshaler)
}
