const {Grid} = require('./grid');
const {LiveCell} = require('../cell/cell');
describe('Grid', () => {
    let grid;

    beforeEach(() => {
        grid = new Grid(3, 3);
    });

    describe('constructor', () => {
        it("should not throw error for valid sizes" , () => {
            expect(() => new Grid(3, 3)).not.toThrowError();
        });
    
        it("should throw error for negative rows" , () => {
            expect(() => new Grid(-3, 3)).toThrowError('Invalid rows');
        });
    
        it("should throw error for negative columns" , () => {
            expect(() => new Grid(3, -3)).toThrowError('Invalid columns');
        });
    
        it("should throw error for zero rows" , () => {
            expect(() => new Grid(0, 3)).toThrowError('Invalid rows');
        });
    
        it("should throw error for zero columns" , () => {
            expect(() => new Grid(3, 0)).toThrowError('Invalid columns');
        });
    });

    
    describe('generateSeed', () => {
        it("should return a new grid" , () => {
            expect(grid.generateSeed()).not.toBe(grid);
        });
    });


    describe('TickNewGeneration', () => {
        it("should return a new grid" , () => {
            grid.put(0, 0, new LiveCell());
            grid.put(0, 1, new LiveCell());
            grid.put(0, 2, new LiveCell());
            console.log(grid);
            let newGrid = grid.tickNewGeneration();
            console.log(newGrid);
            let expectedGrid = new Grid(3, 3);
            console.log(expectedGrid);

            expect(newGrid).not.toBe(grid);
        });
    });
});