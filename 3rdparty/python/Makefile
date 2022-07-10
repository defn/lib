all: default.lock

venv: requirements.txt
	python -mvenv .venv
	.venv/bin/pip install -r requirements.txt

requirements.txt: packages.txt
	pip-compile $< -o $@

.venv/bin/activate: venv

default.lock: requirements.txt .venv/bin/activate
	 . .venv/bin/activate && ../../pants generate-lockfiles
	 . .venv/bin/activate && ../../pants export ::

clean:
	rm -rf .venv
