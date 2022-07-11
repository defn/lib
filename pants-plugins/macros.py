# type: ignore


def defn_sources(**kwargs):
    kwargs["interpreter_constraints"] = [">=3.10,<4"]
    kwargs["resolve"] = "defn"
    python_sources(**kwargs)


def defn_binary(**kwargs):
    kwargs["interpreter_constraints"] = [">=3.10,<4"]
    kwargs["resolve"] = "defn"
    pex_binary(**kwargs)


def pp_sources(**kwargs):
    kwargs["interpreter_constraints"] = [">=3.8,<4"]
    kwargs["resolve"] = "pants-plugins"
    python_sources(**kwargs)
