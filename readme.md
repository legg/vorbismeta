# Add vorbis meta for flac

Command line utility for adding vorbis meta data to flac files
Using this [example]
Added the cmd flags for track, track number, artist...


## Build

$ git clone https://github.com/legg/vorbismeta.git
$ cd vorbismeta/
$ go get
$ go build


## Usage

$ vorbismeta -src="source.flac" -title="foo" -tracknumber="1" -artist="bar" -albumartist="coo" -album="dar" -year="2017" -dst="destination.flac"


## Public domain

The source code and any original content of this repository is hereby released into the [public domain].

[example]: https://play.golang.org/p/P3efeHuhHO
[public domain]: https://creativecommons.org/publicdomain/zero/1.0/