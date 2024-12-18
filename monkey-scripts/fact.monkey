let fact = fn(n) {
    if (n == 1) {
        1
    } else {
        n * fact(n - 1)
    }
}
puts(fact(5))
