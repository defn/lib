import json
import os

context = {"excludeStackIdFromLogicalIds": True, "allowSepCharsInLogicalIds": True}
os.environ.setdefault("CDKTF_CONTEXT_JSON", json.dumps(context))
