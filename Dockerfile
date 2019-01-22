FROM node:10.15-alpine

# set working directory
RUN echo `cd usr && ls -a`
RUN mkdir -p usr/src/shyft_ui
WORKDIR /usr/src/shyft_ui

# add `/usr/src/app/node_modules/.bin` to $PATH
#ENV PATH /usr/src/shyft_ui/node_modules/.bin:$PATH

# install and cache app dependencies
COPY . .
RUN npm install
# RUN npm add react-scripts@1.1.1 -g --silent


# start app

EXPOSE 3000

CMD ["npm", "start"]
