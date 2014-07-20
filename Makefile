TARGETS := build install test

$(TARGETS):
	go $@
