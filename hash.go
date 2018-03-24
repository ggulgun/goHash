package main

import (
    "fmt"
    "log"
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "os"
    "io"
    "encoding/hex"
)


type MaliciousFile struct {
    md5File string
    sha1File string
    sha2File string
    filePath string
}


func calc_md5(filePath string) string {
    f, err := os.Open(filePath)
    if err != nil {
       log.Fatal(err)
    }

    h := md5.New()
    if _, err := io.Copy(h, f); err != nil {
       log.Fatal(err)
    }
    return hex.EncodeToString(h.Sum(nil))
}


func calc_sha1(filePath string) string {
    f, err := os.Open(filePath)
    if err != nil {
       log.Fatal(err)
    }


    h := sha1.New()
    if _, err := io.Copy(h, f); err != nil {
       log.Fatal(err)
    }
    return hex.EncodeToString(h.Sum(nil))
}

func calc_sha2(filePath string) string {
    f, err := os.Open(filePath)
    if err != nil {
       log.Fatal(err)
    }

    h := sha256.New()
    if _, err := io.Copy(h, f); err != nil {
       log.Fatal(err)
    }
    return hex.EncodeToString(h.Sum(nil))
}

func main() {
    filePath := os.Args[1]
    fmt.Println("Your File :",filePath)

    md5File := calc_md5(filePath)
    sha1File := calc_sha1(filePath)
    sha2File := calc_sha2(filePath)

    maliciousClass := MaliciousFile{md5File,sha1File,sha2File,filePath}
    fmt.Println("Md5 :",maliciousClass.md5File)
    fmt.Println("Sha128 :",maliciousClass.sha1File)
    fmt.Println("Sha256 :",maliciousClass.sha2File)


}
