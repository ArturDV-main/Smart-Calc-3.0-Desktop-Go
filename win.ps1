mkdir build -Force
mkdir build/bin -Force
Copy-Item cpp/src/* ./build
Set-Location ./build
g++ -static-libstdc++ -static-libgcc --static -lstdc++ -lpthread -shared -o smart_calc.dll ./*.cc
Set-Location ../
Copy-Item ./build/smart_calc.dll ./build/bin
Copy-Item ./build/smart_calc.dll ./
Copy-Item ./build/smart_calc.dll ./pkg/calcadapter/
Copy-Item ./appicon.png ./build
Copy-Item ./icon.ico ./build
wails build -nsis
./build/bin/smartcalc.exe
