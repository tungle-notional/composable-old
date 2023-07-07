package keepers

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"

	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"

	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"

	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"

	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govtypesv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/cosmos/cosmos-sdk/x/group"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"

	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"

	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	
	"github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v7/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibcclient "github.com/cosmos/ibc-go/v7/modules/core/02-client"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibchost "github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"
	icq "github.com/strangelove-ventures/async-icq/v7"
	icqkeeper "github.com/strangelove-ventures/async-icq/v7/keeper"
	icqtypes "github.com/strangelove-ventures/async-icq/v7/types"

	custombankkeeper "github.com/notional-labs/centauri/v3/custom/bank/keeper"

	"github.com/strangelove-ventures/packet-forward-middleware/v7/router"
	routerkeeper "github.com/strangelove-ventures/packet-forward-middleware/v7/router/keeper"
	routertypes "github.com/strangelove-ventures/packet-forward-middleware/v7/router/types"

	alliancemodule "github.com/terra-money/alliance/x/alliance"
	alliancemoduletypes "github.com/terra-money/alliance/x/alliance/types"
	alliancemodulekeeper "github.com/terra-money/alliance/x/alliance/keeper"
	
	transfermiddleware "github.com/notional-labs/centauri/v3/x/transfermiddleware"
	transfermiddlewarekeeper "github.com/notional-labs/centauri/v3/x/transfermiddleware/keeper"
	transfermiddlewaretypes "github.com/notional-labs/centauri/v3/x/transfermiddleware/types"

	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"

	mintkeeper "github.com/notional-labs/centauri/v3/x/mint/keeper"
	minttypes "github.com/notional-labs/centauri/v3/x/mint/types"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	wasm08 "github.com/cosmos/ibc-go/v7/modules/light-clients/08-wasm/keeper"
	wasmtypes "github.com/cosmos/ibc-go/v7/modules/light-clients/08-wasm/types"
)

const (
	AccountAddressPrefix = "centauri"
	authorityAddress     = "centauri10556m38z4x6pqalr9rl5ytf3cff8q46nk85k9m"
)

type AppKeepers struct {
		// keys to access the substores
		keys    map[string]*storetypes.KVStoreKey
		tkeys   map[string]*storetypes.TransientStoreKey
		memKeys map[string]*storetypes.MemoryStoreKey
	
		// keepers
		AccountKeeper    authkeeper.AccountKeeper
		BankKeeper       custombankkeeper.Keeper
		CapabilityKeeper *capabilitykeeper.Keeper
		StakingKeeper    *stakingkeeper.Keeper
		SlashingKeeper   slashingkeeper.Keeper
		MintKeeper       mintkeeper.Keeper
		DistrKeeper      distrkeeper.Keeper
		GovKeeper        govkeeper.Keeper
		CrisisKeeper     *crisiskeeper.Keeper
		UpgradeKeeper    *upgradekeeper.Keeper
		ParamsKeeper     paramskeeper.Keeper
		IBCKeeper        *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
		EvidenceKeeper   evidencekeeper.Keeper
		TransferKeeper   ibctransferkeeper.Keeper
		ICQKeeper        icqkeeper.Keeper
		FeeGrantKeeper   feegrantkeeper.Keeper
		GroupKeeper      groupkeeper.Keeper
		Wasm08Keeper     wasm08.Keeper // TODO: use this name ?
		// make scoped keepers public for test purposes
		ScopedIBCKeeper       capabilitykeeper.ScopedKeeper
		ScopedTransferKeeper  capabilitykeeper.ScopedKeeper
		ConsensusParamsKeeper consensusparamkeeper.Keeper
		// this line is used by starport scaffolding # stargate/app/keeperDeclaration
		TransferMiddlewareKeeper transfermiddlewarekeeper.Keeper
		RouterKeeper             *routerkeeper.Keeper
		AllianceKeeper           alliancemodulekeeper.Keeper
}

