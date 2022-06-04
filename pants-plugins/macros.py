# type: ignore

def defn_sources(**kwargs):
    kwargs['resolve'] = 'defn'
    python_sources(**kwargs)
