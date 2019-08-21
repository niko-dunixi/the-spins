.default: install
.phony: install generate clean

export GO111MODULE=on
bin_directory:=bin
asset_directory:=data/assets
song_url:=https://www.youtube.com/watch?v=9W6AN_eQeZo
song_file:=$(bin_directory)/halogen-u-got-that.mp3
loop_file:=$(asset_directory)/spin-loop.mp3

install:
	go install .

.git/hooks/pre-commit:
	cp ./pre-commit .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit

$(song_file): .git/hooks/pre-commit
	mkdir -p "${bin_directory}"
	youtube-dl "${song_url}" --extract-audio --audio-format mp3 --exec "mv {} ${song_file}"
	# youtube-dl gives files incorrect timestamps causing us to always
	# download a fresh copy, which is time consuming. `touch` fixes
	# this.
	touch $@

$(loop_file): $(song_file)
	# Optionally cut a segment out for looping. This song
	# is better as it has less instrumentation, so copying
	# the whole file is alright
	mkdir -p "${asset_directory}"
	# ffmpeg -i "${song_file}" -ss 00:01:13.30 -to 00:01:30.38 -c copy "${loop_file}"
	cp ${song_file} ${loop_file}

generate: $(loop_file)
	go generate ./data

clean:
	git clean -xdf