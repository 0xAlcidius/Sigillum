# Sigillum
Sigillum is a tool created to obfuscate /rypt a payload. In Sigillum this process is called "Sealing". Sealing a payload with Sigillum will provide both the sealed payload as well as the desealing function in the programming language chosen by the user.
## How to use
Sigillum requires a payload the seal and a key to seal it with. If both are provided the standard seal algorithm will be RC4 and the output language will be C. Here follows an example of how to use Sigillum:
```shell
./main --payload "Sigillum is cool!" --key "secret"
```
A path to a file can also be used. Although with some files (like `.png` files or `.jpg` files), the seal can become quite substantial. Therefore it is recommended to use the `-o` switch to directly write the output to a file.
```shell
./main --payload "~/image.png" --key "secret" -o out.c
```
Which will provide a file as output. When a path is provided as the payload, Sigillum will also output the desealing function with the output of that function being a file.
## Installation
To install Sigillum to your machine simply clone the repository:
```shell
git clone <repo.git>
```
Make sure [Golang](https://golang.google.cn/dl/) is installed on your machine