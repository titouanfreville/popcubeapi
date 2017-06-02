#!/bin/bash


echo -e "____   ___   ____    __  __ __  ____     ___ "
echo -e "|    \ /   \ |    \  /  ]|  |  ||    \   /  _]"
echo -e "|  o  )     ||  o  )/  / |  |  ||  o  ) /  [_ "
echo -e "|   _/|  O  ||   _//  /  |  |  ||     ||    _]"
echo -e "|  |  |     ||  | /   \_ |  :  ||  O  ||   [_ "
echo -e "|  |  |     ||  | \     ||     ||     ||     |"
echo -e "|__|   \___/ |__|  \____| \__,_||_____||_____|"
echo ""

if [[ -z "$MYSQL_HOST" ]]; then
  echo "Env var not set : MYSQL_HOST"
  echo "MYSQL_HOST is IP/hostame for database"
  exit 1
fi
waitforit --host=${MYSQL_HOST} --port=${MYSQL_PORT} -t 0 -- echo "Db is ready"
echo "=> Installing deps..."
go install
echo "=> Starting popcube api..."
popcubeapi
