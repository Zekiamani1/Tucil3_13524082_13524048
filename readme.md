<h1 align="center">Tucil 3 Stima Ice Sliding Solver</h1>

---
<p align="center">
  <img src="assets/bocchi.png" alt="Alt text">
</p>

<div align="center" id="contributor">
  <strong>
    <table align="center">
      <tr align="center">
        <td></td>
        <td>NIM</td>
        <td>Nama</td>
      </tr>
      <tr align="center">
        <td><img src="https://github.com/achideon.png" width="50" /></td>
        <td>13524048</td>
        <td>Josh Reinhart Zidik</td>
      </tr>
      <tr align="center">
        <td><img src="https://github.com/zekiamani1.png" width="50" /></td>
        <td>13524082</td>
        <td>Zeki Amani</td>
      </tr>
    </table>
  </strong>
</div>

<div align="center">
  <h3 align="center">Tech Stacks</h3>

  <p align="center">
    <img alt="Golang" src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white">
  </p>
</div>


## What is this
Ice Sliding Puzzle is a logic game in which players must move a character from the starting point to the exit point across a slippery ice surface. Players can only move horizontally or vertically, but because the surface is slippery, the character will not stop moving until it hits a wall or an obstacle.\
This program is a solver to the ice sliding puzzle using pathfinding algorithm such as GBFS,UCS, and A*

## Installation & Setup

### Requirements
- Golang (version 1.26.2 or higher)
- Git
- GCC/Build Tools (for Fyne GUI compilation on Linux)


### Installing Dependencies
1. Golang

   ```bash
   sudo apt install golang-go
   ```

2. Git

   ```bash
   sudo apt install git
   ```

3. Build Tools (for Linux users, required for Fyne GUI)

   ```bash
   sudo apt install build-essential libgl1-mesa-dev xorg-dev
   ```

4. Go Module Dependencies

   After cloning the repository, dependencies will be automatically downloaded:
   ```bash
   go mod download
   ```



## How to Run

### Clone the Repository

Open your terminal and clone the repository:

```bash
git clone https://github.com/Zekiamani1/Tucil3_13524082_13524048.git
```

### Navigate to the Directory

```bash
cd Tucil3_13524082_13524048/src
```

### Run the Application

```bash
go run .
```
