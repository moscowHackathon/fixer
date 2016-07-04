# Fixer

Slack client to communication consumers with arkon

## Installation 

	go get github.com/moscowHackathon/fixer
	export GO15VENDOREXPERIMENT=1
	cd $GOPATH/src/github.com/moscowHackathon/fixer && glide install

## Run
	$GOPATH/bin/fixer --slack-token=YOUR-TOKEN-IN-SLACK --cert=/cert/fullchain.pem --key=/cert/privkey.pem
