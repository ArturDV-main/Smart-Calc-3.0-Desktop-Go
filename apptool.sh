#!/bin/bash -e

BINARY="build/bin/smartcalc.app/Contents/MacOS/smartcalc"
FRAMEW_FOLDER="build/bin/smartcalc.app/Contents/Frameworks/"


function DoInstallNameTool {
    xLIB="$1"
    xLIB_NAME="$2"
    xBINARY="$3"
    echo install_name_tool -change \"${xLIB}\" \"@executable_path/../Frameworks/${xLIB_NAME}\" \"${xBINARY}\"
    install_name_tool -change ${xLIB} "@executable_path/../Frameworks/${xLIB_NAME}" "${xBINARY}"
}

for LIB in $(otool -L "${BINARY}"|grep libsmart_calc.dylib|cut -d '(' -f -1)
do
    echo "Handling Lib: $LIB"
    LIB_NAME=$(basename "$LIB")
    echo "    Adding ${LIB_NAME}"
    mkdir "${FRAMEW_FOLDER}"
    cp -Rf "${LIBSNDFILE_DIR}build/${LIB_NAME}" "${FRAMEW_FOLDER}"

    DoInstallNameTool "$LIB" "$LIB_NAME" "$BINARY"
done
echo "Done"