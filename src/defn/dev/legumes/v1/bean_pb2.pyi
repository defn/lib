from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar, Optional, Text

DESCRIPTOR: _descriptor.FileDescriptor

class Bean(_message.Message):
    __slots__ = ["sha256", "url"]
    SHA256_FIELD_NUMBER: ClassVar[int]
    URL_FIELD_NUMBER: ClassVar[int]
    sha256: str
    url: str
    def __init__(self, url: Optional[str] = ..., sha256: Optional[str] = ...) -> None: ...
