echo "creating docker image..."
VERSION=$(git rev-parse HEAD)

echo commmit SHA:
echo $VERSION

docker build -t employees-manager-api:$VERSION .