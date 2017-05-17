#!/bin/bash
# CUSTOM
POPCUBE_CONFIG="/docker-entrypoint-initdb.d/2_custom_init.sql"

if [[ "$ORG_ORGANISATIONNAME" == "None" ]]; then
  echo "Popcube custom env varibale not set"
  exit 1
fi
if [[ "$ORG_DESCRIPTION" == "None" ]]; then
  echo "Popcube custom env varibale not set"
  exit 1
fi
if [[ "$ORG_AVATAR" == "None" ]]; then
  echo "Popcube custom env varibale not set"
  exit 1
fi
if [[ "$ORG_DOMAIN" == "None" ]]; then
  echo "Popcube custom env varibale not set"
  exit 1
fi

# Setting orgs starter params
sed -i "s/%org_organisationName%/${ORG_ORGANISATIONNAME}/g" "$POPCUBE_CONFIG"
sed -i "s/%org_description%/${ORG_DESCRIPTION}/g"  "$POPCUBE_CONFIG"
sed -i "s/%org_avatar%/${ORG_AVATAR}/g"  "$POPCUBE_CONFIG"
sed -i "s/%org_domain%/${ORG_DOMAIN}/g"  "$POPCUBE_CONFIG"
