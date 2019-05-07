package main

import (
	"errors"
	"github.com/idata-shopee/gopcp"
	"github.com/idata-shopee/gopcp_rpc"
	"github.com/idata-shopee/gopcp_service"
	"github.com/idata-shopee/gopcp_stream"
	"github.com/idata-shopee/mc_service/mc"
	"os"
	"strconv"
)

func StartTcpServer(port int) error {
	var mm = mc.GetMemMap()

	// TODO file set and get
	return gopcp_service.StartTcpServer(port, func(streamServer *gopcp_stream.StreamServer) *gopcp.Sandbox {
		return gopcp.GetSandbox(map[string]*gopcp.BoxFunc{
			// (set, key string, value interface{})
			"set": gopcp.ToSandboxFun(func(args []interface{}, attachment interface{}, pcpServer *gopcp.PcpServer) (interface{}, error) {
				if len(args) != 2 {
					return nil, errors.New(`set method signature "(key string, value string)"`)
				} else if key, ok := args[0].(string); !ok {
					return nil, errors.New(`set method signature "(key string, value string)"`)
				} else {
					return nil, mm.Set(key, args[1])
				}
			}),

			// (get, key string)
			"get": gopcp.ToSandboxFun(func(args []interface{}, attachment interface{}, pcpServer *gopcp.PcpServer) (interface{}, error) {
				if len(args) != 1 {
					return nil, errors.New(`set method signature "(key string)"`)
				} else if key, ok := args[0].(string); !ok {
					return nil, errors.New(`set method signature "(key string, value string)"`)
				} else {
					return mm.Get(key)
				}
			}),
		})
	}, func() *gopcp_rpc.ConnectionEvent {
		return &gopcp_rpc.ConnectionEvent{
			// on close of connection
			func(err error) {
			},
			// new connection
			func(pcpConnectionHandler *gopcp_rpc.PCPConnectionHandler) {
			},
		}
	})
}

// read port from env
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("missing env PORT which must exists.")
	} else {
		if port_int, err := strconv.Atoi(port); err != nil {
			panic("Env PORT must be a number.")
		} else {
			if err = StartTcpServer(port_int); err != nil {
				panic(err)
			}
		}
	}
}
