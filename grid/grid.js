const {Cell, LiveCell, DeadCell} = require('../cell/cell');

class Grid {

    constructor(rows, columns) {
        if (rows <= 0) {
            throw new Error('Invalid rows');
        }
        if (columns <= 0) {
            throw new Error('Invalid columns');
        }
        this.rows = rows;
        this.columns = columns;
        this.cells = new Array(rows);
        for (let i = 0; i < rows; i++) {
            this.cells[i] = new Array(columns);
        }
        for(let i = 0; i < rows; i++) {
            for(let j = 0; j < columns; j++) {
                this.cells[i][j] = new Cell(new DeadCell());
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
}

module.exports = {Grid};