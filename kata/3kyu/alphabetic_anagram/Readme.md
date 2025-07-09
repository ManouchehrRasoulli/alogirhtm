https://www.codewars.com/kata/53e57dada0cb0400ba000688/train/go

Consider a "word" as any sequence of capital letters A-Z (not limited to just "dictionary words"). For any word with at least two different letters, there are other words composed of the same letters but in a different order (for instance, STATIONARILY/ANTIROYALIST, which happen to both be dictionary words; for our purposes "AAIILNORSTTY" is also a "word" composed of the same letters as these two).

We can then assign a number to every word, based on where it falls in an alphabetically sorted list of all words made up of the same group of letters. One way to do this would be to generate the entire list of words and find the desired one, but this would be slow if the word is long.

Given a word, return its number. Your function should be able to accept any word 25 letters or less in length (possibly with some letters repeated), and take no more than 500 milliseconds to run. To compare, when the solution code runs the 27 test cases in JS, it takes 101ms.

For very large words, you'll run into number precision issues in JS (if the word's position is greater than 2^53). For the JS tests with large positions, there's some leeway (.000000001%). If you feel like you're getting it right for the smaller ranks, and only failing by rounding on the larger, submit a couple more times and see if it takes.

Python, Java and Haskell have arbitrary integer precision, so you must be precise in those languages (unless someone corrects me).

C# is using a long, which may not have the best precision, but the tests are locked so we can't change it.

Sample words, with their rank:
ABAB = 2
AAAB = 1
BAAA = 4
QUESTION = 24572
BOOKKEEPER = 10743


This problem is about **finding the rank (or position)** of a word in the list of all its **permutations**, sorted **alphabetically**, assuming:

* Only **uppercase English letters (A–Z)** are used.
* Words may contain **repeated letters**.
* The "rank" is **1-based** (first permutation is rank 1).

---

### 🔍 **What’s being asked?**

Given a word (e.g., `"BOOKKEEPER"`), compute its **position (rank)** in the list of **all unique permutations** of its letters, sorted alphabetically.

**BUT** don’t generate all permutations (too slow). Instead, **calculate** the rank mathematically.

---

### 📘 **How does it work?**

Let’s take an example:

#### Example: `"ABAB"`

* Step 1: List all **distinct permutations** alphabetically:

    1. AABB
    2. ABAB ← this is our input
    3. ABBA
    4. BAAB
    5. BABA
    6. BBAA

→ So `"ABAB"` is at position **2**.

But instead of generating them, we use **combinatorics** to **count how many permutations would appear before** the given one.

---

### 🧮 **General Method**

For a word like `"BOOKKEEPER"`:

1. Start with the first letter `'B'`.
2. For every letter **lexicographically less than `'B'`**, count how many permutations **start** with that letter (and use up the remaining letters).
3. Add up those permutations → that gives us how many permutations come **before** `"BOOKKEEPER"`.
4. Do this for **each character** in the word, moving left to right.
5. Finally, add 1 (because rank is 1-based).

To compute permutations with **repeated letters**, use:

$$
\text{Permutations} = \frac{n!}{\text{product of }(f_i!)}
$$

Where:

* $n$ = total number of letters remaining.
* $f_i$ = frequency of each letter.

---

### ✅ **Let's verify the examples**

#### 1. `"AAAB"`

* Alphabetical permutations:

    1. AAAB ← our word
    2. AABA
    3. ABAA
    4. BAAA

Rank = **1**

---

#### 2. `"ABAB"`

Alphabetical permutations:

1. AABB
2. ABAB ← our word
3. ABBA
4. BAAB
5. BABA
6. BBAA

Rank = **2**

---

#### 3. `"BAAA"`

Permutations:

1. AAAB
2. AABA
3. ABAA
4. BAAA ← our word

Rank = **4**

---

#### 4. `"QUESTION"`

