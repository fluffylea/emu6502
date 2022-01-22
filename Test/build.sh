#!/bin/zsh
#set -xe

OPTHIS_CMD="Ophis/bin/ophis"
EMU6502_CMD="../emu6502"
BUILD_DIR="build"
SOURCE_FILES=(
  "helloWorld.oph"
  "printChar.oph"
  "printDec8.oph"
  "printDec16.oph"
  "fib.oph"
)

function clean() {
  rm -r $BUILD_DIR
}

function compile() {
  pushd ..
  go build
  popd || exit 1

  for file in $SOURCE_FILES; do
    name=${file%.*}
    if ! $OPTHIS_CMD -l "$BUILD_DIR/$name.l" \
                     -m "$BUILD_DIR/$name.m" \
                     -o "$BUILD_DIR/$name.rom" \
                     "$file"; then
      echo "Failed to build $file"
      exit 1
    fi
  done
}

function run() {
  compile
  for file in $SOURCE_FILES; do
    name=${file%.*}
    if ! $EMU6502_CMD -loglevel info \
                      -runtime 1 \
                      -rom "$BUILD_DIR/$name.rom" \
                      -mapping "$BUILD_DIR/$name.m" \
                      -listing "$BUILD_DIR/$name.l"; then
      echo "Failed to run $file"
      exit 1
    fi
  done
}

function run_single() {
  compile
  file=$1
  name=${file%.*}
  if ! $EMU6502_CMD -loglevel info \
                    -runtime 1 \
                    -rom "$BUILD_DIR/$name.rom" \
                    -mapping "$BUILD_DIR/$name.m" \
                    -listing "$BUILD_DIR/$name.l"; then
    echo "Failed to run $file"
    exit 1
  fi
}

function save_test() {
  compile
  for file in $SOURCE_FILES; do
    name=${file%.*}
    if ! $EMU6502_CMD -loglevel info \
                      -runtime 1 \
                      -rom "$BUILD_DIR/$name.rom" \
                      -mapping "$BUILD_DIR/$name.m" \
                      -listing "$BUILD_DIR/$name.l" >"$name.output"; then
      echo "Failed to run $file"
      exit 1
    fi
  done
}

function test() {
  compile
  for file in $SOURCE_FILES; do
    name=${file%.*}
    if ! $EMU6502_CMD -loglevel info \
      -runtime 1 \
      -rom "$BUILD_DIR/$name.rom" \
      -mapping "$BUILD_DIR/$name.m" \
      -listing "$BUILD_DIR/$name.l" >"$BUILD_DIR/$name.output2"; then
      echo "Failed to run $file"
      exit 1
    else
      # Verify that the output is the same
      # If not, exit with error
      if ! diff -q "$name.output" "$BUILD_DIR/$name.output2"; then
        echo "Test $file failed: Output differs"
        exit 1
      fi
    fi
  done
}

# if build dir doesn't exist, create it
if [ ! -d "$BUILD_DIR" ]; then
  mkdir "$BUILD_DIR"
fi

# Clean command deletes all .l, .m and .rom files
case "$1" in
clean)
  clean
  ;;
compile)
  compile
  ;;
run)
  run
  ;;
run_single)
  run_single "$2"
  ;;
save_test)
  save_test
  ;;
test)
  test
  ;;
*)
  echo "Usage: $0 {clean|compile|run|save_test|test}"
  ;;
esac
