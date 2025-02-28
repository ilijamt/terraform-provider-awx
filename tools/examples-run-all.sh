DIRNAME=$(dirname -- "$0")
ROOTDIR=$(realpath "$DIRNAME/../")
TARGETDIR=$(realpath "$DIRNAME/../examples")
COVERAGEDIR=$ROOTDIR/build/coverage

export GOCOVERDIR=$COVERAGEDIR

cd $TARGETDIR
source .envrc

function run() {
    T=$TARGETDIR/$1
    cd $T
    terraform plan -out plan
    terraform apply plan
    terraform apply -destroy -auto-approve
}

set -x
rm -rf $COVERAGEDIR
mkdir -p $COVERAGEDIR
cd $TARGETDIR
bash $ROOTDIR/tools/examples-clean.sh
cd $ROOTDIR
make build-debug

run application
run credential_with_input_source
run credentials
# run general
run instance_groups
run inventory
# run job_template_survey_spec
# run preload_data
# run objectroles
run roles
run settings
run settings_authentication
run settings_github


go tool covdata textfmt -i=$COVERAGEDIR -o $ROOTDIR/build/coverage.data.out
go tool cover -html=$ROOTDIR/build/coverage.data.out -o $ROOTDIR/build/coverage.data.html