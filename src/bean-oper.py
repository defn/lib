import logging

import kopf
from google.protobuf.json_format import ParseDict

import defn.dev.legumes.v1.bean_pb2 as bean_pb2


@kopf.on.create("beans")  # type: ignore
def create_fn(body, **kwargs):
    bean = ParseDict(body.spec, bean_pb2.Bean())
    logging.info(f"A handler is called with bean: url={bean.url}, sha256={bean.sha256}")
