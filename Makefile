.phony: install copyright-infringement clean

install:
	echo "Installer"

copyright-infringement: .git/hooks/pre-commit
	./can-i-haz-muzic.sh
	go generate ./data

.git/hooks/pre-commit:
	echo "Installing git pre-commit hook"
	cp ./pre-commit .git/hooks/pre-commit

clean:
	rm -rfv ./bin
	rm -rfv ./data/assets/*