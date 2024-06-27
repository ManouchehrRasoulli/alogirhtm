function combat(health, damage) {
    // Write your code here
    h = health - damage;
    if (h < 0) {
        return 0;
    }
    return h;
}