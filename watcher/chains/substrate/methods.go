package substrate

type Method string

const Pallets = "pallets"

var (
	BaseMethod      Method = ""
	RegisterWatcher Method = Pallets + ".register_watcher"
)
