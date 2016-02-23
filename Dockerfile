FROM maven:3.3.3-jdk-8

COPY settings.xml /root/.m2/
COPY . /build/
WORKDIR /build

ENTRYPOINT [ "/usr/bin/mvn" ]
