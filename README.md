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
## Disclaimer
This software is provided "as is", without warranty of any kind, express or implied, including but not limited to the warranties of merchantability, fitness for a particular purpose, title and non-infringement. In no event shall the authors or copyright holders be liable for any claim, damages or other liability, whether in an action of contract, tort, or otherwise arising form, out of, or in connection with the software or the use or other dealings in the software.
### Security Notice
While Sigillum is designed simplify encryption and decryption of payloads, no software can guarantee absolute protection. Users are responsible for implementing and maintaining comprehensive security measures to protect their systems and data from cyber threats. The authors do not warrant that the software will integrate error-free with all other software running on users's systems.
### Compliance and Legal Use
Users are responsible for ensuring that their use of this software complies with all applicable local, national and international laws and regulations. The authors disclaim any liability for unauthorized or unlawful use of this software.