git pull
go build
cd ./simple-auth-docker
cp ../asr-auth ./
cp ../conf.ini ./
cp ../registry.html ./
docker build -t "simple-auth:$1" .
rm simple-auth
rm conf.ini
rm registry.html
