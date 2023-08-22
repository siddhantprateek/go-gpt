<h1 align="center">
    Go-GPT
</h1>
<p align="center">
    A simple way to communicate with OpenAI using your CLI.
</p>

![go-gpt-cli-thumb](https://user-images.githubusercontent.com/43869046/229339476-4bb3d0fa-f03c-4142-b825-7d7cc8ef3003.png)

---

## Application Features

### 👌 Features
- Currently `Go-GPT` uses OpenAI's `gpt-3.5-turbo` model. While building this project `GPT-4` and `gpt-4-32k` has already been released on 14 March, 2023.
- Support for `GPT-4` model (*coming soon*...)

### ⚡Performance
- `GPT-3` language model currently have limited knowledge of the world until September 2021, soon after upgrading to `GPT-4`, it will have better results.
- `GPT-3` language model can’t produce video, sound, or images currently.

## ⚙️ Installation

- Golang `1.20` or higher is required to build and run the project. You can find the installer on the official Golang download page [Go.dev](go.dev).

-  ```bash
    make go-gpt
    make build
   ```

---

```bash
$ go-gpt
         ____               ____ ____ _____ 
	/ ___|  ___        / ___|  _ \_   _|
	| |  _ / _ \ _____| |  _| |_) || |  
	| |_| | (_) |_____| |_| |  __/ | |  
	\____| \___/       \____|_|    |_|  

  Go-GPT is a simple way to communicate with OpenAI  
    using your CLI created by @siddhantprateek
```

## 🐳 To run Go-GPT using Docker 

To run Go-GPT with Docker, use the following command.
```bash
 docker run -d --name go-gpt -p 8090:8090 
```

To view logs from the Docker container, use the following command:
```bash
 docker logs -f go-gpt
```


## 🏗️ Features to work on

- [ ] Token Streaming Over CLI
- [ ] Implement a loading indicator during response generation.
- [ ] Add a `.go-gpt` history file to maintain chat histories and responses.

## 😉 Author

- [@siddhantprateek](https://github.com/siddhantprateek)

## 📝 License

[MIT](./LICENSE)
