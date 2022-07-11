from __future__ import print_function

import logging

import grpc

import defn.dev.legumes.v1.bean_pb2 as bean_pb2
import defn.dev.legumes.v1.bean_pb2_grpc as bean_pb2_grpc


def run():
    with grpc.insecure_channel("localhost:50051") as channel:
        stub = bean_pb2_grpc.BeanStoreServiceStub(channel)
        response: bean_pb2.Bean = stub.GetBean(
            bean_pb2.Bean(url="hello", sha256="world")
        )
    print("Bean client received: " + response.url + " " + response.sha256)


if __name__ == "__main__":
    logging.basicConfig()
    run()
