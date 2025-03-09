vine_execution_var = VINE_EXECUTE
#Do not change the vine_roo_dir
vine_root_dir = .vine
vine_rood_cmd = vine
vine_binay = vine_aly.exe

vine_path = $(HOME)/$(vine_root_dir)

.PHONY: help
help:
    @echo 'Usage:'

.PHONY: build
build:
	go build -o bin/$(vine_binay) main.go

.PHONY: clean
clean:
	# TODO rm vine_path probably
	# Preserve alias config and scripts
	rm -f $(vine_path)/$(vine_binay)
	rm -f bin/$(vine_binay)
	rm -rf $(vine_path)/.config
	rm -rf $(vine_path)/.auth

.PHONY: install
install: clean build
	mkdir -p $(vine_path)
	cp bin/$(vine_binay) $(vine_path)

	mkdir -p $(vine_path)/.config $(vine_path)/.auth
	cp config/vine_exposed_config.json $(vine_path)/.config/
	cp config/vine_run_auth.json $(vine_path)/.auth/

	env $(vine_execution_var)=make $(vine_path)/$(vine_binay) install

.PHONY: update
update:
	./scripts/update-refs.sh

	cp $(vine_path)/.auth/vine_run_auth.json ./config/

	# env $(vine_execution_var)=make go run main.go update
	# $(vine_path)/$(vine_binay) update
	# cp $(HOME)/.bashrc ./config/template_files/
