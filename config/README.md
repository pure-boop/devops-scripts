#!/bin/bash

# Define a function to greet the user
greet() {
  echo "Hello, $1!"
}

# Define a function to calculate the sum of two numbers
add() {
  echo $(($1 + $2))
}

# Define a function to multiply two numbers
multiply() {
  echo $(($1 * $2))
}

# Define a function to calculate the factorial of a number
factorial() {
  if [ $1 -eq 0 ]; then
    echo 1
  else
    echo $(($1 * $(factorial $(( $1 - 1 ))))
  fi
}

# Define a function to calculate the greatest common divisor of two numbers
gcd() {
  local a=$1
  local b=$2
  while [ $b -ne 0 ]; then
    local temp=$b
    b=$(( a % b ))
    a=$temp
  done
  echo $a
}

# Call the functions
greet "Alice"
result=$(add 5 3)
echo "The sum is: $result"
result=$(multiply 5 3)
echo "The product is: $result"
result=$(factorial 5)
echo "The factorial is: $result"
result=$(gcd 48 18)
echo "The GCD is: $result"