package go_zero_pprof

// A Config is a pprof config.
type Config struct {
	Host string `json:",optional"`
	Port int    `json:",default=6060"`
}
