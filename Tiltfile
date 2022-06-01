#!/usr/bin/env python

update_settings(max_parallel_updates=6)

local_resource(
    "pants", cmd="sudo apt install -y python3-pip && p --loop fmt lint check package ::", allow_parallel=True
)

local_resource("dev", cmd="cd && make dev", allow_parallel=True)

local_resource("pre-commit", cmd="cd && pre-commit run --all", allow_parallel=True)
