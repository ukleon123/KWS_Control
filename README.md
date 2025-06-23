# KWS Control

KWS Control

## system requirements

- Go >= 1.21
- libvirt
- Docker

## run

### 1. dependencies

#### Ubuntu/Debian
```sh
sudo apt install libvirt-dev pkg-config
```

#### CentOS/RHEL
```sh
sudo yum install libvirt-devel pkgconfig
```

#### I use Arch BTW
```sh
sudo pacman -S libvirt qemu dnsmasq openbsd-netcat
```

### 3. Go mod download
```sh
go mod download
```

## run

### method1: build & run
```sh
make run

make build

./kws

make clean
```

### method2: docker
```sh
docker build -t kws-control .

docker run -p 8081:8081 kws-control
```

## tree

```
KWS_Control/
|── api/
│   |── server/     # HTTP API server
│   |── workercont/ # worker control
|── config/         # config
|── vm/             # VM management
|── util/           # utils
|── main.go         # main.go
|── Dockerfile      # dockerfile
|── Makefile        # biuld script
|── go.mod          # go module setting
```