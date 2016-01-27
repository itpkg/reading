export const randomColor = function () {
    return {color: "#" + ((1 << 24) * Math.random() | 0).toString(16)};
};