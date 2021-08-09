package go_zero_pprof

import (
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/threading"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

var (
	once sync.Once
)

// StartAgent starts a pprof agent.
func StartAgent(c Config) {
	once.Do(func() {
		if len(c.Host) == 0 {
			return
		}

		threading.GoSafe(func() {
			addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
			logx.Infof("Starting pprof agent at %s", addr)
			if err := http.ListenAndServe(addr, nil); err != nil {
				logx.Error(err)
			}
		})
	})
}
