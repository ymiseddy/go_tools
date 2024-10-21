# Argon2Idhash Wrapper

## What is it?

Easy to use wrapper around the Argon2id hashing algorithm to allow easy
hashing and verification. Salts are automatically generated and encoded
into the hash. Optionally, a global pepper can be used along with the
salt for further security.

## Basic Usage

Generate the hash bytes for a string:

```go
package main

import (
    "fmt"
    "github.com/ymiseddy/go_tools/idhash"
)

func main() {
    password := "Secret password"
    hashed := idhash.DefaultArgon2Id.GenerateHashBytes(password)

    fmt.Printf("Hashed: %x\n", hashed)

    verified := idhash.DefaultArgon2Id.Verify([]byte(password), hashed)
    if verified {
        fmt.Println("Passwords match.")
    } else {
        fmt.Println("Passwords do not match.")
    }
}
```

## Base64 encoded hash

Using convenience functions automatically convert to and from base64:

```go
    password := "Secret password"
    hashed, err := idhash.DefaultArgon2Id.GenerateHashBase64(password)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Hashed: %s\n", hashed)

    verified, err := idhash.DefaultArgon2Id.VerifyBase64(password, hashed)
    if err != nil {
        panic(err)
    }

    if verified {
        fmt.Println("Passwords match.")
    } else {
        fmt.Println("Passwords do not match.")
    }
```

## Using Pepper

An example of using a pepper when generating the hash by creating a copy of the
default argon2id struct:

```go
    password := "Secret password"
    hashgenerator := idhash.DefaultArgon2Id
    hashgenerator.Pepper = []byte{23, 21, 24, 17, 19, 21, 22, 27}

    hashed, err := hashgenerator.GenerateHashBase64(password)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Hashed: %s\n", hashed)

    verified, err := hashgenerator.VerifyBase64(password, hashed)
    if err != nil {
        panic(err)
    }

    if verified {
        fmt.Println("Passwords match.")
    } else {
        fmt.Println("Passwords do not match.")
    }
```
