package conf

type configDef struct {
	arg, conf, env, def string
}

var (
	confDefs = map[string]configDef{
		"address":         configDef{arg: "address", conf: "server.address", env: "ADDRESS", def: ":2018"},
		"shutdownTimeout": configDef{arg: "shutdownTimeout", conf: "system.timeout", env: "SHUTDOWN_TIMEOUT", def: "5"},
	}
)
