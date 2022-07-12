from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar, Iterable, Mapping, Optional, Text, Union

DESCRIPTOR: _descriptor.FileDescriptor
MARKDOWN: StringKind
PLAIN: StringKind

class ApplyResourceChange(_message.Message):
    __slots__ = []
    class Request(_message.Message):
        __slots__ = ["config", "planned_private", "planned_state", "prior_state", "provider_meta", "type_name"]
        CONFIG_FIELD_NUMBER: ClassVar[int]
        PLANNED_PRIVATE_FIELD_NUMBER: ClassVar[int]
        PLANNED_STATE_FIELD_NUMBER: ClassVar[int]
        PRIOR_STATE_FIELD_NUMBER: ClassVar[int]
        PROVIDER_META_FIELD_NUMBER: ClassVar[int]
        TYPE_NAME_FIELD_NUMBER: ClassVar[int]
        config: DynamicValue
        planned_private: bytes
        planned_state: DynamicValue
        prior_state: DynamicValue
        provider_meta: DynamicValue
        type_name: str
        def __init__(self, type_name: Optional[str] = ..., prior_state: Optional[Union[DynamicValue, Mapping]] = ..., planned_state: Optional[Union[DynamicValue, Mapping]] = ..., config: Optional[Union[DynamicValue, Mapping]] = ..., planned_private: Optional[bytes] = ..., provider_meta: Optional[Union[DynamicValue, Mapping]] = ...) -> None: ...
    class Response(_message.Message):
        __slots__ = ["diagnostics", "legacy_type_system", "new_state", "private"]
        DIAGNOSTICS_FIELD_NUMBER: ClassVar[int]
        LEGACY_TYPE_SYSTEM_FIELD_NUMBER: ClassVar[int]
        NEW_STATE_FIELD_NUMBER: ClassVar[int]
        PRIVATE_FIELD_NUMBER: ClassVar[int]
        diagnostics: _containers.RepeatedCompositeFieldContainer[Diagnostic]
        legacy_type_system: bool
        new_state: DynamicValue
        private: bytes
        def __init__(self, new_state: Optional[Union[DynamicValue, Mapping]] = ..., private: Optional[bytes] = ..., diagnostics: Optional[Iterable[Union[Diagnostic, Mapping]]] = ..., legacy_type_system: bool = ...) -> None: ...
    def __init__(self) -> None: ...

class AttributePath(_message.Message):
    __slots__ = ["steps"]
    class Step(_message.Message):
        __slots__ = ["attribute_name", "element_key_int", "element_key_string"]
        ATTRIBUTE_NAME_FIELD_NUMBER: ClassVar[int]
        ELEMENT_KEY_INT_FIELD_NUMBER: ClassVar[int]
        ELEMENT_KEY_STRING_FIELD_NUMBER: ClassVar[int]
        attribute_name: str
        element_key_int: int
        element_key_string: str
        def __init__(self, attribute_name: Optional[str] = ..., element_key_string: Optional[str] = ..., element_key_int: Optional[int] = ...) -> None: ...
    STEPS_FIELD_NUMBER: ClassVar[int]
    steps: _containers.RepeatedCompositeFieldContainer[AttributePath.Step]
    def __init__(self, steps: Optional[Iterable[Union[AttributePath.Step, Mapping]]] = ...) -> None: ...

class Configure(_message.Message):
    __slots__ = []
    class Request(_message.Message):
        __slots__ = ["config", "terraform_version"]
        CONFIG_FIELD_NUMBER: ClassVar[int]
        TERRAFORM_VERSION_FIELD_NUMBER: ClassVar[int]
        config: DynamicValue
        terraform_version: str
        def __init__(self, terraform_version: Optional[str] = ..., config: Optional[Union[DynamicValue, Mapping]] = ...) -> None: ...
    class Response(_message.Message):
        __slots__ = ["diagnostics"]
        DIAGNOSTICS_FIELD_NUMBER: ClassVar[int]
        diagnostics: _containers.RepeatedCompositeFieldContainer[Diagnostic]
        def __init__(self, diagnostics: Optional[Iterable[Union[Diagnostic, Mapping]]] = ...) -> None: ...
    def __init__(self) -> None: ...

