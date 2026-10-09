FROM ubuntu:latest
LABEL authors="franklinesable"

ENTRYPOINT ["top", "-b"]