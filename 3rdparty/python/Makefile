all: venv default.lock

venv: requirements.txt
	/usr/bin/python3.8 -mvenv .venv
	.venv/bin/pip install -r requirements.txt

requirements.txt: packages.txt
	pip-compile $< -o $@

.venv/bin/activate: venv

default.lock: requirements.txt .venv/bin/activate
	source .venv/bin/activate && ../../pants generate-lockfiles
