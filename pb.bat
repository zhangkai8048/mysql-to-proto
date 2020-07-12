@echo off

for %%i in (proto/*.proto) do (
    protoc ./proto/%%i --java_out=./src/main/java
    jupiter protoc -f ./proto/%%i   -o ./services/ -
    echo exchange %%i To go file successfully!
)

