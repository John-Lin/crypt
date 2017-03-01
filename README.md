# Hermescrypt

A simple CLI tool for Encrypt/Decrypt files. Using 256-bit AES-GCM.

## Usage

Encrypt (automatically generated a encryption key)

```
$ ./hermescrypt enc --help
Usage of enc:
  -f string
    	Specify a filename.
```

Decrypt

```
$ ./hermescrypt dec --help
Usage of dec:
  -f string
    	Specify a filename.
  -key string
    	Specify a key for decrypting.
```
## Build

```
$ go build hermescrypt.go cryptobox.go
```

### Make It Runnable from the Command-Line
After created an executable file, `hermescrypt`. Copy `hermescrypt` to `/usr/local/bin`.

```
$ cp hermescrypt /usr/local/bin
```

## Example

```
$ ./hermescrypt enc -f=mysecret

$ ./hermescrypt dec -f=mysecretEnc -key=3ejnUzswPQm9tiZ47EKTCQoGK4h03uK7heutnhYI14Q=
```