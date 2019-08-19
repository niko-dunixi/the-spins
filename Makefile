.phony: install copyright-infringement clean

install:
	echo "Installer"

copyright-infringement:
	./can-i-haz-muzic.sh
	go generate ./data

clean:
	rm -rfv ./bin
	rm -rfv ./data/assets/*