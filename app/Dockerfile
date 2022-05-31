# Stage 1) Building Flutter web
FROM ubuntu:20.04 as builder

RUN apt update && apt install -y curl git unzip xz-utils zip libglu1-mesa

# download Flutter SDK from Flutter Github repo
RUN git clone https://github.com/flutter/flutter.git /usr/local/flutter

# Set flutter environment path
ENV PATH="/usr/local/flutter/bin:/usr/local/flutter/bin/cache/dart-sdk/bin:${PATH}"

# Run flutter doctor
RUN flutter doctor

# Enable flutter web
RUN flutter channel master
RUN flutter upgrade
RUN flutter config --enable-web

# Copy files to container and build
RUN mkdir /app/
COPY . /app/
WORKDIR /app/
RUN flutter build web

# Stage 2) Running the web server
FROM nginx
COPY --from=builder /app/build/web /usr/share/nginx/html