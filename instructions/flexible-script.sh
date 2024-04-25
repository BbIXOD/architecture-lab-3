#!/bin/bash

# Перевіряємо, чи були передані аргументи
if [ "$#" -eq 0 ]; then
    echo "Використання: $0 <аргумент1> <аргумент2> ..."
    exit 1
fi

# Цикл для обробки всіх аргументів
for arg in "$@"
do
    curl -X POST http://localhost:17000 -d "$arg"
done