class Diagnostic(_message.Message):
    __slots__ = ["attribute", "detail", "severity", "summary"]
    class Severity(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = []
    ATTRIBUTE_FIELD_NUMBER: ClassVar[int]
    DETAIL_FIELD_NUMBER: ClassVar[int]
    ERROR: Diagnostic.Severity
    INVALID: Diagnostic.Severity
    SEVERITY_FIELD_NUMBER: ClassVar[int]
    SUMMARY_FIELD_NUMBER: ClassVar[int]
    WARNING: Diagnostic.Severity
    attribute: AttributePath
    detail: str
    severity: Diagnostic.Severity
    summary: str
    def __init__(self, severity: Optional[Union[Diagnostic.Severity, str]] = ..., summary: Optional[str] = ..., detail: Optional[str] = ..., attribute: Optional[Union[AttributePath, Mapping]] = ...) -> None: ...

class DynamicValue(_message.Message):
    __slots__ = ["json", "msgpack"]
    JSON_FIELD_NUMBER: ClassVar[int]
    MSGPACK_FIELD_NUMBER: ClassVar[int]
    json: bytes
    msgpack: bytes
    def __init__(self, msgpack: Optional[bytes] = ..., json: Optional[bytes] = ...) -> None: ...

class GetProviderSchema(_message.Message):
    __slots__ = []
    class Request(_message.Message):
        __slots__ = []
        def __init__(self) -> None: ...
    class Response(_message.Message):
        __slots__ = ["data_source_schemas", "diagnostics", "provider", "provider_meta", "resource_schemas"]
        class DataSourceSchemasEntry(_message.Message):
            __slots__ = ["key", "value"]
            KEY_FIELD_NUMBER: ClassVar[int]
            VALUE_FIELD_NUMBER: ClassVar[int]
            key: str
            value: Schema
            def __init__(self, key: Optional[str] = ..., value: Optional[Union[Schema, Mapping]] = ...) -> None: ...
        class ResourceSchemasEntry(_message.Message):
            __slots__ = ["key", "value"]
            KEY_FIELD_NUMBER: ClassVar[int]
            VALUE_FIELD_NUMBER: ClassVar[int]
            key: str
            value: Schema
            def __init__(self, key: Optional[str] = ..., value: Optional[Union[Schema, Mapping]] = ...) -> None: ...
        DATA_SOURCE_SCHEMAS_FIELD_NUMBER: ClassVar[int]
        DIAGNOSTICS_FIELD_NUMBER: ClassVar[int]
        PROVIDER_FIELD_NUMBER: ClassVar[int]
        PROVIDER_META_FIELD_NUMBER: ClassVar[int]
        RESOURCE_SCHEMAS_FIELD_NUMBER: ClassVar[int]
        data_source_schemas: _containers.MessageMap[str, Schema]
        diagnostics: _containers.RepeatedCompositeFieldContainer[Diagnostic]
        provider: Schema
        provider_meta: Schema
        resource_schemas: _containers.MessageMap[str, Schema]
        def __init__(self, provider: Optional[Union[Schema, Mapping]] = ..., resource_schemas: Optional[Mapping[str, Schema]] = ..., data_source_schemas: Optional[Mapping[str, Schema]] = ..., diagnostics: Optional[Iterable[Union[Diagnostic, Mapping]]] = ..., provider_meta: Optional[Union[Schema, Mapping]] = ...) -> None: ...
    def __init__(self) -> None: ...

class GetProvisionerSchema(_message.Message):
    __slots__ = []
    class Request(_message.Message):
        __slots__ = []
        def __init__(self) -> None: ...
    class Response(_message.Message):
        __slots__ = ["diagnostics", "provisioner"]
        DIAGNOSTICS_FIELD_NUMBER: ClassVar[int]
        PROVISIONER_FIELD_NUMBER: ClassVar[int]
        diagnostics: _containers.RepeatedCompositeFieldContainer[Diagnostic]
        provisioner: Schema
        def __init__(self, provisioner: Optional[Union[Schema, Mapping]] = ..., diagnostics: Optional[Iterable[Union[Diagnostic, Mapping]]] = ...) -> None: ...
    def __init__(self) -> None: ...

class ImportResourceState(_message.Message):
    __slots__ = []
    class ImportedResource(_message.Message):
        __slots__ = ["private", "state", "type_name"]
        PRIVATE_FIELD_NUMBER: ClassVar[int]
        STATE_FIELD_NUMBER: ClassVar[int]
        TYPE_NAME_FIELD_NUMBER: ClassVar[int]
        private: bytes
        state: DynamicValue
        type_name: str
        def __init__(self, type_name: Optional[str] = ..., state: Optional[Union[DynamicValue, Mapping]] = ..., private: Optional[bytes] = ...) -> None: ...
    class Request(_message.Message):
        __slots__ = ["id", "type_name"]
        ID_FIELD_NUMBER: ClassVar[int]
        TYPE_NAME_FIELD_NUMBER: ClassVar[int]
        id: str
        type_name: str
        def __init__(self, type_name: Optional[str] = ..., id: Optional[str] = ...) -> None: ...
    class Response(_message.Message):
        __slots__ = ["diagnostics", "imported_resources"]
        DIAGNOSTICS_FIELD_NUMBER: ClassVar[int]
        IMPORTED_RESOURCES_FIELD_NUMBER: ClassVar[int]
        diagnostics: _containers.RepeatedCompositeFieldContainer[Diagnostic]
        imported_resources: _containers.RepeatedCompositeFieldContainer[ImportResourceState.ImportedResource]
        def __init__(self, imported_resources: Optional[Iterable[Union[ImportResourceState.ImportedResource, Mapping]]] = ..., diagnostics: Optional[Iterable[Union[Diagnostic, Mapping]]] = ...) -> None: ...
    def __init__(self) -> None: ...

class PlanResourceChange(_message.Message):
    __slots__ = []
    class Request(_message.Message):
        __slots__ = ["config", "prior_private", "prior_state", "proposed_new_state", "provider_meta", "type_name"]
        CONFIG_FIELD_NUMBER: ClassVar[int]
        PRIOR_PRIVATE_FIELD_NUMBER: ClassVar[int]
        PRIOR_STATE_FIELD_NUMBER: ClassVar[int]
        PROPOSED_NEW_STATE_FIELD_NUMBER: ClassVar[int]
        PROVIDER_META_FIELD_NUMBER: ClassVar[int]
        TYPE_NAME_FIELD_NUMBER: ClassVar[int]
        config: DynamicValue
        prior_private: bytes
        prior_state: DynamicValue
        proposed_new_state: DynamicValue
        provider_meta: DynamicValue
        type_name: str
        def __init__(self, type_name: Optional[str] = ..., prior_state: Optional[Union[DynamicValue, Mapping]] = ..., proposed_new_state: Optional[Union[DynamicValue, Mapping]] = ..., config: Optional[Union[DynamicValue, Mapping]] = ..., prior_private: Optional[bytes] = ..., provider_meta: Optional[Union[DynamicValue, Mapping]] = ...) -> None: ...
    class Response(_message.Message):
        __slots__ = ["diagnostics", "legacy_type_system", "planned_private", "planned_state", "requires_replace"]
        DIAGNOSTICS_FIELD_NUMBER: ClassVar[int]
        LEGACY_TYPE_SYSTEM_FIELD_NUMBER: ClassVar[int]
        PLANNED_PRIVATE_FIELD_NUMBER: ClassVar[int]
        PLANNED_STATE_FIELD_NUMBER: ClassVar[int]
        REQUIRES_REPLACE_FIELD_NUMBER: ClassVar[int]
        diagnostics: _containers.RepeatedCompositeFieldContainer[Diagnostic]
        legacy_type_system: bool
        planned_private: bytes
        planned_state: DynamicValue
        requires_replace: _containers.RepeatedCompositeFieldContainer[AttributePath]
        def __init__(self, planned_state: Optional[Union[DynamicValue, Mapping]] = ..., requires_replace: Optional[Iterable[Union[AttributePath, Mapping]]] = ..., planned_private: Optional[bytes] = ..., diagnostics: Optional[Iterable[Union[Diagnostic, Mapping]]] = ..., legacy_type_system: bool = ...) -> None: ...
    def __init__(self) -> None: ...

class PrepareProviderConfig(_message.Message):
    __slots__ = []
    class Request(_message.Message):
        __slots__ = ["config"]
        CONFIG_FIELD_NUMBER: ClassVar[int]
        config: DynamicValue
        def __init__(self, config: Optional[Union[DynamicValue, Mapping]] = ...) -> None: ...
    class Response(_message.Message):
        __slots__ = ["diagnostics", "prepared_config"]
        DIAGNOSTICS_FIELD_NUMBER: ClassVar[int]
        PREPARED_CONFIG_FIELD_NUMBER: ClassVar[int]
        diagnostics: _containers.RepeatedCompositeFieldContainer[Diagnostic]
        prepared_config: DynamicValue
        def __init__(self, prepared_config: Optional[Union[DynamicValue, Mapping]] = ..., diagnostics: Optional[Iterable[Union[Diagnostic, Mapping]]] = ...) -> None: ...
    def __init__(self) -> None: ...

class ProvisionResource(_message.Message):
    __slots__ = []
    class Request(_message.Message):
        __slots__ = ["config", "connection"]
        CONFIG_FIELD_NUMBER: ClassVar[int]
        CONNECTION_FIELD_NUMBER: ClassVar[int]
        config: DynamicValue
        connection: DynamicValue
        def __init__(self, config: Optional[Union[DynamicValue, Mapping]] = ..., connection: Optional[Union[DynamicValue, Mapping]] = ...) -> None: ...
    class Response(_message.Message):
        __slots__ = ["diagnostics", "output"]
        DIAGNOSTICS_FIELD_NUMBER: ClassVar[int]
        OUTPUT_FIELD_NUMBER: ClassVar[int]
        diagnostics: _containers.RepeatedCompositeFieldContainer[Diagnostic]
        output: str
        def __init__(self, output: Optional[str] = ..., diagnostics: Optional[Iterable[Union[Diagnostic, Mapping]]] = ...) -> None: ...
    def __init__(self) -> None: ...

class RawState(_message.Message):
    __slots__ = ["flatmap", "json"]
    class FlatmapEntry(_message.Message):
        __slots__ = ["key", "value"]
        KEY_FIELD_NUMBER: ClassVar[int]
        VALUE_FIELD_NUMBER: ClassVar[int]
        key: str
        value: str
        def __init__(self, key: Optional[str] = ..., value: Optional[str] = ...) -> None: ...
    FLATMAP_FIELD_NUMBER: ClassVar[int]
    JSON_FIELD_NUMBER: ClassVar[int]
    flatmap: _containers.ScalarMap[str, str]
    json: bytes
    def __init__(self, json: Optional[bytes] = ..., flatmap: Optional[Mapping[str, str]] = ...) -> None: ...

class ReadDataSource(_message.Message):
    __slots__ = []
    class Request(_message.Message):
        __slots__ = ["config", "provider_meta", "type_name"]
        CONFIG_FIELD_NUMBER: ClassVar[int]
        PROVIDER_META_FIELD_NUMBER: ClassVar[int]
        TYPE_NAME_FIELD_NUMBER: ClassVar[int]
        config: DynamicValue
        provider_meta: DynamicValue
        type_name: str
        def __init__(self, type_name: Optional[str] = ..., config: Optional[Union[DynamicValue, Mapping]] = ..., provider_meta: Optional[Union[DynamicValue, Mapping]] = ...) -> None: ...
    class Response(_message.Message):
        __slots__ = ["diagnostics", "state"]
        DIAGNOSTICS_FIELD_NUMBER: ClassVar[int]
        STATE_FIELD_NUMBER: ClassVar[int]
        diagnostics: _containers.RepeatedCompositeFieldContainer[Diagnostic]
        state: DynamicValue
        def __init__(self, state: Optional[Union[DynamicValue, Mapping]] = ..., diagnostics: Optional[Iterable[Union[Diagnostic, Mapping]]] = ...) -> None: ...
    def __init__(self) -> None: ...

class ReadResource(_message.Message):
    __slots__ = []
    class Request(_message.Message):
        __slots__ = ["current_state", "private", "provider_meta", "type_name"]
        CURRENT_STATE_FIELD_NUMBER: ClassVar[int]
        PRIVATE_FIELD_NUMBER: ClassVar[int]
        PROVIDER_META_FIELD_NUMBER: ClassVar[int]
        TYPE_NAME_FIELD_NUMBER: ClassVar[int]
        current_state: DynamicValue
        private: bytes
        provider_meta: DynamicValue
        type_name: str
        def __init__(self, type_name: Optional[str] = ..., current_state: Optional[Union[DynamicValue, Mapping]] = ..., private: Optional[bytes] = ..., provider_meta: Optional[Union[DynamicValue, Mapping]] = ...) -> None: ...
    class Response(_message.Message):
        __slots__ = ["diagnostics", "new_state", "private"]
        DIAGNOSTICS_FIELD_NUMBER: ClassVar[int]
        NEW_STATE_FIELD_NUMBER: ClassVar[int]
        PRIVATE_FIELD_NUMBER: ClassVar[int]
        diagnostics: _containers.RepeatedCompositeFieldContainer[Diagnostic]
        new_state: DynamicValue
        private: bytes
        def __init__(self, new_state: Optional[Union[DynamicValue, Mapping]] = ..., diagnostics: Optional[Iterable[Union[Diagnostic, Mapping]]] = ..., private: Optional[bytes] = ...) -> None: ...
    def __init__(self) -> None: ...

class Schema(_message.Message):
    __slots__ = ["block", "version"]
    class Attribute(_message.Message):
        __slots__ = ["computed", "deprecated", "description", "description_kind", "name", "optional", "required", "sensitive", "type"]
        COMPUTED_FIELD_NUMBER: ClassVar[int]
        DEPRECATED_FIELD_NUMBER: ClassVar[int]
        DESCRIPTION_FIELD_NUMBER: ClassVar[int]
        DESCRIPTION_KIND_FIELD_NUMBER: ClassVar[int]
        NAME_FIELD_NUMBER: ClassVar[int]
        OPTIONAL_FIELD_NUMBER: ClassVar[int]
        REQUIRED_FIELD_NUMBER: ClassVar[int]
        SENSITIVE_FIELD_NUMBER: ClassVar[int]
        TYPE_FIELD_NUMBER: ClassVar[int]
        computed: bool
        deprecated: bool
        description: str
        description_kind: StringKind
        name: str
        optional: bool
        required: bool
        sensitive: bool
        type: bytes
        def __init__(self, name: Optional[str] = ..., type: Optional[bytes] = ..., description: Optional[str] = ..., required: bool = ..., optional: bool = ..., computed: bool = ..., sensitive: bool = ..., description_kind: Optional[Union[StringKind, str]] = ..., deprecated: bool = ...) -> None: ...
    class Block(_message.Message):
        __slots__ = ["attributes", "block_types", "deprecated", "description", "description_kind", "version"]
        ATTRIBUTES_FIELD_NUMBER: ClassVar[int]
        BLOCK_TYPES_FIELD_NUMBER: ClassVar[int]
        DEPRECATED_FIELD_NUMBER: ClassVar[int]
        DESCRIPTION_FIELD_NUMBER: ClassVar[int]
        DESCRIPTION_KIND_FIELD_NUMBER: ClassVar[int]
        VERSION_FIELD_NUMBER: ClassVar[int]
        attributes: _containers.RepeatedCompositeFieldContainer[Schema.Attribute]
        block_types: _containers.RepeatedCompositeFieldContainer[Schema.NestedBlock]
        deprecated: bool
        description: str
        description_kind: StringKind
        version: int
        def __init__(self, version: Optional[int] = ..., attributes: Optional[Iterable[Union[Schema.Attribute, Mapping]]] = ..., block_types: Optional[Iterable[Union[Schema.NestedBlock, Mapping]]] = ..., description: Optional[str] = ..., description_kind: Optional[Union[StringKind, str]] = ..., deprecated: bool = ...) -> None: ...
    class NestedBlock(_message.Message):
        __slots__ = ["block", "max_items", "min_items", "nesting", "type_name"]
        class NestingMode(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
            __slots__ = []
        BLOCK_FIELD_NUMBER: ClassVar[int]
        GROUP: Schema.NestedBlock.NestingMode
        INVALID: Schema.NestedBlock.NestingMode
        LIST: Schema.NestedBlock.NestingMode
        MAP: Schema.NestedBlock.NestingMode
        MAX_ITEMS_FIELD_NUMBER: ClassVar[int]
        MIN_ITEMS_FIELD_NUMBER: ClassVar[int]
        NESTING_FIELD_NUMBER: ClassVar[int]
        SET: Schema.NestedBlock.NestingMode
        SINGLE: Schema.NestedBlock.NestingMode
        TYPE_NAME_FIELD_NUMBER: ClassVar[int]
        block: Schema.Block
        max_items: int
        min_items: int
        nesting: Schema.NestedBlock.NestingMode
        type_name: str
        def __init__(self, type_name: Optional[str] = ..., block: Optional[Union[Schema.Block, Mapping]] = ..., nesting: Optional[Union[Schema.NestedBlock.NestingMode, str]] = ..., min_items: Optional[int] = ..., max_items: Optional[int] = ...) -> None: ...
    BLOCK_FIELD_NUMBER: ClassVar[int]
    VERSION_FIELD_NUMBER: ClassVar[int]
    block: Schema.Block
    version: int
    def __init__(self, version: Optional[int] = ..., block: Optional[Union[Schema.Block, Mapping]] = ...) -> None: ...

class Stop(_message.Message):
    __slots__ = []
    class Request(_message.Message):
        __slots__ = []
        def __init__(self) -> None: ...
    class Response(_message.Message):
        __slots__ = ["Error"]
        ERROR_FIELD_NUMBER: ClassVar[int]
        Error: str
        def __init__(self, Error: Optional[str] = ...) -> None: ...
    def __init__(self) -> None: ...

class UpgradeResourceState(_message.Message):
    __slots__ = []
    class Request(_message.Message):
        __slots__ = ["raw_state", "type_name", "version"]
        RAW_STATE_FIELD_NUMBER: ClassVar[int]
        TYPE_NAME_FIELD_NUMBER: ClassVar[int]
        VERSION_FIELD_NUMBER: ClassVar[int]
        raw_state: RawState
        type_name: str
        version: int
        def __init__(self, type_name: Optional[str] = ..., version: Optional[int] = ..., raw_state: Optional[Union[RawState, Mapping]] = ...) -> None: ...
    class Response(_message.Message):
        __slots__ = ["diagnostics", "upgraded_state"]
        DIAGNOSTICS_FIELD_NUMBER: ClassVar[int]
        UPGRADED_STATE_FIELD_NUMBER: ClassVar[int]
        diagnostics: _containers.RepeatedCompositeFieldContainer[Diagnostic]
        upgraded_state: DynamicValue
        def __init__(self, upgraded_state: Optional[Union[DynamicValue, Mapping]] = ..., diagnostics: Optional[Iterable[Union[Diagnostic, Mapping]]] = ...) -> None: ...
    def __init__(self) -> None: ...

class ValidateDataSourceConfig(_message.Message):
    __slots__ = []
    class Request(_message.Message):
        __slots__ = ["config", "type_name"]
        CONFIG_FIELD_NUMBER: ClassVar[int]
        TYPE_NAME_FIELD_NUMBER: ClassVar[int]
        config: DynamicValue
        type_name: str
        def __init__(self, type_name: Optional[str] = ..., config: Optional[Union[DynamicValue, Mapping]] = ...) -> None: ...
    class Response(_message.Message):
        __slots__ = ["diagnostics"]
        DIAGNOSTICS_FIELD_NUMBER: ClassVar[int]
        diagnostics: _containers.RepeatedCompositeFieldContainer[Diagnostic]
        def __init__(self, diagnostics: Optional[Iterable[Union[Diagnostic, Mapping]]] = ...) -> None: ...
    def __init__(self) -> None: ...

class ValidateProvisionerConfig(_message.Message):
    __slots__ = []
    class Request(_message.Message):
        __slots__ = ["config"]
        CONFIG_FIELD_NUMBER: ClassVar[int]
        config: DynamicValue
        def __init__(self, config: Optional[Union[DynamicValue, Mapping]] = ...) -> None: ...
    class Response(_message.Message):
        __slots__ = ["diagnostics"]
        DIAGNOSTICS_FIELD_NUMBER: ClassVar[int]
        diagnostics: _containers.RepeatedCompositeFieldContainer[Diagnostic]
        def __init__(self, diagnostics: Optional[Iterable[Union[Diagnostic, Mapping]]] = ...) -> None: ...
    def __init__(self) -> None: ...

class ValidateResourceTypeConfig(_message.Message):
    __slots__ = []
    class Request(_message.Message):
        __slots__ = ["config", "type_name"]
        CONFIG_FIELD_NUMBER: ClassVar[int]
        TYPE_NAME_FIELD_NUMBER: ClassVar[int]
        config: DynamicValue
        type_name: str
        def __init__(self, type_name: Optional[str] = ..., config: Optional[Union[DynamicValue, Mapping]] = ...) -> None: ...
    class Response(_message.Message):
        __slots__ = ["diagnostics"]
        DIAGNOSTICS_FIELD_NUMBER: ClassVar[int]
        diagnostics: _containers.RepeatedCompositeFieldContainer[Diagnostic]
        def __init__(self, diagnostics: Optional[Iterable[Union[Diagnostic, Mapping]]] = ...) -> None: ...
    def __init__(self) -> None: ...

class StringKind(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
