const {Cell, LiveCell, DeadCell} = require('../cell/cell');
const { CellAddress } = require('../cellAddress/cellAddress');

class Grid {

    constructor(rows, columns) {
        if (rows <= 0) {
            throw new Error('Invalid rows');
        }
        if (columns <= 0) {
            throw new Error('Invalid columns');
        }
        this.cells = new Array(rows);
        for (let i = 0; i < rows; i++) {
            this.cells[i] = new Array(columns);
        }
        for(let i = 0; i < rows; i++) {
            for(let j = 0; j < columns; j++) {
                this.cells[i][j] = new DeadCell();
            }
        }
    }

    generateSeed() {
        let rand = Math.floor(Math.random() * 2);

        for(let i = 0; i < this.rows; i++) {
            for(let j = 0; j < this.columns; j++) {
                if(rand == 0) {
                    this.cells[i][j].cell = new DeadCell();
                } else {
                    this.cells[i][j].cell = new LiveCell();
                }
            }
        }
    }

    countAliveNeighbours(neighbours) {
        let aliveNeighbours = 0;
        for(let neighbour of neighbours) {
            if(this.cells[neighbour.x][neighbour.y].isAlive()) {
                aliveNeighbours++;
            }
        }
        return aliveNeighbours;
    }

    put(row, column, cell) {
        if( row < 0 || row >= this.rows || column < 0 || column >= this.columns) {
            throw new Error('Invalid row or column');
        }
        this.cells[row][column] = cell;

    }

    tickNewGeneration() {
        let newGrid = new Grid(this.cells.length, this.cells.length);

        for(let row = 0; row < this.cells.length; row++) {
            for(let column = 0; column < this.cells[0].length; column++) {
                let cell = this.cells[row][column];
                let neighbours = new CellAddress(row, column).getNeighbours(this);
                let aliveNeighbours = this.countAliveNeighbours(neighbours);
                let newCell = cell.nextGeneration(aliveNeighbours);
                newGrid.put(row, column, newCell);
            }
        }
        return newGrid;
    }
}

module.exports = {Grid};