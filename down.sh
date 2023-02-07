docker-compose down
docker rmi $(docker images | grep "douyin" | awk '{print $3}')
cd ../
docker system prune
sudo rm -rf douyin/
