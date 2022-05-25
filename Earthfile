VERSION --shell-out-anywhere --use-chmod --use-host-command --earthly-version-arg --use-copy-link 0.6

IMPORT github.com/defn/cloud/lib:master AS lib

FROM lib+platform

warm:
    COPY pyproject.toml poetry.lock .
    RUN ~/bin/e poetry install

build:
    FROM +warm
    COPY src src
    RUN ~/bin/e poetry build
    SAVE ARTIFACT dist AS LOCAL dist

publish:
    FROM +warm
    COPY dist dist
    RUN --push --secret POETRY_PYPI_TOKEN_PYPI ~/bin/e poetry publish

