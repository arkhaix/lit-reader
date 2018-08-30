FROM docker.elastic.co/logstash/logstash-oss:6.4.0
COPY build/logstash/logstash.yml /usr/share/logstash/config/
USER root
RUN chown logstash:logstash config/logstash.yml
USER logstash