FROM node:10.15-alpine as builder

# set working directory
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# add `/usr/src/app/node_modules/.bin` to $PATH
ENV PATH /usr/src/app/node_modules/.bin:$PATH

# install and cache app dependencies
#COPY package.json /usr/src/app/package.json
COPY ./ /usr/src/app
RUN npm install --silent
RUN npm install react-scripts@1.1.1 -g --silent

# build
RUN npm run build


FROM nginx:1.14-alpine
COPY --from=builder /usr/src/app/build/ /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]


