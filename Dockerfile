FROM scratch

ADD /bin/main /

EXPOSE 8090
CMD ["/main"]