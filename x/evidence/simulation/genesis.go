package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/evidence/exported"
	"github.com/cosmos/cosmos-sdk/x/evidence/types"
)

// Simulation parameter constants
const evidence = "evidence"

// GenEvidences returns an empty slice of evidences.
func GenEvidences(_ *rand.Rand, _ []simtypes.Account) []exported.Evidence {
	return []exported.Evidence{}
}

// RandomizedGenState generates a random GenesisState for evidence
func RandomizedGenState(simState *module.SimulationState) {
	var ev []exported.Evidence

	simState.AppParams.GetOrGenerate(evidence, &ev, simState.Rand, func(r *rand.Rand) { ev = GenEvidences(r, simState.Accounts) })

	evidenceGenesis := types.NewGenesisState(ev)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(evidenceGenesis)
}
