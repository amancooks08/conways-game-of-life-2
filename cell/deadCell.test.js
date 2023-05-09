// Tests of the DeadCell class
const {DeadCell} = require('./cell');
describe('DeadCell', () => {

    let deadCell;

    beforeEach(() => {
        deadCell = new DeadCell();
    });

    it('should return false', () => {
        expect(deadCell.isAlive()).toBe(false);
    });
});