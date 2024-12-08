mkdir build -Force
cp cpp/src/* ./build
cd ./build
g++ -shared -o smart_calc.dll ./*.cc
cp ./smart_calc.dll ../
cp ./smart_calc.dll ../pkg/calcadapter/
cd ..
echo "done"