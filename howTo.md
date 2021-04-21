# GENERAL INFORMATION I gathered when started dabbling in GO

Based on golang, linode and digital ocean linux/ubuntu install instructions

Sources:
<https://www.linode.com/docs/development/go/install-go-on-ubuntu/>
<https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-18-04#conclusion>
<https://golang.org/doc/install/source>

at time -v 1.13.1 installed
note: for security download in ~/ dir first, once integrity of download
confirmed moved to /usr/local

## step 1: download

``` bash
curl -O https://dl.google.com/go/go1.13.1.linux-amd64.tar.gz

# linode points to storage.googleapis, but golang official website <a.../> has
# href pointing to dl/google which digital ocean and used, which imo is a better
# approach.
```

## step 2: verify integrity

``` bash
sha256sum go1.13.1.linux-amd64.tar.gz

# sample output:
94f874037b82ea5353f4061e543681a0e79657f787437974214629af8407d124  go1.13.1.linux-amd64.tar.gz

94f874037b82ea5353f4061e543681a0e79657f787437974214629af8407d124 from golang download page, match == True
```

## step 3: extract iff integrity verified

``` bash
tar -xvf go1.13.1.linux-amd64.tar.gz
# from digital ocean:
# The x flag tells tar to extract, v tells it we want verbose output (a listing 
# of the files being extracted), and f tells it we’ll specify a filename:
```

## step 4: change owner to root(ie adjust permissions), move to /usr/local

``` bash
sudo chown -R root:root ./go
sudo mv go /usr/local
```

<https://en.wikipedia.org/wiki/Chown>

## step 5: path variable

for user in ~/.profile add new PATH for golang

``` bash
## from experience don't use export, type out GOPATH in profile...
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

I prefer '/go' over 'work' naming convention from dig.O
