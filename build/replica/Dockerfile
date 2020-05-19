FROM mongo

WORKDIR /conf

COPY replicaSet.js .
COPY setup.sh .

RUN chmod +x /conf/setup.sh

CMD ["./setup.sh"]