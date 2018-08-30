FROM prom/prometheus:v2.3.2

COPY build/prometheus/prometheus.yml /etc/prometheus/

USER root
RUN chown nobody:nogroup /etc/prometheus/prometheus.yml
USER nobody