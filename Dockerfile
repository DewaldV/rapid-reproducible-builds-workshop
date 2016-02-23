FROM maven:3.3.3-jdk-8

COPY . /build/
WORKDIR /build
COPY settings.xml /root/.m2/

ENTRYPOINT [ "/usr/bin/mvn" ]
CMD [ "clean test" ]
