FROM golang:1.14.4
RUN mkdir /app
ADD . /app/
WORKDIR /app
EXPOSE 80
CMD ["make", "run"]