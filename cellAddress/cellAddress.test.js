const {CellAddress} = require('./cellAddress');

describe('CellAddress', () => {

    let cellAddress;

    beforeEach(() => {
        cellAddress = new CellAddress(1, 2);
    });

    it('should return the correct row', () => {
        expect(cellAddress.row).toBe(1);
    });

    it('should return the correct column', () => {
        expect(cellAddress.column).toBe(2);
    });
});