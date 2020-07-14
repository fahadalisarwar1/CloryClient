package core

import (
	"encoding/gob"
	"fmt"
	"os"
)

type FileStruct struct {
	FileName string
	FileSize int
	FileContent []byte
	EncryptionKey []byte
}

func (c *Client) DownloadFile(){
	dec := gob.NewDecoder(c.Conn)
	fs := &FileStruct{}
	err := dec.Decode(fs)

	DisplayMsg(err)
	fmt.Println("Before Decryption: ", fs.FileContent)
	DecryptedData := Decrypt(fs.FileContent, string(fs.EncryptionKey))
	fmt.Println("Decrypted data", DecryptedData)
	fo, err := os.Create(fs.FileName)
	DisplayMsg(err)
	_, err = fo.Write(DecryptedData)
	DisplayMsg(err)
	err = fo.Close()
	DisplayMsg(err)
}
