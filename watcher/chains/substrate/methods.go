package substrate

type Method string

const NuProxy = "nuproxy"

var (
	BaseMethod       Method = ""
	RegisterWatcher  Method = NuProxy + ".register_watcher"
	UpdateStakerInfo Method = NuProxy + ".update_staker_infos_and_mint"
)
