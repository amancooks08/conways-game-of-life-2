class CellAddress {
    //throws error on negative row or column
    constructor(row, column) {
        if (row < 0) {
            throw new Error('Invalid row');
        }
        if (column < 0) {
            throw new Error('Invalid column');
        }
        this.row = row;
        this.column = column;
    }
}

module.exports = {
    CellAddress
};