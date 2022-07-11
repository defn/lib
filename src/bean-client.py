from __future__ import print_function

import logging

import grpc
from google.protobuf.json_format import Parse, ParseDict

import defn.dev.legumes.v1.bean_pb2 as bean_pb2
import defn.dev.legumes.v1.bean_pb2_grpc as bean_pb2_grpc


a = {"url": "hello", "sha256": "dict"}
b = '{"url": "hello", "sha256": "string"}'


def run():
    with grpc.insecure_channel("localhost:50051") as channel:
        stub = bean_pb2_grpc.BeanStoreServiceStub(channel)

        response: bean_pb2.Bean = stub.GetBean(ParseDict(a, bean_pb2.Bean()))
        print("Bean client received: " + response.url + " " + response.sha256)

        response: bean_pb2.Bean = stub.GetBean(Parse(b, bean_pb2.Bean()))
        print("Bean client received: " + response.url + " " + response.sha256)


if __name__ == "__main__":
    logging.basicConfig()
    run()
