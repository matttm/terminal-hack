# terminal-hack

Development: Inactive

## Description

This is a terminal based game written in Go, which is inspired by terminal hacking in Fallout 3. In this game, a 2d grid is filled with single 'dud' characters and words. The objective is to determine which word is the winning word in a limited number of guesses. Each incorrect guess will tell you how many of the selected letters are correct. A letter is correct if it has the same index and actual character as in the winning word. For example. if the winning word is `book` and selected word is `bang`, the selected word has one correct character.

<img width="1728" alt="Screenshot 2024-12-29 at 12 43 47 PM" src="https://github.com/user-attachments/assets/d285eb01-1a59-4f7a-b3de-99a8e0336ee2" />

## Design

This game is created with a very low-level drawing library called Termbox. Because of this, I had to program all navigation/animation functionaliy.

## Authors

-   Matt Maloney : matttm

## Contribute

If you want to contribute, just send me a message.
