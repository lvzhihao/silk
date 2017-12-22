// Copyright © 2017 edwin <edwin.lzh@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"mime"
	"net/http"

	"go.uber.org/zap"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gws "github.com/lvzhihao/silk/pbs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// proxyCmd represents the proxy command
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "proxy",
	Long:  `proxy`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{
			grpc.WithInsecure(),
		}
		err := gws.RegisterSilkerHandlerFromEndpoint(ctx, mux, viper.GetString("grpc_host"), opts)
		if err != nil {
			logger.Fatal("failed connect grpc", zap.Error(err))
		}
		app := http.NewServeMux()
		app.Handle("/", mux)
		mime.AddExtensionType(".svg", "image/svg+xml")
		prefix := "/swagger-ui/"
		app.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir("swagger-ui/dist/"))))
		prefix = "/json/"
		app.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir("swagger-json/"))))
		err = http.ListenAndServe(viper.GetString("proxy_host"), app)
		if err != nil {
			logger.Fatal("failed listen", zap.Error(err))
		}
	},
}

func init() {
	rootCmd.AddCommand(proxyCmd)
}
