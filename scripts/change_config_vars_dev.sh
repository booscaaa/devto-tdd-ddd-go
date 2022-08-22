cd $1

cp config.example.json config.json
sed -i -e 's/$DB_USER/'postgres'/g' config.json
sed -i -e 's/$DB_PASS/'postgres'/g' config.json
sed -i -e 's/$DB_HOST/'postgres'/g' config.json
sed -i -e 's/$DB_PORT/'5432'/g' config.json
sed -i -e 's/$DB_NAME/'tdd_ddd_go'/g' config.json
sed -i -e 's/$SERVER_PORT/'3000'/g' config.json
sed -i -e 's/$SSL_MODE/'disable'/g' config.json
sed -i -e 's/$SERVER_HOST/'localhost:3000'/g' config.json
sed -i -e 's/$SERVER_SCHEME/'http'/g' config.json