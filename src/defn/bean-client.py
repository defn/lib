import logging
import os

import grpc
from google.protobuf.json_format import Parse, ParseDict

import defn.dev.legumes.v1.bean_pb2 as bean_pb2
import defn.dev.legumes.v1.bean_pb2_grpc as bean_pb2_grpc


def run():
    with grpc.insecure_channel(
        os.environ.get("server", "kourier-internal-x-kourier-system-x-vc1.vc1.svc:80"),
        options=[
            (
                "grpc.default_authority",
                os.environ.get("authority", "hello.demo.svc.cluster.local"),
            )
        ],
    ) as channel:
        stub = bean_pb2_grpc.BeanStoreServiceStub(channel)

        response: bean_pb2.Bean = stub.GetBean(
            bean_pb2.Bean(
                url=os.environ.get("url", "cool"),
                sha256=os.environ.get("sha256", "beans"),
            )
        )
        print(f"Bean client received: {response.url}, {response.sha256}")


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    run()
