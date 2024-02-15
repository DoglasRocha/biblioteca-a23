export const isValidUserInput = (isInvalidObject) => {
    for (let key in isInvalidObject) {
        if (isInvalidObject[key]) return false;
    }
    return true;
};