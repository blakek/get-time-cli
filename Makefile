GO ?= go
NAME ?= get-time
PREFIX ?= /usr/local/bin

.PHONY: all
all: get-time

# Remove built files
.PHONY: clean
clean:
	$(RM) -r ./dist

# Moves the built binary to the installation prefix
.PHONY: install
install: $(NAME)
	mv -i `pwd`/dist/$(NAME) $(PREFIX)/$(NAME)

# Builds the executable
$(NAME):
	$(GO) build -o dist/$(NAME) main.go
	@echo "Successfully built!"

# Alternate to install (for local development). Symlinks from this directory to
# the installation prefix.
.PHONY: symlink
symlink: $(NAME)
	ln -is `pwd`/dist/$(NAME) $(PREFIX)/$(NAME)
