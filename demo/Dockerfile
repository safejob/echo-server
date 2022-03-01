FROM ubuntu


WORKDIR /app
COPY demo /app/demo
RUN chmod +x /app/demo

EXPOSE 80

CMD ["./demo","-s","200"]
