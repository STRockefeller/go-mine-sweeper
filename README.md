# Mine Sweeper

## Overview

This Mine Sweeper project is a modern take on the classic Minesweeper game. Developed in Go, it features a terminal-based user interface for a nostalgic yet fresh gaming experience. The game allows players to select from multiple difficulty levels and interact through simple, intuitive commands.

## Installation

To run this project, you need to have Go installed on your machine. Follow these steps to get started:

1. **Clone the Repository:**

   ```shell
   git clone https://github.com/STRockefeller/go-mine-sweeper.git
   ```

2. **Navigate to the Project Directory:**

   ```shell
   cd go-mine-sweeper
   ```

## Dependencies

This project relies on the following external libraries:

- `github.com/charmbracelet/huh`: For creating interactive forms in the terminal.

Ensure you have these dependencies installed by running:

```shell
go get -u
```

## Usage

To start the game, run:

```shell
go run .
```

Upon launch, you'll be prompted to select a difficulty level: Easy, Medium, or Hard. The game grid and the number of mines will vary based on your choice.

### Game Commands

- **Reveal a Cell:** Type `R` followed by the x and y coordinates of the cell (e.g., `R 3 4`).
- **Flag a Cell:** Type `F` followed by the x and y coordinates of the cell (e.g., `F 2 1`).

Coordinates are 0-indexed from the top-left corner of the board.

### Example

```shell
Enter your move (e.g., 'R 3 4' to reveal or 'F 2 1' to flag) (0-indexed from the top-left corner): R 0 2
```

### Screenshots

![difficulty selection](https://i.imgur.com/Ievjiex.png)
![game play](https://i.imgur.com/DhKX8VA.png)
![game play](https://i.imgur.com/rio4Jai.png)
![game play](https://i.imgur.com/P6nAAxy.png)

## License

This project is licensed under the [MIT License](LICENSE).

---

Happy gaming! 🚀
