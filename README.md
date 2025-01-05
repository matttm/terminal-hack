# terminal-hack

Development: Main development completed with intermittent improvements

## Description

This is a terminal based game written in Go, which is inspired by terminal hacking in Fallout 3. In this game, a 2d grid is filled with single 'dud' characters and words. The objective is to determine which word is the winning word in a limited number of guesses. Each incorrect guess will tell you how many of the selected letters are correct. A letter is correct if it has the same index and actual character as in the winning word. For example. if the winning word is `book` and selected word is `bang`, the selected word has one correct character.

<img width="1920" alt="Screenshot 2025-01-04 at 2 13 04 PM" src="https://github.com/user-attachments/assets/2436e090-8509-4d30-b40b-b78cff1148ef" />

## Design

This game is created with a very low-level drawing library called Termbox. Because of this, I had to program all navigation/animation functionaliy.

## Authors

-   Matt Maloney : matttm

## Contribute

If you want to contribute, just send me a message.
