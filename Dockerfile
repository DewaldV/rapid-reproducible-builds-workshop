FROM maven:3.3.3-jdk-8

COPY . /build/
WORKDIR /build

ENTRYPOINT [ "/usr/bin/mvn" ]
CMD [ "clean test" ]
