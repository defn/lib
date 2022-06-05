from pants.engine.target import (
    COMMON_TARGET_FIELDS,
    Dependencies,
    SingleSourceField,
    StringField,
    Target,
)


class MehField(StringField):
    alias = "bleh"
    help = "What's making you feel bleh."


class MehTarget(Target):
    alias = "meh"
    core_fields = (*COMMON_TARGET_FIELDS, Dependencies, SingleSourceField, MehField)
    help = (
        "A meh target to describe your state of meh.\n\n"
    )


def target_types():
    return [MehTarget]
