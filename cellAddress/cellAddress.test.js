const {CellAddress} = require('./cellAddress');

describe('CellAddress', () => {

    let cellAddress;

    beforeEach(() => {
        cellAddress = new CellAddress(1, 2);
    });

    it('should return the correct row', () => {
        expect(cellAddress.x).toBe(1);
    });

    it('should return the correct column', () => {
        expect(cellAddress.y).toBe(2);
    });

    it('should throw error for negative row', () => {
        expect(() => new CellAddress(-1, 2)).toThrowError('Invalid row');
    });

    it('should throw error for negative column', () => {
        expect(() => new CellAddress(1, -2)).toThrowError('Invalid column');
    });
});