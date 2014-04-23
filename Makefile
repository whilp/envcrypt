.PHONY: install test clean

install:
	python setup.py install

test:
	python setup.py test

clean:
	rm -rf build dist *egg-info *.pyc
