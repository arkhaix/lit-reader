FROM docker.elastic.co/elasticsearch/elasticsearch:6.4.0
COPY build/elasticsearch/elasticsearch.yml /usr/share/elasticsearch/config/
USER root
RUN chown elasticsearch:elasticsearch config/elasticsearch.yml
USER elasticsearch