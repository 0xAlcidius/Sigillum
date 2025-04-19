# Sigillum
Sigillum is a tool created to obfuscate / enrypt a payload. In Sigillum this process is called "Sealing". Sealing a payload with Sigillum will provide both the sealed payload as well as the desealing function in the programming language chosen by the user. Sigillum can be used to simplify the process of developing software with encrypted payloads, ranging from images, to text, to binaries.
## How to use
Sigillum requires a payload the seal and a key to seal it with. If both are provided the standard seal algorithm will be RC4 and the output language will be C. Here follows an example of how to use Sigillum:
###### Windows
```shell
go run .\bin\main.go --payload "Sigillum is cool!" --key "secret"
```
###### Linux
```shell
go run ./bin/main.go --payload "Sigillum is cool!" --key "secret"
```
##### Choosing algorithm
To choose a specific algorithm for the sealing and desealing process. One can use the `--seal` switch:
###### Windows
```shell
go run .\bin\main.go --payload "Sigillum is cool!" --key "secret" --seal "AES"
```
###### Linux
```shell
go run ./bin/main.go --payload "Sigillum is cool!" --key "secret" --seal "AES"
```
##### Save output
A path to a file can also be used. Although with some files (like `.png` files or `.jpg` files), the seal can become quite substantial. Therefore it is recommended to use the `--output` switch to directly write the output to a file.
###### Windows
```shell
go run .\bin\main.go  --payload "Sigillum is cool!" --key "secret" --output out.c
```
###### Linux
```shell
go run ./bin/main.go  --payload "Sigillum is cool!" --key "secret" --output out.c
```
##### File as payload
If you want to use a file as a payload. There's a switch to also name the file once it has been desealed by the desealing algorithm. This switch is 
Which will provide a file as output. When a path is provided as the payload, Sigillum will also output the desealing function with the output of that function being a file. This can be achieved by using the `--filename` switch:
###### Windows
```shell
go run .\bin\main.go --payload "Sigillum is cool!" --key "secret" --filename "my_image.png"
```
###### Linux
```shell
go run ./bin/main.go --payload "Sigillum is cool!" --key "secret" --filename "my_image.png"
```
## Installation
To install Sigillum to your machine simply clone the repository:
```shell
git clone <repo.git>
```
Make sure [Golang](https://golang.google.cn/dl/) is installed on your machine
### Usage in code
To implement Sigillum in your own code, first of all make sure your code complies with the CPL license. That said, make sure you have Golang installed. You can use Sigillum as follows:
```go
func main() {
	seal, found := sigillum.Seal["RC4"]
	if !found {
		fmt.Printf("Seal function not found.")
		return
	}

	seal.ExecuteSeal([]byte("secret"), []byte("Sigillum is cool!"))
}
```
You can change RC4 also for XOR or AES for example.
## Disclaimer
This software is provided "as is", without warranty of any kind, express or implied, including but not limited to the warranties of merchantability, fitness for a particular purpose, title and non-infringement. In no event shall the authors or copyright holders be liable for any claim, damages or other liability, whether in an action of contract, tort, or otherwise arising form, out of, or in connection with the software or the use or other dealings in the software.
### Security Notice
While Sigillum is designed simplify encryption and decryption of payloads, no software can guarantee absolute protection. Users are responsible for implementing and maintaining comprehensive security measures to protect their systems and data from cyber threats. The authors do not warrant that the software will integrate error-free with all other software running on users's systems.
### Compliance and Legal Use
Users are responsible for ensuring that their use of this software complies with all applicable local, national and international laws and regulations. The authors disclaim any liability for unauthorized or unlawful use of this software.