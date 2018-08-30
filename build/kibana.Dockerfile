FROM docker.elastic.co/kibana/kibana-oss:6.4.0
COPY build/kibana/kibana.yml /usr/share/kibana/
USER root
RUN chown kibana:kibana kibana.yml
USER kibana