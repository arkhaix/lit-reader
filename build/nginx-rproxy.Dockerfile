FROM nginx:1.15
COPY build/nginx/nginx.conf /etc/nginx/nginx.conf
#COPY build/nginx/letsencrypt /etc/letsencrypt