Given two strings S and T, return if they equal when both are typed out.
Any '#' that appears in the string counts as backspace.

e.g. S = 'a#b#' and T = 'b#c#' are equal
since S = '' and T = '' after using # as backspace chars.

Constraints:
- Are two empty strings equal?
    - Yes
- What happens if two hashes appears beside each other?
    - remove the 2 chars on the left if chars are present
- If there are no chars left to remove for the corresponding hashes
    - then ignore the chars i.e. 'a##b' => 'b'
- Does case sensitivity matters?
    - Yes, 'A' is not equal to 'a'

Test cases:
- S = 'a#b#c' T = 'b#c#c' true
    - equal since S = 'c' and T = 'c'
- S = 'a##c' and T = 'b#c#c'
    - equal since S = 'c' and T = 'c'