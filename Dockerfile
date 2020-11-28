FROM ketouem/ag-alpine

WORKDIR /
COPY ./gowebexec /gowebexec
COPY ./data /data
RUN mv /the_silver_searcher/ag /
ENV PORT 3000
ENTRYPOINT ["/gowebexec"]

