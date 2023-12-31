### Develop
dev-run-back:
	go run main.go

dev-run-front:
	npm start --prefix front

### Prod
prod-build:
	make prod-build-front
	$(call build_prod,linux,amd64,linux_amd64,file_server)
	$(call build_prod,windows,amd64,windows_amd64,file_server.exe)
	$(call build_prod,darwin,amd64,mac_amd64,file_server)

prod-build-front:
	# build front and copy to files to backend folder
	npm run build --prefix front
	cp -r front/public/assets/css src/public/assets/
	cp -v front/build/static/js/main.*.js src/public/assets/js/react.js

define build_prod
	# create dir for result files, copy files to this dir
	mkdir -p builds/$(3)/src/
	mkdir -p builds/$(3)/user/files/
	env GOOS=$(1) GOARCH=$(2) go build -o builds/$(3)/$(4) homeHelper
	cp -r src/public builds/$(3)/src/

	# create archive
	cd builds/ && zip -r $(3) ./* -x "*.zip" && cd -
	rm -r builds/$(3)
endef

