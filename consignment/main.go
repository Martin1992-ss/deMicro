package main

import (
	"log"

	"github.com/Martin1992-ss/deMicro/common"
	"github.com/Martin1992-ss/deMicro/config"
	"github.com/Martin1992-ss/deMicro/consignment/handler"
	pb "github.com/Martin1992-ss/deMicro/interface-center/out/consignment"
	userPb "github.com/Martin1992-ss/deMicro/interface-center/out/user"
	vesselPb "github.com/Martin1992-ss/deMicro/interface-center/out/vessel"
	"github.com/micro/go-micro"
)

const service = "consignment"

func main() {
	session, err := common.CreateDBSession(service)
	if err != nil {
		log.Fatalf("create session error: %v\n", err)
	}
	// 创建于 MongoDB 的主会话，需在退出 main() 时候手动释放连接
	defer session.Close()

	// 直接调用自己写的公有的库获取 server，保持配置同步
	// common.AuthWrapper 为前置认证，采用JWT
	srv := common.GetMicroServer(service, micro.WrapHandler())
	// 作为 vessel-service 的客户端
	vClient := vesselPb.NewVesselServiceClient(config.GetServiceName("vessel"), srv.Client())
	uClient := userPb.NewUserServiceClient(config.GetServiceName("user"), srv.Client())
	bk := srv.Server().Options().Broker
	h := handler.GetHandler(session, vClient, uClient, bk)

	// 将 server 作为微服务的服务端
	pb.RegisterShippingServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
