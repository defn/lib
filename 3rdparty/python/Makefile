all: venv default.lock

venv: requirements.txt
	python -mvenv .venv
	.venv/bin/pip install -r requirements.txt

requirements.txt: packages.txt
	pip-compile $< -o $@

default.lock: requirements.txt
	../../pants generate-lock-files
