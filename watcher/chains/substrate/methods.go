package substrate

type Method string

const NuProxy = "NulinkNuproxy"

const Watchers = "Watchers"

var (
	RegisterWatcher Method = NuProxy + ".register_watcher"
	UpdateStakeInfo Method = NuProxy + ".update_staker_infos_and_mint"
)
