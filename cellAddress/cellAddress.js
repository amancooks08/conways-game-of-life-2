class CellAddress {
    //throws error on negative row or column
    constructor(row, column) {
        if (row < 0) {
            throw new Error('Invalid row');
        }
        if (column < 0) {
            throw new Error('Invalid column');
        }
        this.x = row;
        this.y = column;
    }

    //returns an array of CellAddress objects representing the neighbours of the current cell
    getNeighbours(grid) {
        let lastRowIndex = grid.cells.length - 1;
        let lastColumnIndex = grid.cells[0].length - 1;
        let neighbours = [];
        for(let row = this.x - 1; row <= this.x + 1; row++) {
            for(let column = this.y - 1; column <= this.y + 1; column++) {
                if(row < 0 || row > lastRowIndex || column < 0 || column > lastColumnIndex) {
                    continue;
                }
                if(row == this.x && column == this.y) {
                    continue;
                }
                neighbours.push(new CellAddress(row, column));
            }
        }
        return neighbours;
    }
}

module.exports = {
    CellAddress
};