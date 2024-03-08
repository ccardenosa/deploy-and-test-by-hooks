# Makefile

PYTHON_INTERPRETER=python3.12
PYTHON_VENV_PATH=.venv

all: install test

install: venv
	: # Activate venv and install somthing inside
	source $(PYTHON_VENV_PATH)/bin/activate && python3 -m pip install --upgrade pip && pip install -r requirements.txt
	source $(PYTHON_VENV_PATH)/bin/activate && (\
		ansible-galaxy collection install -r requirements.yml -p ./playbooks/collections \
	)

devel: install git-setup
	: # Other commands here
	source $(PYTHON_VENV_PATH)/bin/activate && (\
		ansible-galaxy collection install -r requirements-devel.yml -p ./playbooks/collections \
	)

git-setup:
	: # Install pre-commit hooks
	source $(PYTHON_VENV_PATH)/bin/activate && (\
		pushd . ; \
		cd ../.. ; \
		test -d .git -a ! -f .git/hooks/pre-commit && pre-commit install || echo \
		popd ; \
	)

venv:
	: # Create $(PYTHON_VENV_PATH) if it doesn't exist
	test -f $(PYTHON_VENV_PATH)/bin/activate || $(PYTHON_INTERPRETER) -m venv $(PYTHON_VENV_PATH)

run:
	: # Run your app here, e.g
	source $(PYTHON_VENV_PATH)/bin/activate && pip -V

	: # Exec multiple commands
	source $(PYTHON_VENV_PATH)/bin/activate && (\
		python3 -c 'import sys; print(sys.prefix)'; \
		echo ; \
		pip3 -V ;\
		echo ; \
		ansible-playbook --version ; \
		echo ; \
		echo "ansible-builder Version: " ; \
		ansible-builder --version \
	)

test:
	: # Running Testing...
	: # NOTE: Uncomment the following lines when you have tests to run
	: source $(PYTHON_VENV_PATH)/bin/activate && (\
	:	export PYTHONPATH=${PWD}:${PATH}; \
	:	pytest tests/ \
	:)

build-image:
	: # Building container image:
	podman build -f Containerfile -t quay.io/ccardenosa/virtual-hub:stream9 .

push-image:
	: # Upload container image:
	podman push quay.io/ccardenosa/virtual-hub:stream9

clean:
	rm -rf $(PYTHON_VENV_PATH) playbooks/collections context
	find . -type f -iname '*.pyc' -delete
