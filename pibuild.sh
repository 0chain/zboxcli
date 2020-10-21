wget https://golang.org/dl/go1.15.3.linux-arm64.tar.gz
sudo tar -C /usr/local -xzf go1.15.3.linux-arm64.tar.gz
rm go1.15.3.linux-arm64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
sudo apt install -y build-essential git nano
cd ~
mkdir .zcn
mkdir go
cd go
git clone https://github.com/herumi/mcl.git
git clone https://github.com/herumi/bls.git
git clone https://github.com/0chain/gosdk.git
git clone https://github.com/0chain/zboxcli.git
git clone https://github.com/0chain/zwalletcli.git
cd mcl
sudo make MCL_USE_GMP=0
cd ..
cd bls
sudo make MCL_USE_GMP=0
cd ..
cd gosdk
sed -i 's+github.com/herumi/bls-go-binary+//github.com/herumi/bls-go-binary+g' go.mod
sudo make install
cd ..
cd zboxcli
sed -i 's+// replace+replace+g' go.mod
make install
cp zbox ~/.zcn
cp network/one.yaml ~/.zcn/config.yaml
cd ..
cd zwalletcli
sed -i 's+// replace+replace+g' go.mod
make install
cp zwallet ~/.zcn
cd ..
cd ..
cd .zcn
./zbox version
./zwallet version
