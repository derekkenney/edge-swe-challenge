FROM mongo:latest
COPY init_messages.json /init_messages.json
CMD mongoimport --host mongo --db events --collection messages --drop --file /init_messages.json --jsonArray
