package substrate

type Method string

const NuProxy = "NulinkNuproxy"

var (
	BaseMethod      Method = ""
	RegisterWatcher Method = NuProxy + ".register_watcher"
	UpdateStakeInfo Method = NuProxy + ".update_staker_infos_and_mint"
)
