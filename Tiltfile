#!/usr/bin/env python

update_settings(max_parallel_updates=6)

local_resource(
    "pants", cmd="pants --loop fmt lint check package ::", allow_parallel=True
)

local_resource("dev", cmd="cd && make dev", allow_parallel=True)