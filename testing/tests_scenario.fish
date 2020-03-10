#! /usr/bin/env fish

set TEST_CREDS "Eleanor\npass\n"
set CORE "\n"
set CONNECTOR "\033[B\n"

function echoed
  echo -ne "$argv[1]"
end

function npm_init
  npm init -y 1>/dev/null
end

function dessert_init
  if [ "$argv[1]" = "core" ]
    echoed "$CORE" | ./dessert-cli init
  else if [ "$argv[1]" = "connector" ]
    echoed "$CONNECTOR" | ./dessert-cli init
  else
    echoed "Invalid type. Aborting!"
    exit 1
  end
end

function dessert_login
  echoed "$TEST_CREDS" | ./dessert-cli login
end

dessert_init "core"
dessert_login