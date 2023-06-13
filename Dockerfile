# Debian'nun resmi imajını kullan
FROM debian:latest

# Gerekli paketlerin güncellenmesi
RUN apt-get update && apt-get upgrade -y

# MongoDB'nin eklenmesi için gerekli olan public key'in import edilmesi
RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 9DA31620334BD75D9DCB49F368818C72E52529D4

# MongoDB reposunun eklenmesi
RUN echo "deb http://repo.mongodb.org/apt/debian stretch/mongodb-org/6.0 main" | tee /etc/apt/sources.list.d/mongodb-org-6.0.list

# MongoDB'nin kurulumu
RUN apt-get update && apt-get install -y mongodb-org

# MongoDB'nin default veritabanı dizini
RUN mkdir -p /data/db

# MongoDB portu
EXPOSE 27017

# MongoDB'nin başlatılması
CMD ["mongod"]
