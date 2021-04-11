FROM gcr.io/distroless/base
COPY dotenv /
ENTRYPOINT ["/dotenv"] 