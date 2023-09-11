echo "Building the Go project..."
go build -o tetris-optimiser .
if [ $? -ne 0 ]; then
    echo "Error building the project. Exiting."
    exit 1
fi
echo "Testing against input files..."
echo "----------------------------------"
echo "----------------------------------"

for file in ./test/*; do
    output=$(./tetris-optimiser "$file")
    echo "$output"    
done
echo "All tests completed."
