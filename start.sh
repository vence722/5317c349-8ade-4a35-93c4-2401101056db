docker build -t order-api .
docker-compose up -d
echo "Start sleep for 60s, in order to wait for MySQL finish starting"
sleep 60
echo "Wake up. Now will init the database and the table"
docker cp ./init_table.sql order-mysql:/init_table.sql
docker exec order-mysql "mysql -u root -ppassword < /init_table.sql"
echo "Done."