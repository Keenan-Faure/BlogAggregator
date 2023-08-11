FROM debian:stable-slim

COPY out /bin/out

CMD [ "/bin/out" ]