// InitNormalKeepers initializes all 'normal' keepers.
func (appKeepers *AppKeepers) InitNormalKeepers(
	appCodec codec.Codec,
	cdc *codec.LegacyAmino,
	bApp *baseapp.BaseApp,
	maccPerms map[string][]string,
	invCheckPeriod uint,
	skipUpgradeHeights map[int64]bool,
	homePath string,
) {
	// add keepers
	appKeepers.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec, appKeepers.keys[authtypes.StoreKey], authtypes.ProtoBaseAccount, maccPerms, AccountAddressPrefix, authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	appKeepers.BankKeeper = custombankkeeper.NewBaseKeeper(
		appCodec, appKeepers.keys[banktypes.StoreKey], appKeepers.AccountKeeper, appKeepers.BlacklistedModuleAccountAddrs(maccPerms), authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	appKeepers.StakingKeeper = stakingkeeper.NewKeeper(
		appCodec, appKeepers.keys[stakingtypes.StoreKey], appKeepers.AccountKeeper, appKeepers.BankKeeper, authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	appKeepers.MintKeeper = mintkeeper.NewKeeper(
		appCodec, appKeepers.keys[minttypes.StoreKey], appKeepers.StakingKeeper,
		appKeepers.AccountKeeper, appKeepers.BankKeeper, authtypes.FeeCollectorName, authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	appKeepers.DistrKeeper = distrkeeper.NewKeeper(
		appCodec, appKeepers.keys[distrtypes.StoreKey], appKeepers.AccountKeeper, appKeepers.BankKeeper,
		appKeepers.StakingKeeper, authtypes.FeeCollectorName, authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	appKeepers.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec, cdc, appKeepers.keys[slashingtypes.StoreKey], appKeepers.StakingKeeper, authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	
	appKeepers.CrisisKeeper = crisiskeeper.NewKeeper(appCodec, appKeepers.keys[crisistypes.StoreKey],
		invCheckPeriod, appKeepers.BankKeeper, authtypes.FeeCollectorName, authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	groupConfig := group.DefaultConfig()
	/*
		Example of setting group params:
		groupConfig.MaxMetadataLen = 1000
	*/
	appKeepers.GroupKeeper = groupkeeper.NewKeeper(
		appKeepers.keys[group.StoreKey],
		appCodec,
		bApp.MsgServiceRouter(),
		appKeepers.AccountKeeper,
		groupConfig,
	)

	appKeepers.FeeGrantKeeper = feegrantkeeper.NewKeeper(appCodec, appKeepers.keys[feegrant.StoreKey], appKeepers.AccountKeeper)
	appKeepers.UpgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, appKeepers.keys[upgradetypes.StoreKey], appCodec, homePath, bApp, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	appKeepers.AllianceKeeper = alliancemodulekeeper.NewKeeper(
		appCodec,
		appKeepers.keys[alliancemoduletypes.StoreKey],
		appKeepers.GetSubspace(alliancemoduletypes.ModuleName),
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.StakingKeeper,
		appKeepers.DistrKeeper,
	)

	appKeepers.BankKeeper.RegisterKeepers(appKeepers.AllianceKeeper, appKeepers.StakingKeeper)
	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	appKeepers.StakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(appKeepers.DistrKeeper.Hooks(), appKeepers.SlashingKeeper.Hooks(), appKeepers.AllianceKeeper.StakingHooks()),
	)

	// ... other modules keepers

	// Create IBC Keeper
	appKeepers.IBCKeeper = ibckeeper.NewKeeper(
		appCodec, appKeepers.keys[ibchost.StoreKey], appKeepers.GetSubspace(ibchost.ModuleName), appKeepers.StakingKeeper, appKeepers.UpgradeKeeper, appKeepers.ScopedIBCKeeper,
	)

	appKeepers.Wasm08Keeper = wasm08.NewKeeper(appCodec, appKeepers.keys[wasmtypes.StoreKey], authorityAddress, homePath)
	// Create Transfer Keepers
	appKeepers.TransferMiddlewareKeeper = transfermiddlewarekeeper.NewKeeper(
		appKeepers.keys[transfermiddlewaretypes.StoreKey],
		appCodec,
		appKeepers.IBCKeeper.ChannelKeeper,
		&appKeepers.TransferKeeper,
		appKeepers.BankKeeper,
		authorityAddress,
	)

	appKeepers.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec, appKeepers.keys[ibctransfertypes.StoreKey],
		appKeepers.GetSubspace(ibctransfertypes.ModuleName),
		&appKeepers.TransferMiddlewareKeeper, // ICS4Wrapper
		appKeepers.IBCKeeper.ChannelKeeper,
		&appKeepers.IBCKeeper.PortKeeper,
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.ScopedTransferKeeper,
	)

	appKeepers.RouterKeeper = routerkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[routertypes.StoreKey],
		appKeepers.GetSubspace(routertypes.ModuleName),
		appKeepers.TransferKeeper,
		appKeepers.IBCKeeper.ChannelKeeper,
		&appKeepers.DistrKeeper,
		appKeepers.BankKeeper,
		appKeepers.TransferMiddlewareKeeper,
		appKeepers.IBCKeeper.ChannelKeeper,
	)

	transferIBCModule := transfer.NewIBCModule(appKeepers.TransferKeeper)
	
	scopedICQKeeper := appKeepers.CapabilityKeeper.ScopeToModule(icqtypes.ModuleName)

	appKeepers.ICQKeeper = icqkeeper.NewKeeper(
		appCodec, appKeepers.keys[icqtypes.StoreKey], appKeepers.GetSubspace(icqtypes.ModuleName),
		appKeepers.IBCKeeper.ChannelKeeper, appKeepers.IBCKeeper.ChannelKeeper, &appKeepers.IBCKeeper.PortKeeper,
		scopedICQKeeper, bApp,
	)

	icqIBCModule := icq.NewIBCModule(appKeepers.ICQKeeper)
	transfermiddlewareStack := transfermiddleware.NewIBCMiddleware(
		transferIBCModule,
		appKeepers.TransferMiddlewareKeeper,
	)

	ibcMiddlewareStack := router.NewIBCMiddleware(
		transfermiddlewareStack,
		appKeepers.RouterKeeper,
		0,
		routerkeeper.DefaultForwardTransferPacketTimeoutTimestamp,
		routerkeeper.DefaultRefundTransferPacketTimeoutTimestamp,
	)

	// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec, appKeepers.keys[evidencetypes.StoreKey], appKeepers.StakingKeeper, appKeepers.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	appKeepers.EvidenceKeeper = *evidenceKeeper

	// Register Gov (must be registered after stakeibc)
	govRouter := govtypesv1beta1.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govtypesv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(appKeepers.ParamsKeeper)).
		// AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(appKeepers.DistrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(appKeepers.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(appKeepers.IBCKeeper.ClientKeeper)).
		AddRoute(alliancemoduletypes.RouterKey, alliancemodule.NewAllianceProposalHandler(appKeepers.AllianceKeeper))

	govKeeper := *govkeeper.NewKeeper(
		appCodec, appKeepers.keys[govtypes.StoreKey], appKeepers.AccountKeeper, appKeepers.BankKeeper,
		appKeepers.StakingKeeper, bApp.MsgServiceRouter(), govtypes.DefaultConfig(), authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	govKeeper.SetLegacyRouter(govRouter)

	appKeepers.GovKeeper = *govKeeper.SetHooks(
		govtypes.NewMultiGovHooks(
		// register the governance hooks
		),
	)

	ibcRouter := porttypes.NewRouter()
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, ibcMiddlewareStack)
	ibcRouter.AddRoute(icqtypes.ModuleName, icqIBCModule)

	// this line is used by starport scaffolding # ibc/app/router
	appKeepers.IBCKeeper.SetRouter(ibcRouter)
}

// InitSpecialKeepers initiates special keepers (upgradekeeper, params keeper)
func (appKeepers *AppKeepers) InitSpecialKeepers(
	appCodec codec.Codec,
	cdc *codec.LegacyAmino,
	bApp *baseapp.BaseApp,
	invCheckPeriod uint,
	skipUpgradeHeights map[int64]bool,
	homePath string,
) {
	appKeepers.GenerateKeys()
	appKeepers.ParamsKeeper = appKeepers.initParamsKeeper(appCodec, cdc, appKeepers.keys[paramstypes.StoreKey], appKeepers.tkeys[paramstypes.TStoreKey])
	appKeepers.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, appKeepers.keys[capabilitytypes.StoreKey], appKeepers.memKeys[capabilitytypes.MemStoreKey])

	// set the BaseApp's parameter store
	appKeepers.ConsensusParamsKeeper = consensusparamkeeper.NewKeeper(appCodec, appKeepers.keys[consensusparamtypes.StoreKey], authtypes.NewModuleAddress(govtypes.ModuleName).String())
	bApp.SetParamStore(&appKeepers.ConsensusParamsKeeper)

	// grant capabilities for the ibc and ibc-transfer modules
	appKeepers.ScopedIBCKeeper = appKeepers.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	appKeepers.ScopedTransferKeeper = appKeepers.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)

	appKeepers.UpgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, appKeepers.keys[upgradetypes.StoreKey], appCodec, homePath, bApp, authtypes.NewModuleAddress(govtypes.ModuleName).String())
}

// initParamsKeeper init params keeper and its subspaces
func (appKeepers *AppKeepers) initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(routertypes.ModuleName).WithKeyTable(routertypes.ParamKeyTable()) // TODO:
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypesv1.ParamKeyTable())     //nolint:staticcheck
	paramsKeeper.Subspace(minttypes.ModuleName).WithKeyTable(minttypes.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(icqtypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(alliancemoduletypes.ModuleName)

	return paramsKeeper
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (appKeepers *AppKeepers) BlacklistedModuleAccountAddrs(maccPerms map[string][]string) map[string]bool {
	modAccAddrs := make(map[string]bool)
	// DO NOT REMOVE: StringMapKeys fixes non-deterministic map iteration
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}
	return modAccAddrs
}