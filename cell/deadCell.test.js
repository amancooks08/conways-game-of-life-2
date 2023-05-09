// Tests of the DeadCell class
const {DeadCell, LiveCell} = require('./cell');
describe('DeadCell', () => {

    let deadCell;

    beforeEach(() => {
        deadCell = new DeadCell();
    });

    it('should return false', () => {
        expect(deadCell.isAlive()).toBe(false);
    });


    describe ('nextGeneration', () => {
        it("should return a LiveCell if it has 3 live neighbours", () => {
           expect(deadCell.nextGeneration(3) instanceof LiveCell).toBe(true);
        });

        it("should return a DeadCell if it has less than 3 live neighbours", () => {
            expect(deadCell.nextGeneration(0) instanceof DeadCell).toBe(true);
            expect(deadCell.nextGeneration(1) instanceof DeadCell).toBe(true);
        });

        it("should return a DeadCell if it has more than 3 live neighbours", () => {
            expect(deadCell.nextGeneration(4) instanceof DeadCell).toBe(true);
        });
    });
});