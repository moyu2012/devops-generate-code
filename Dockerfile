FROM node:slim

MAINTAINER ShenQiang   <qiang.shen@cm-dt.com>

RUN echo "deb http://ppa.launchpad.net/webupd8team/java/ubuntu trusty main" >> /etc/apt/sources.list.d/oracje-java.list && \
    echo "deb-src http://ppa.launchpad.net/webupd8team/java/ubuntu trusty main" >> /etc/apt/sources.list.d/oracje-java.list && \
    apt-key adv --keyserver keyserver.ubuntu.com --recv-keys EEA14886 && \
    echo oracle-java8-installer shared/accepted-oracle-license-v1-1 select true | /usr/bin/debconf-set-selections

RUN apt-get update && apt-get install --no-install-recommends -y oracle-java8-installer oracle-java8-set-default maven 
RUN npm install --quiet --global vue-cli
RUN npm install --quiet --global angular-cli
RUN apt-get install -y expect
COPY ./vue.sh /usr/bin/vue-generate
RUN chmod +x /usr/bin/vue-generate

RUN apt-get autoremove && apt-get clean
WORKDIR /root/
COPY devops-generate-code ./
RUN chmod +x ./devops-generate-code
EXPOSE 8080
CMD ["./devops-generate-code"]