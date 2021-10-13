ARG NGINX_VERSION="latest"
FROM nginx:$NGINX_VERSION
ARG COMMIT_HASH="none"

COPY ./index.html /usr/share/nginx/html/
RUN echo ${COMMIT_HASH}
RUN sed -i -e "s/{{COMMIT_HASH}}/${COMMIT_HASH}/g" /usr/share/nginx/html/index.html
