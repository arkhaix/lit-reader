FROM cockroachdb/cockroach:v2.0.5
COPY build/cockroach-init/cr-init.sh /cockroach/
COPY build/cockroach-init/cr-init.sql /cockroach/
