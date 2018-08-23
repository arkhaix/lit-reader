FROM cockroachdb/cockroach:v2.0.5
COPY cr-init.sh /cockroach/
COPY cr-init.sql /cockroach/
