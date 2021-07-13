GO111MODULE=on
SONG_URL := https://www.youtube.com/watch?v=9W6AN_eQeZo
# SONG_FILE := $(shell mktemp -d)/halogen-u-got-that.mp3
# SONG_START_TIME = 00:01:13.30
SONG_START_TIME = 00:00:32.90
SONG_END_TIME = 00:01:50.00

.phony: install
install:
	go install .

.phony: run
run:
	go run .

build-cache/full-audio.mp3:
	@mkdir -p build-cache
	youtube-dl "${SONG_URL}" --extract-audio --audio-format mp3 --exec "mv {} $@"

assets/spin-loop.mp3: build-cache/full-audio.mp3
	ffmpeg -i "$<" -ss $(SONG_START_TIME) -to $(SONG_END_TIME) -c copy "$@"

.phony: test
test: build-cache/full-audio.mp3
	[ ! -f ./assets/spin-loop.mp3 ] || rm ./assets/spin-loop.mp3
	$(MAKE) assets/spin-loop.mp3 run

clean:
	git clean -xdf