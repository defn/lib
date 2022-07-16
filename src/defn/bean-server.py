import asyncio
import logging

import grpc

import defn.dev.legumes.v1.bean_pb2 as bean_pb2
import defn.dev.legumes.v1.bean_pb2_grpc as bean_pb2_grpc


class GreeterService(bean_pb2_grpc.BeanStoreServiceServicer):
    async def GetBean(
        self, request: bean_pb2.Bean, context: grpc.aio.ServicerContext
    ) -> bean_pb2.Bean:
        return bean_pb2.Bean(url=request.url, sha256=request.sha256)


async def serve() -> None:
    server = grpc.aio.server()
    bean_pb2_grpc.add_BeanStoreServiceServicer_to_server(GreeterService(), server)
    listen_addr = "0.0.0.0:50051"
    server.add_insecure_port(listen_addr)
    logging.info("Starting server on %s", listen_addr)
    await server.start()
    await server.wait_for_termination()


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    asyncio.run(serve())
