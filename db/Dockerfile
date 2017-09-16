FROM postgres:9.3
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD docker
ENV POSTGRES_DB clientserver
ADD db.sql /docker-entrypoint-initdb.d/
