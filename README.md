# crypt

A simple CLI tool for Encrypt/Decrypt files and upload/download to/from AWS S3. Using 256-bit AES-GCM.

## Build

```
$ go build crypt.go cryptobox.go s3.go
```

## Configure AWS S3
Create a AWS S3 bucket and configure `config/awsS3Conf.json`. bucket name and region is required.

### Make It Runnable from the Command-Line
After created an executable file, `crypt`. Copy `crypt` to `/usr/local/bin`.

```
$ cp crypt /usr/local/bin
```

## Usage

Encrypt & Decrypt a file
```
$ ./crypt enc -f=mysecret

$ ./crypt dec -f=mysecretEnc -key=3ejnUzswPQm9tiZ47EKTCQoGK4h03uK7heutnhYI14Q=
```

Encrypt file and Upload to AWS S3 (automatically generated an encryption key)

```
$ ./crypt push mysecret
```

Download and Decrypt file

```
$ ./crypt pull -key vWWgPEgRIOgWyTVRs2tzDYKqHWAHa6hSnX+C+N3i4jg= myfolder/mysecret
```

List objects on S3 bucket
```
$ ./crypt list
```
