package substrate

type Method string

const NuProxy = "nuproxy"

var (
	BaseMethod      Method = ""
	RegisterWatcher Method = NuProxy + ".register_watcher"
)
