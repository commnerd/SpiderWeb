#!/bin/bash

err() { printf "%s\n" "$*" >&2; exit 1; }

if [ $(whoami) != "root" ]
then
	err "Please run provisioning as root."
fi

if [ ! ${PROJECT_SLUG} ]
then
	PROJECT_SLUG=spiderweb
fi

mkdir -p /var/log/${PROJECT_SLUG}
mkdir -p /var/run/${PROJECT_SLUG}
mkdir -p /etc/${PROJECT_SLUG}
