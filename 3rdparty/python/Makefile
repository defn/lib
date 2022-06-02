requirements.txt: packages.txt
	pip-compile $< -o $@
	python -mvenv .venv
	.venv/bin/pip install -r $@
	../../pants generate-lock-files
