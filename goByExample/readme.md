# Go by example

<https://gobyexample.com/>

## Installing Golang on Windows WSL2

<https://medium.com/@benzbraunstein/how-to-install-and-setup-golang-development-under-wsl-2-4b8ca7720374>

By default `wget` downloads files in the current working directory where it is run.

``` bash
wget https://dl.google.com/go/go1.17.1.linux-amd64.tar.gz
sudo tar -xvf go1.17.1.linux-amd64.tar.gz
sudo mv go /usr/local
```

replace go version with appropriate version when downloading
after above commands have been invoked in .bashrc add the following lines

``` bash
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```
