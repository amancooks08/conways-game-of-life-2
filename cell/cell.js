class Cell {

    // constructor() {}

    isAlive() {}
    
    nextGeneration(){}
}

class DeadCell extends Cell {
    constructor() {
        super();
    }

    isAlive() {
        return false;
    }

    nextGeneration(numberOfLiveNeighbours) {
        if (numberOfLiveNeighbours === 3) {
            return new LiveCell();
        }
        return this;
    }
}

class LiveCell extends Cell {
    constructor() {
        super();
    }

    isAlive() {
        return true;
    }
    
    nextGeneration(numberOfLiveNeighbours) {
        if (numberOfLiveNeighbours < 2 || numberOfLiveNeighbours > 3) {
            return new DeadCell();
        }
        return this;
    }
}

module.exports = {
    Cell,
    DeadCell,
    LiveCell
};