This is trickier due to 8 unique letters.
The rank is **24572**, computed using the mathematical process above.

---

#### 5. `"BOOKKEEPER"`

Also has duplicate letters:

* B:1, O:2, K:2, E:3, P:1, R:1

→ There are **millions** of permutations possible. But this word ranks **10743rd** alphabetically among them.

---

### 🧠 Why not brute force?

* 25-letter word with many repeated letters could have **billions** of permutations.
* Generating all and sorting would be **impractical**.
* So we calculate using **math (factorials and frequency counts)**.

---

Would you like to see the full Python function that calculates the rank efficiently?

Great! Here's a clear **pseudocode** for computing the **lexicographic rank** of a word among all its **unique permutations**, accounting for **repeated letters**:

---

### 🧩 **Pseudocode: Lexicographic Rank with Repeated Letters**

```plaintext
function findLexicographicRank(word):
    rank = 1
    length = length of word
    freq = array[26] of integers initialized to 0

    // Count frequency of each character
    for each character c in word:
        freq[ord(c) - ord('A')] += 1

    // Precompute factorials up to length
    factorial = array[length + 1]
    factorial[0] = 1
    for i from 1 to length:
        factorial[i] = factorial[i - 1] * i

    for i from 0 to length - 1:
        current_char = word[i]

        // Count characters less than current_char
        for ch from 'A' to current_char - 1:
            if freq[ord(ch) - ord('A')] == 0:
                continue

            // Temporarily reduce frequency
            freq[ord(ch) - ord('A')] -= 1

            // Calculate number of permutations with this smaller character at position i
            denom = 1
            for j from 0 to 25:
                denom *= factorial[freq[j]]
            rank += factorial[length - i - 1] / denom

            // Restore frequency
            freq[ord(ch) - ord('A')] += 1

        // Decrease frequency of current_char (it is now fixed)
        freq[ord(current_char) - ord('A')] -= 1

    return rank
```

---

### 🧠 **Key Concepts Used**

* **Factorials** to count permutations.
* **Multiset permutations formula** to handle repeated letters:

  $$
  \frac{n!}{f_1! \cdot f_2! \cdot \dots \cdot f_k!}
  $$
* For each character, we **count how many permutations start with a smaller letter**.

---

Let me know if you want this translated into a specific programming language like **Python**, **Go**, or **C++**.


Sure! Here's **clear and commented pseudocode** for computing the **lexicographic rank** of a word among all its **unique permutations**, accounting for **repeated characters**:

---

### 🧠 **Pseudocode: Lexicographic Rank of a Word with Repeated Characters**

```
function rank(word):
    rank ← 1
    n ← length(word)
    freq ← map from char to int

    // Step 1: Count frequency of each character
    for char in word:
        freq[char] += 1

    // Step 2: For each position in the word
    for i from 0 to n - 1:
        current_char ← word[i]

        // Count characters less than current_char
        for ch from 'A' to current_char - 1:
            if freq[ch] > 0:
                freq[ch] -= 1

                // Calculate number of permutations of remaining letters
                perms ← factorial(n - i - 1)
                for each letter in freq:
                    perms = perms / factorial(freq[letter])

                rank += perms

                // Restore frequency
                freq[ch] += 1

        // Reduce the frequency of the current_char (we are "using" it)
        freq[current_char] -= 1

    return rank
```

---

### 📦 **Helper Functions (assumed available)**

* `factorial(x)` — returns `x!` (factorial of x).
* `freq` is a dictionary or map tracking how many times each letter appears.

---

### 📝 Example Input

If you call:

```
rank("BOOKKEEPER")
```

It will return:

```
10743
```

---

### ✅ Benefits

* Handles **repeated characters** correctly.
* Doesn't generate permutations — it **counts** them efficiently.
* Works for words up to \~20–25 characters without overflow if you use big integers.

---

Would you like me to convert this into **Python**, **Go**, or another programming language?
