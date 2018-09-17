FROM golang
ENV GOPATH /apt:$GOPATH
COPY . /apt/src/5317c349-8ade-4a35-93c4-2401101056db
RUN cd /apt/src/5317c349-8ade-4a35-93c4-2401101056db && \
	sh install.sh && go build
CMD /apt/src/5317c349-8ade-4a35-93c4-2401101056db/5317c349-8ade-4a35-93c4-2401101056db
EXPOSE 8080