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
}

module.exports = {Grid};