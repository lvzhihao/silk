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
	"net"
	"os"
	"os/signal"

	"go.uber.org/zap"

	"github.com/jinzhu/gorm"
	pbs "github.com/lvzhihao/silk/pbs"
	"github.com/lvzhihao/silk/servers"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	logger *zap.Logger
)

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "grpc",
	Long:  `grpc`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := gorm.Open("mysql", viper.GetString("mysql_dns"))
		if err != nil {
			logger.Fatal("failed connect db", zap.Error(err))
		}
		defer db.Close()
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return viper.GetString("table_prefix") + "_" + defaultTableName
		}
		lis, err := net.Listen("tcp", viper.GetString("grpc_host"))
		if err != nil {
			logger.Fatal("failed to listen", zap.Error(err))
		}
		s := grpc.NewServer()
		pbs.RegisterSilkerServer(s, &servers.Server{
			Logger: logger,
			DB:     db,
		})
		// Register reflection service on gRPC server.
		reflection.Register(s)
		go func() {
			if err := s.Serve(lis); err != nil {
				logger.Fatal("failed to serve", zap.Error(err))
			}
		}()
		// Stop
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)
		<-stop
		s.GracefulStop()
		logger.Info("Server gracefully stopped")
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
	if os.Getenv("DEBUG") == "true" {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}
	logger.Sync()
}
