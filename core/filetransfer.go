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
	// decoder to receive file sent from server
	dec := gob.NewDecoder(c.Conn)
	fs := &FileStruct{}
	err := dec.Decode(fs)		// receiving serialized file
	DisplayMsg(err)

	// since the recieved data is encrypted, it must be decrypted using the key
	DecryptedData := Decrypt(fs.FileContent, string(fs.EncryptionKey))

	// create a new file to store sent content
	fo, err := os.Create(fs.FileName)
	DisplayMsg(err)
	_, err = fo.Write(DecryptedData) // write decrypted data to file
	DisplayMsg(err)
	err = fo.Close()
	DisplayMsg(err)
	fmt.Println("successfully saved file")
}
