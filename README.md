# rp 

An implementation of Discord's rich presence for Linux, MacOS and Windows in Go. Fork of [fawni/rp](https://github.com/fawni/rp). Modified for music player rich presence builds.

## Installation

Install `https://github.com/Victiniiiii/rp`:

```sh
$ go get -u https://github.com/Victiniiiii/rp
```

## Usage

Create a new client

```go
c, err := rp.NewClient("DISCORD_APP_ID")
if err != nil {
	panic(err)
}
```

Set the rich presence activity

```go
if err := c.SetActivity(&rpc.Activity{
	State:      "Hey!",
	Details:    "Running on rp.go!",
	LargeImage: "largeimageid",
	LargeText:  "This is the large image",
	SmallImage: "smallimageid",
	SmallText:  "And this is the small image",
}) err != nil {
	panic(err)
}
```

More details in the [example](example/main.go)

## License

[ISC](LICENSE)
