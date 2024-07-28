package gotorrentparser

import "github.com/zeebo/bencode"

type Dictionary struct {
	Path []string `bencode:"path"`

	PathUtf8 []string `bencode:"path.utf-8"`

	Length int64 `bencode:"length"`
}

type Info struct {
	// In the single file case, the name key is the name of a file, in the muliple file case, it's the name of a directory.
	Name string `bencode:"name"`

	NameUtf8 string `bencode:"name.utf-8"`

	PieceLength int64 `bencode:"piece length"`

	Pieces []byte `bencode:"pieces"`

	// There is also a key length or a key files, but not both or neither. If length is present then the download represents a single file, otherwise it represents a set of files which go in a directory structure.

	// single file context
	Length int64 `bencode:"length"`

	// multi file context. type Dictionary struct
	Files []Dictionary `bencode:"files"`
}

// Known as .torrent file.
// Data structure:
// https://fileformats.fandom.com/wiki/Torrent_file.
// http://bittorrent.org/beps/bep_0003.html
// Announce List: http://bittorrent.org/beps/bep_0012.html
type MetaInfo struct {
	// required
	Announce string `bencode:"announce"`

	// extension
	AnnounceList [][]string `bencode:"announce-list"`

	// optional
	Comment string `bencode:"comment"`

	// optional
	CreatedBy string `bencode:"created by"`

	// optional
	CreatedAt int64 `bencode:"creation date"`

	// required. type Info struct
	RawInfo bencode.RawMessage `bencode:"info"`
}
