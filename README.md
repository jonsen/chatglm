# ChatGLM

[![Go Reference](https://pkg.go.dev/badge/github.com/jonsen/chatglm/vulndb.svg)][1]

Go binding for [ChatGLM.cpp][2].

## Usage

### Get the Code

```bash
git clone --recursive https://github.com/jonsen/chatglm.git && cd chatglm
```

Or:

```bash
git clone https://github.com/jonsen/chatglm.git && cd chatglm
git submodule update --init --recursive
```

### Quantize Model

Transform ChatGLM-6B into 4-bit quantized GGML format:

```bash
make convert
```

For other transformations, see [Quantize Model][3].

### Build & Test

Build the ChatGLM.cpp libraries:

```bash
make build
```

Run the tests:

```bash
go test -v -race ./...
```

Run the demo:

```bash
go run demo/main.go -m [your model file] -i
```

## License

[MIT](LICENSE)

[1]: https://pkg.go.dev/github.com/jonsen/chatglm
[2]: https://github.com/li-plus/chatglm.cpp
[3]: https://github.com/li-plus/chatglm.cpp#getting-started
