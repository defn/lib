requirements.txt: packages.txt
	pip-compile $< -o $@

default.lock: requirements.txt
	../../pants generate-lock-files

venv:
	python -mvenv .venv
	.venv/bin/pip install -r requirements.txt
