FROM node:8.11.3-alpine

# set working directory
RUN echo `cd usr && ls -a`
RUN mkdir -p usr/src/shyft_ui
WORKDIR /usr/src/shyft_ui

# add `/usr/src/app/node_modules/.bin` to $PATH
ENV PATH /usr/src/shyft_ui/node_modules/.bin:$PATH

# install and cache app dependencies
COPY ./shyft_ui/package.json /usr/src/shyft_ui/package.json
RUN npm install
# RUN npm add react-scripts@1.1.1 -g --silent
COPY ./shyft_ui/src /usr/src/shyft_ui/src
COPY ./shyft_ui/scripts /usr/src/shyft_ui/scripts
COPY ./shyft_ui/config /usr/src/shyft_ui/config
COPY ./shyft_ui/public /usr/src/shyft_ui/public

# start app

EXPOSE 3000

CMD ["npm", "start"]