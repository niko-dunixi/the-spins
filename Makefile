.default: install
.phony: install generate clean

export bin_directory = bin
export asset_directory = data/assets
export song_url = https://www.youtube.com/watch?v=z8cgNLGnnK4
export song_file = ${bin_directory}/nsp-you-spin-me-cover.mp3
export loop_file = ${asset_directory}/spin-loop.mp3

install:
	go install .

generate: $(loop_file)
	go generate ./data

$(loop_file): $(song_file)
	mkdir -p "${asset_directory}"
	ffmpeg -i "${song_file}" -ss 00:01:13.30 -to 00:01:30.38 -c copy "${loop_file}"

$(song_file): .git/hooks/pre-commit
	mkdir -p "${bin_directory}"
	youtube-dl "${song_url}" --extract-audio --audio-format mp3 --exec "mv {} ${song_file}"

.git/hooks/pre-commit:
	cp ./pre-commit .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit

clean:
	git clean -xdf