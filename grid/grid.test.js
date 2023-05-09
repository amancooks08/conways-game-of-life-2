const {Grid} = require('./grid');

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
        // generate a new grid and with that a seed, which shouldn't be equal to the original grid
        it("should return a new grid" , () => {
            expect(grid.generateSeed()).not.toBe(grid);
        });
    });
});