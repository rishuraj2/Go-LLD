# Functional Requirements
- The game is played on 3x3 board
- Players take alternate turns to place their Symbol ('X' or 'O')
- Player vs Player only
- Game should detect and announce winner
- Game should declare a draw is all the cells are filled and no player has won
- Reject invalid move by notifying the player
- Game moves should be hardcoded by either a driver class or main function

# Non-Functional Requirements
- Design should follow object-oriented principles with clear responsibility and seperation of concerns.
- System should be modular and extensible to support future features like larger board, AI opponent, scoreboard and history tracking to allow undo or move replay.
- Game logic should be testable and easy to maintain.
- System should provide clear console output, reflecting the current state of the game board.

# Core Entities

### Enums
- Symbols ('X', 'O', 'EMPTY')
- GameState ('WINNER_X', 'WINNER_O', 'DRAW')

### Data Class
- Cell
- Player

### Core Class
- Board
- Game

# Responsibility
- **Symbol**: Enum having 'X', 'O' and 'Empty' used by **Cell** and **Player**.
- **GameState**: Enum having 'WINNER_X', 'WINNER_O' AND 'DRAW' used by **Game** to determine it's status.
- **Cell**: It has **Symbol** and is used as single unit in **Board**. It has behaviour to add and get the **Symbol**.
- **Board**: It has a square matrix of **Cells** of dynamic size that **Game** owner can provide. It also checks for game status on each move.
- **Player**: It holds player's details and is used by **Game**.
- **Game**: It is the orchestrator that holds **Players**, **GameState** and **Board** and manages the overall game creation and termination.