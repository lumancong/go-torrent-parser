package main

import (
	"crypto/sha1"
	"fmt"
	"os"
	"strings"
	"time"

	gotorrentparser "github.com/lumancong/go-torrent-parser"
	"github.com/zeebo/bencode"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("torrent file path missing")
		os.Exit(1)
	}
	fileBytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	var metaInfo gotorrentparser.MetaInfo
	err = bencode.DecodeBytes(fileBytes, &metaInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Announce: ", metaInfo.Announce)
	fmt.Println("AnnounceList: ")
	for _, v := range metaInfo.AnnounceList {
		fmt.Printf("\t%s\n", v)
	}
	fmt.Println("Comment: ", metaInfo.Comment)
	fmt.Println("CreatedBy: ", metaInfo.CreatedBy)
	fmt.Println("CreatedAt: ", time.Unix(metaInfo.CreatedAt, 0))
	fmt.Println("InfoHash: ", toSHA1(metaInfo.RawInfo))

	var info gotorrentparser.Info
	err = bencode.DecodeBytes(metaInfo.RawInfo, &info)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name: ", info.Name)
	fmt.Println("PieceLength: ", formatFileSize(float64(info.PieceLength), 1024.0))
	fmt.Println("TotalPieces: ", len(info.Pieces)/20)
	//fmt.Println("PIECESHASH", info.Pieces[:20])
	if info.Length == 0 { // multi-files case
		fmt.Println("===")
		for _, v := range info.Files {
			fmt.Printf("%s, %s\n", formatFileSize(float64(v.Length), 1024.0), strings.Join(v.Path, "/"))
		}
		fmt.Println("===")
	} else { // single file case
		fmt.Println("Length: ", formatFileSize(float64(info.Length), 1024.0))
	}
}

func toSHA1(data []byte) string {
	hash := sha1.New()
	hash.Write(data)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func formatFileSize(s float64, base float64) string {
	var sizes = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	unitsLimit := len(sizes)
	i := 0
	for s >= base && i < unitsLimit {
		s = s / base
		i++
	}

	f := "%.0f %s"
	if i > 1 {
		f = "%.2f %s"
	}

	return fmt.Sprintf(f, s, sizes[i])
}
