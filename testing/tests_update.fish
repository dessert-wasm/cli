set FOLDER "./cmd/stdoutTests/"
set CONFIGS "./cmd/configTests/"
export DESSERT_GRAPHQL_URI="http://localhost:5000/"

# version
./dessert-cli version > "$FOLDER""version_basic"

# login basic
npm init -y 1>/dev/null
echo -ne "Y\n" | ./dessert-cli init 1>/dev/null
echo -ne "Eleanor\npass\n" | ./dessert-cli login > "$FOLDER""login_basic"
# login failed
echo -ne "a\na\n" | ./dessert-cli login > "$FOLDER""login_failed"
rm package.json

# logout basic
echo -ne "Eleanor\npass\n" | ./dessert-cli login 1>/dev/null
./dessert-cli logout > "$FOLDER""logout_basic"

# logout already
./dessert-cli logout > "$FOLDER""logout_already"

# missing yml

# missing json
# init missing json
echo -ne "Y\n" | ./dessert-cli init > "$FOLDER""init_missing_json"

# publish missing json
npm init -y 1>/dev/null
echo -ne "Y\n" | ./dessert-cli init 1>/dev/null
rm package.json
echo -ne "whatever\n" | ./dessert-cli publish > "$FOLDER""publish_missing_json"
rm dessert.yml

# publish missing yml
npm init -y 1>/dev/null
echo -ne "Y\n" | ./dessert-cli init 1>/dev/null
rm dessert.yml
echo -ne "whatever\n" | ./dessert-cli publish > "$FOLDER""publish_missing_yml"
rm package.json

# publish ok
npm init -y 1>/dev/null
echo -ne "Y\n" | ./dessert-cli init 1>/dev/null
echo -ne "Eleanor\npass\n" | ./dessert-cli login 1>/dev/null
./dessert-cli publish > "$FOLDER""publish_ok"
./dessert-cli logout 1>/dev/null 
rm dessert.yml
rm package.json

# replaces missing json
./dessert-cli replaces module1 module2 > "$FOLDER""replaces_missing_json"

# -------------------- #
# init core
npm init -y 1>/dev/null
echo -ne "Y\n" | ./dessert-cli init > "$FOLDER""init_core"
cp package.json "$CONFIGS""core_dessert.json"
rm package.json

# replaces
npm init -y 1>/dev/null
./dessert-cli replaces module1 module2 > "$FOLDER""replaces_basic"
cp package.json "$CONFIGS""replaces_dessert.json"
rm package.json

# replaces invalid arguments
./dessert-cli replaces > "$FOLDER""replaces_invalid_args"

# init connector
npm init -y 1>/dev/null
echo -ne "N\n" | ./dessert-cli init > "$FOLDER""init_connector"
cp package.json "$CONFIGS""connector_dessert.json"
rm package.json

# init invalid answer
npm init -y 1>/dev/null
echo -ne "invalid\n" | ./dessert-cli init > "$FOLDER""init_invalid_answer"

rm package.json
rm dessert.yml

echo "Done!"