FROM scratch

ADD main /

EXPOSE 8090
CMD ["/main"]