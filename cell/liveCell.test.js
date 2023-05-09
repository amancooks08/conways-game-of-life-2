const { DeadCell, LiveCell } = require("./cell");

describe("LiveCell", () => {
  let liveCell;

  beforeEach(() => {
    liveCell = new LiveCell();
  });

  it("should return true", () => {
    expect(liveCell.isAlive()).toBe(true);
  });

  describe("nextGeneration", () => {
    it("should return a DeadCell if it has less than 2 live neighbours", () => {
      expect(liveCell.nextGeneration(0) instanceof DeadCell).toBe(true);
      expect(liveCell.nextGeneration(1) instanceof DeadCell).toBe(true);
    });

    it("should return a LiveCell if it has 2 or 3 live neighbours", () => {
      expect(liveCell.nextGeneration(2) instanceof LiveCell).toBe(true);
      expect(liveCell.nextGeneration(3) instanceof LiveCell).toBe(true);
    });

    it("should return a DeadCell if it has more than 3 live neighbours", () => {
      expect(liveCell.nextGeneration(4) instanceof DeadCell).toBe(true);
    });
  });
});
