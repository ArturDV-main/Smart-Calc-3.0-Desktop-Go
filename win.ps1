mkdir build -Force
Copy-Item cpp/src/* ./build
Set-Location ./build
g++ -shared -o smart_calc.dll ./*.cc
Set-Location ../
Copy-Item ./build/smart_calc.dll ./pkg/calcadapter/
Copy-Item ./appicon.png ./build
Copy-Item ./icon.ico ./build
wails build
Copy-Item ./smart_calc.dll ./build/bin
./build/bin/smartcalc.exe