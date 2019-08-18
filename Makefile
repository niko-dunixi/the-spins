.phony: install clean

install: spin-loop.mp3
	echo "Installer"

spin-loop.mp3:
	./can-i-haz-muzic.sh

clean:
	rm -rfv ./bin