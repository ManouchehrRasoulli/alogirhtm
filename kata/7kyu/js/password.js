function password(str) {
    console.log(str);

    if (str.length < 8) {
        return false;
    }

    return /[a-z]/.test(str) && /[A-Z]/.test(str) && /[0-9]/.test(str);
}